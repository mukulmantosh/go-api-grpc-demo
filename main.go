package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "go-api-grpc-demo/api"
    "go-api-grpc-demo/model"
    pb "go-api-grpc-demo/proto"

    "github.com/google/uuid"
    "google.golang.org/grpc"
)

type grpcServer struct {
    pb.UnimplementedUserServiceServer
    store *model.UserStore
}

func (s *grpcServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
    user, exists := s.store.Get(req.Id)
    if !exists {
        return nil, fmt.Errorf("user not found")
    }
    return &pb.User{
        Id:    user.ID,
        Name:  user.Name,
        Email: user.Email,
        Age:   user.Age,
    }, nil
}

func (s *grpcServer) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
    users := s.store.List()
    pbUsers := make([]*pb.User, len(users))
    for i, user := range users {
        pbUsers[i] = &pb.User{
            Id:    user.ID,
            Name:  user.Name,
            Email: user.Email,
            Age:   user.Age,
        }
    }
    return &pb.ListUsersResponse{Users: pbUsers}, nil
}

func (s *grpcServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
    user := &model.User{
        ID:    uuid.New().String(),
        Name:  req.Name,
        Email: req.Email,
        Age:   req.Age,
    }
    s.store.Create(user)
    return &pb.User{
        Id:    user.ID,
        Name:  user.Name,
        Email: user.Email,
        Age:   user.Age,
    }, nil
}

func (s *grpcServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {
    user := &model.User{
        ID:    req.Id,
        Name:  req.Name,
        Email: req.Email,
        Age:   req.Age,
    }
    if !s.store.Update(user) {
        return nil, fmt.Errorf("user not found")
    }
    return &pb.User{
        Id:    user.ID,
        Name:  user.Name,
        Email: user.Email,
        Age:   user.Age,
    }, nil
}

func (s *grpcServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
    success := s.store.Delete(req.Id)
    return &pb.DeleteUserResponse{Success: success}, nil
}

func main() {
    // Create shared store
    store := model.NewUserStore()

    // Setup REST server
    handler := api.NewHandler(store)
    router := api.SetupRouter(handler)
    httpServer := &http.Server{
        Addr:    ":8080",
        Handler: router,
    }

    // Setup gRPC server
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterUserServiceServer(s, &grpcServer{store: store})

    // Start servers in goroutines
    go func() {
        log.Printf("Starting HTTP server on :8080")
        if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("HTTP server error: %v", err)
        }
    }()

    go func() {
        log.Printf("Starting gRPC server on :50051")
        if err := s.Serve(lis); err != nil {
            log.Fatalf("gRPC server error: %v", err)
        }
    }()

    // Wait for interrupt signal
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    // Graceful shutdown
    log.Println("Shutting down servers...")
    
    // Shutdown HTTP server
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := httpServer.Shutdown(ctx); err != nil {
        log.Printf("HTTP server shutdown error: %v", err)
    }

    // Shutdown gRPC server
    s.GracefulStop()
    log.Println("Servers stopped")
}
