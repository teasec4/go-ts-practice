package handler

import (
	"encoding/json"
	"justforfun/internal/service"
	"net/http"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{
		Service: s,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.Service.CreateUser(req.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Получить пользователя по UUID
func (h *UserHandler) FindUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	user, err := h.Service.FindUserByID(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
