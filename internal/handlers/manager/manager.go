package handlermanager

import (
	"test/pkg/client/postgresql"
	"test/pkg/logging"

	"github.com/gin-gonic/gin"
)

func Manager(client postgresql.Client, logger *logging.Logger) *gin.Engine {
	router := gin.Default()


	return router
}