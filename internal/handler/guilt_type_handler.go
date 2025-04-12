package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"guilt-type-service/internal/excel"
	"guilt-type-service/internal/model"
	"guilt-type-service/internal/repository"
	"guilt-type-service/internal/service"
)

type GuiltTypeHandler struct {
	service service.GuiltTypeService
	repo    repository.GuiltTypeRepository
}

func NewGuiltTypeHandler(s service.GuiltTypeService, r repository.GuiltTypeRepository) *GuiltTypeHandler {
	return &GuiltTypeHandler{service: s, repo: r}
}

func (h *GuiltTypeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, data)
}

func (h *GuiltTypeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.GuiltType
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	res, err := h.service.Create(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, http.StatusOK, res)
}

func (h *GuiltTypeHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id") // you can use mux.Vars if using path params
	id, _ := strconv.Atoi(idStr)

	var req model.GuiltType
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	res, err := h.service.Update(uint(id), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, res)
}

func (h *GuiltTypeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	err := h.service.Delete(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *GuiltTypeHandler) ImportExcel(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Query().Get("path")
	if filePath == "" {
		filePath = "./data.xlsx"
	}

	err := excel.ImportFromExcel(filePath, h.repo)
	if err != nil {
		http.Error(w, "Импорт не удался: "+err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"message": "Импорт завершен"})
}

// helper
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}
