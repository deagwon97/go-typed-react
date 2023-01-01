package routes

import (
	"go-typed-react/docs"

	"github.com/gin-gonic/gin"

	apiRouter "go-typed-react/0_backend/4_router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run(address string) error {

	router := gin.New()
	// cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}

	// backend APIs
	v1 := router.Group("/api/v1")
	apiRouter.AddRoutes(v1)

	// api docs (swagger)
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// frontend pages
	router.Use(cors.New(config))
	router.Use(static.Serve("/", static.LocalFile("./1_frontend/build", true)))

	return router.Run(address)
}
