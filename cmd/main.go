package main

import (
	"guilt-type-service/config"
	"guilt-type-service/database"
	"guilt-type-service/internal/auth"
	"guilt-type-service/internal/handler"
	"guilt-type-service/internal/middleware"
	"guilt-type-service/internal/repository"
	"guilt-type-service/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to the database
	db := database.InitDB()

	// Setup dependency injection
	repo := repository.NewGuiltTypeRepository(db)
	svc := service.NewGuiltTypeService(repo)
	h := handler.NewGuiltTypeHandler(svc, repo)
	authClient := auth.NewJavaAuthClient("http://172.20.10.4:8081") // Replace IP if needed

	// Setup router
	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)

	adminRouter := router.PathPrefix("/api/guilt-types").Subrouter()
	adminRouter.Use(auth.JWTWithAuth(authClient, "ROLE_ADMIN"))

	adminRouter.Handle("", middleware.JWTAdminOnly(http.HandlerFunc(h.GetAll))).Methods(http.MethodGet)
	adminRouter.Handle("", middleware.JWTAdminOnly(http.HandlerFunc(h.Create))).Methods(http.MethodPost)
	adminRouter.Handle("/{id}", middleware.JWTAdminOnly(http.HandlerFunc(h.Update))).Methods(http.MethodPut)
	adminRouter.Handle("/{id}", middleware.JWTAdminOnly(http.HandlerFunc(h.Delete))).Methods(http.MethodDelete)
	adminRouter.Handle("/import", middleware.JWTAdminOnly(http.HandlerFunc(h.ImportExcel))).Methods(http.MethodPost)

	// Start HTTP server
	server := &http.Server{
		Addr:    ":8083",
		Handler: router,
	}

	log.Printf("🚀 Guilt Type Service started on port %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
