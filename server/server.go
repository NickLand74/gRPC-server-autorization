package server

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/NickLand74/gRPC-server-autorization/config"
	"github.com/NickLand74/gRPC-server-autorization/handlers"
	"github.com/NickLand74/gRPC-server-autorization/proto/auth/pb"
	"github.com/NickLand74/gRPC-server-autorization/services"
	"github.com/NickLand74/gRPC-server-autorization/services/storage"
)

// Функция для обработки паник
func recoverPanic() {
	if r := recover(); r != nil {
		log.Printf("PANIC RECOVERED: %v\nStack Trace:\n%s", r, debug.Stack())
		os.Exit(1) // Завершаем процесс с кодом ошибки
	}
}

func Run() {
	cfg := config.LoadConfig()

	// Инициализация хранилища
	storage := storage.NewPostgresStorage()
	service := services.NewAuthService(storage)
	handler := handlers.NewAuthHandler(service)

	// Настройка gRPC-сервера
	listener, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(func(ctx context.Context, req interface{},
			info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			// Логирование ошибок и паник
			defer func() {
				if r := recover(); r != nil {
					log.Printf("PANIC IN HANDLER: %v\nStack Trace:\n%s", r, debug.Stack())
				}
			}()
			return handler(ctx, req)
		}),
	)

	pb.RegisterAuthServiceServer(grpcServer, handler)
	reflection.Register(grpcServer)

	// Обработка сигналов
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	// Запуск сервера в горутине
	serverCtx, serverCancel := context.WithCancel(context.Background())
	go func() {
		defer recoverPanic()
		if err := grpcServer.Serve(listener); err != nil && err != grpc.ErrServerStopped {
			log.Printf("gRPC server error: %v", err)
		}
		serverCancel()
	}()

	// Ожидание сигнала завершения
	select {
	case <-quit:
		log.Println("Received shutdown signal. Gracefully stopping server...")
	case <-serverCtx.Done():
		log.Println("Server encountered an error. Gracefully stopping...")
	}

	// Завершение сервера
	grpcServer.GracefulStop()
	log.Println("Server stopped")
}
