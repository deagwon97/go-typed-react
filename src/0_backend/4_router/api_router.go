package router

import (
	repository "go-typed-react/0_backend/1_repository"
	service "go-typed-react/0_backend/2_service"
	controller "go-typed-react/0_backend/3_controller"
	"go-typed-react/0_backend/database"

	"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup) {
	router := rg.Group("")
	gormDB, _ := database.CreateGormDB()
	rps, _ := repository.NewRepository(gormDB)
	svc, _ := service.NewService(rps)
	ctrl, _ := controller.NewController(svc)

	router.GET("/ping", ctrl.Ping)
	router.GET("/file", ctrl.CreateFile)

}
