package handlers

import (
	"NickLand74/gRPC-server-autorization.git/services"
	"context"
)

type AuthHandler struct {
	auth.UnimplementedAuthServiceServer
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	return h.service.Register(ctx, req)
}

func (h *AuthHandler) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	return h.service.Login(ctx, req)
}
