package controller

import (
	"context"
	"log"
	"net"
	pb "user-service/pb"
	"user-service/repository"
	"user-service/service"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var userService service.UserNeo4jService

type UserGrpcController struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserGrpcController) UpdateAvatar(ctx context.Context, request *pb.Avatar) (*pb.Empty, error) {
	err := userService.UpdateAvatar(int(request.Id), request.Avatar)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (s *UserGrpcController) UpdateProfile(ctx context.Context, request *pb.Profile) (*pb.Empty, error) {
	err := userService.UpdateProfile(int(request.Id), request.Profile)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (s *UserGrpcController) GetUser(ctx context.Context, request *pb.ID) (*pb.User, error) {
	user, err := userService.GetUser(int(request.Id))
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:       int64(user.ID),
		Username: user.Username,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Profile:  user.Profile,
	}, nil
}

func InitGrpc(driver neo4j.Driver) {
	userService = service.UserNeo4jService{
		UserRepository: repository.UserNeo4jRepository{
			Drive: driver,
		},
	}

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &UserGrpcController{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
