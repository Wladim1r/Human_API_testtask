package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/Wladim1r/testtask/docs"
	"github.com/Wladim1r/testtask/internal/db"
	"github.com/Wladim1r/testtask/internal/http-server/handlers"
	"github.com/Wladim1r/testtask/internal/http-server/middleware"
	"github.com/Wladim1r/testtask/internal/http-server/repository"
	"github.com/Wladim1r/testtask/internal/http-server/service"
	"github.com/Wladim1r/testtask/internal/lib/sl"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API Server
// @version 1.0.3
// @description API for managing human information with automatic age, gender and nationality detection
// @host localhost:8080
// @BasePath /

func main() {
	gin.SetMode(gin.ReleaseMode)

	logLevel := slog.LevelInfo
	if os.Getenv("DEBUG") == "true" {
		logLevel = slog.LevelDebug
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))
	slog.SetDefault(logger)

	db, err := db.InitDB()
	if err != nil {
		slog.Error("failed to init DB", sl.Err(err))
		os.Exit(1)
	}

	slog.Info("DB initialized")

	repo := repository.NewHumanRepository(db)
	serv := service.NewHumanService(repo)
	hand := handlers.NewHumanHandler(serv)

	router := gin.Default()
	slog.Info("router created")

	router.Use(gin.LoggerWithFormatter(middleware.Log))

	url := ginSwagger.URL("/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.GET("/api", hand.GetInfo)
	router.DELETE("/api/:id", hand.Delete)
	router.PATCH("/api/:id", hand.Patch)
	router.POST("/api", hand.Post)
	slog.Info("routers connected")

	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	defer stop()

	svr := &http.Server{
		Addr:    os.Getenv("SVR_PORT"),
		Handler: router,
	}

	go func() {
		if err := svr.ListenAndServe(); err != nil {
			slog.Error("failed to run server", sl.Err(err))
			os.Exit(1)
		}
	}()

	slog.Info("starting server",
		slog.String("port", os.Getenv("SVR_PORT")),
		slog.String("db_host", os.Getenv("DB_HOST")),
		slog.Bool("debug_mode", os.Getenv("DEBUG") == "true"),
	)

	<-ctx.Done()

	slog.Info("shutting down...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := svr.Shutdown(shutdownCtx); err != nil {
		slog.Error("failed to shutdown", sl.Err(err))
	} else {
		slog.Info("server stopped gracefully")
	}
}
