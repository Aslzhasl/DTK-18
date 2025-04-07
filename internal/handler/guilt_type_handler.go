package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"guilt-type-service/internal/service"
	"guilt-type-service/internal/model"
	"guilt-type-service/internal/excel"
	"guilt-type-service/internal/repository"
	"os"
)

type GuiltTypeHandler struct {
	service service.GuiltTypeService
	repo    repository.GuiltTypeRepository
}

func NewGuiltTypeHandler(s service.GuiltTypeService, r repository.GuiltTypeRepository) *GuiltTypeHandler {
	return &GuiltTypeHandler{service: s, repo: r}
}

func (h *GuiltTypeHandler) GetAll(c *gin.Context) {
	res, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *GuiltTypeHandler) Create(c *gin.Context) {
	var req model.GuiltType
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *GuiltTypeHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req model.GuiltType
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.service.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *GuiltTypeHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *GuiltTypeHandler) ImportExcel(c *gin.Context) {
	filePath := c.Query("path")
	if filePath == "" {
		filePath = "./data.xlsx"
	}
	if _, err := os.Stat(filePath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Файл не найден"})
		return
	}
	err := excel.ImportFromExcel(filePath, h.repo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Импорт завершен"})
}
