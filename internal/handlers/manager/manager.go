package handlermanager

import (
	adminlogin "test/internal/admin/admin"
	admindb "test/internal/admin/admin/db"
	"test/pkg/client/postgresql"
	"test/pkg/logging"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)
const (
	adminURL  = "/api/v1/admin"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func Manager(client postgresql.Client, logger *logging.Logger) *gin.Engine {
	r := gin.Default()

	r.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	AdminRouterManager := r.Group(adminURL)
	AdminRouterRepository := admindb.NewRepository(client, logger)
	AdminRouterHandler := adminlogin.NewHandler(AdminRouterRepository, logger)
	AdminRouterHandler.Register(AdminRouterManager)
	return r
}