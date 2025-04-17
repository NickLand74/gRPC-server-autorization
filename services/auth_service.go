package services

import (
	"context"
	"log"

	"github.com/NickLand74/gRPC-server-autorization/config"
	"github.com/NickLand74/gRPC-server-autorization/storage"
)

type AuthService struct {
	storage storage.Storage
}

func NewAuthService(storage storage.Storage) *AuthService {
	return &AuthService{storage: storage}
}

func (s *AuthService) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	err = s.storage.CreateUser(req.Username, hashedPassword)
	if err != nil {
		return &auth.RegisterResponse{Message: "Username already exists"}, nil
	}

	return &auth.RegisterResponse{Message: "Registration successful"}, nil
}

func (s *AuthService) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	user, err := s.storage.GetUser(req.Username)
	if err != nil {
		return &auth.LoginResponse{}, nil // Пустой ответ при ошибке
	}

	err = auth.CheckPassword(user.Password, req.Password)
	if err != nil {
		return &auth.LoginResponse{}, nil
	}

	token, err := auth.GenerateToken(req.Username, config.JWTSecret)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return nil, err
	}

	return &auth.LoginResponse{Token: token}, nil
}
