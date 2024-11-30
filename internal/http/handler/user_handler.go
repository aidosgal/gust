package handler

import (
	"net/http"

	"github.com/aidosgal/gust/internal/dto"
)

type UserHandler struct {
}

type UserService interface {
    Login(req dto.LoginRequest) (dto.LoginRequest, error)
    Register(req dto.RegisterRequest) (dto.RegisterRequest, error)
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *UserHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	return
}
