package handlers

import (
	"context"

	"github.com/NickLand74/gRPC-server-autorization/proto/auth/pb"
	"github.com/NickLand74/gRPC-server-autorization/services"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer // Используем pb вместо auth
	service                           *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return h.service.Register(ctx, req)
}

func (h *AuthHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return h.service.Login(ctx, req)
}
