package handlers

import (
	"encoding/json"
	"net/http"
	"49hw/models"
	"49hw/services"

	"github.com/go-chi/render"
)

type UserHandler struct {
	UserService *services.UserService
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	err := h.UserService.RegisterUser(&user)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.JSON(w, r, map[string]string{"message": "User registered successfully"})
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	authUser, err := h.UserService.AuthenticateUser(user.Username, user.Password)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	tokenString, err := services.CreateToken(authUser.Username, authUser.Role)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.JSON(w, r, map[string]string{"token": tokenString})
}
