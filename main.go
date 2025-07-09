package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
	"time"

	"ragnarok-multifinance/internal/delivery/http"
	mysql "ragnarok-multifinance/internal/library/mysql"
	redis "ragnarok-multifinance/internal/library/redis"
	pkg "ragnarok-multifinance/internal/pkg"
	repository "ragnarok-multifinance/internal/repository"
	usecase "ragnarok-multifinance/internal/usecase"

	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Load config
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(basepath + "/config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Config error: %v", err)
	}

	// Setup logger
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)

	// Setup MySQL
	db, err := mysql.New(viper.GetString("MySQL.DSN"))
	if err != nil {
		logrus.Fatalf("MySQL connection error: %v", err)
	}

	// Setup Redis
	redisCfg := redis.Config{
		Addr:     viper.GetString("Redis.Addr"),
		Password: viper.GetString("Redis.Password"),
		DB:       viper.GetInt("Redis.DB"),
	}
	redisClient := redis.New(redisCfg)

	// Setup Product domain
	productRepo := repository.NewProductMySQL(db)
	productUC := usecase.NewProductUsecase(productRepo)
	productHandler := http.NewProductHandler(productUC)

	// Setup Gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(secure.New(secure.Config{
		FrameDeny:            true,
		ContentTypeNosniff:   true,
		BrowserXssFilter:     true,
		SSLRedirect:          false,
		STSSeconds:           31536000,
		STSIncludeSubdomains: true,
	}))

	// Health endpoint
	r.GET("/health", func(c *gin.Context) {
		if err := db.Exec("SELECT 1").Error; err != nil {
			c.Status(503)
			return
		}
		if err := redis.Ping(c.Request.Context(), redisClient); err != nil {
			c.Status(503)
			return
		}
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Ready endpoint
	r.GET("/ready", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ready"})
	})

	// Register product routes with JWT middleware
	api := r.Group("/api")
	api.Use(pkg.JWTMiddleware(viper.GetString("JWT.Secret")))
	productHandler.Register(api)

	srv := &httpServer{engine: r}

	wait := gracefulShutdownNew(ctx, 30*time.Second, []operationNew{
		{
			name: "server",
			op: func(ctx context.Context) error {
				return srv.Stop(ctx)
			},
		},
	})

	if err := srv.Run(); err != nil {
		log.Println(err)
	}

	<-wait
	os.Exit(0)
}

type httpServer struct {
	engine *gin.Engine
}

func (s *httpServer) Run() error {
	return s.engine.Run(":8080")
}

func (s *httpServer) Stop(ctx context.Context) error {
	// Implement graceful shutdown if needed
	return nil
}

type operationNew struct {
	name string
	op   func(ctx context.Context) error
}

func gracefulShutdownNew(ctx context.Context, timeout time.Duration, ops []operationNew) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		signal.Notify(s, os.Interrupt)
		<-s

		log.Println("shutting down")

		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("timeout %d ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})
		defer timeoutFunc.Stop()

		for _, op := range ops {
			innerOp := op.op
			innerKey := op.name
			log.Printf("cleaning up: %s", innerKey)
			if err := innerOp(ctx); err != nil {
				log.Printf("%s: clean up failed: %s", innerKey, err.Error())
				return
			}
			log.Printf("%s was shutdown gracefully", innerKey)
		}
		close(wait)
	}()
	return wait
}
