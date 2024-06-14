package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"test/internal/config"
	handlermanager "test/internal/handlers/manager"
	"test/pkg/client/postgresql"
	"test/pkg/logging"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	cfg := config.GetConfig()

	logger := logging.GetLogger()
	postgresSQLClient, err := postgresql.NewClient(context.TODO(), cfg.Storage)
	fmt.Println(postgresSQLClient)
	if err != nil {
		logger.Fatalf("%v", err)
	}

	start(handlermanager.Manager(postgresSQLClient, logger), cfg, postgresSQLClient)
}

func start(router *gin.Engine, cfg *config.Config, pGPool *pgxpool.Pool) {
	logger := logging.GetLogger()
	logger.Info("start application")

	var listener net.Listener
	var listenErr error

	logger.Info("listen tcp")
	listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s",cfg.Listen.BindIP, cfg.Listen.Port))
	logger.Infof("server is listening port %s:%s",cfg.Listen.BindIP, cfg.Listen.Port)

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	}))

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 5000 * time.Second,
		ReadTimeout:  5000 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
