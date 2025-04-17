package services

import (
	"context"
	"log"

	"github.com/NickLand74/gRPC-server-autorization/config"
	"github.com/NickLand74/gRPC-server-autorization/internal/auth"
	"github.com/NickLand74/gRPC-server-autorization/proto/auth/pb"
	"github.com/NickLand74/gRPC-server-autorization/services/storage"
)

type AuthService struct {
	storage storage.Storage
}

func NewAuthService(storage storage.Storage) *AuthService {
	return &AuthService{storage: storage}
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	err = s.storage.CreateUser(req.Username, hashedPassword)
	if err != nil {
		return &pb.RegisterResponse{Message: "Username already exists"}, nil
	}

	return &pb.RegisterResponse{Message: "Registration successful"}, nil
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.storage.GetUser(req.Username)
	if err != nil {
		return nil, err
	}

	err = auth.CheckPassword(user.Password, req.Password)
	if err != nil {
		return nil, err
	}

	token, err := auth.GenerateToken(req.Username, config.LoadConfig().JWTSecret)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return nil, err
	}

	return &pb.LoginResponse{Token: token}, nil
}
