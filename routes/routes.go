package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"guilt-type-service/internal/auth"
	"guilt-type-service/internal/handler"
	"guilt-type-service/internal/middleware"
	"guilt-type-service/internal/repository"
	"guilt-type-service/internal/service"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Инициализация репозитория, сервиса и хендлера
	repo := repository.NewGuiltTypeRepository(db)
	svc := service.NewGuiltTypeService(repo)
	h := handler.NewGuiltTypeHandler(svc, repo)

	authClient := auth.NewJavaAuthClient("http://172.20.10.2:8081")

	adminRoutes := r.Group("/api/guilt-types")
	adminRoutes.Use(middleware.JWTMiddleware(authClient, "ROLE_ADMIN"))

	{
		adminRoutes.GET("", h.GetAll)
		adminRoutes.POST("", h.Create)
		adminRoutes.PUT(":id", h.Update)
		adminRoutes.DELETE(":id", h.Delete)
		adminRoutes.POST("/import", h.ImportExcel)
	}

	return r
}
