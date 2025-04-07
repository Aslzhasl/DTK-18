package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"guilt-type-service/internal/repository"
	"guilt-type-service/internal/service"
	"guilt-type-service/internal/handler"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	repo := repository.NewGuiltTypeRepository(db)
	service := service.NewGuiltTypeService(repo)
	handler := handler.NewGuiltTypeHandler(service, repo)

	api := r.Group("/api/guilt-types")
	{
		api.GET("", handler.GetAll)
		api.POST("", handler.Create)
		api.PUT(":id", handler.Update)
		api.DELETE(":id", handler.Delete)
		api.POST("/import", handler.ImportExcel)
	}

	return r
}