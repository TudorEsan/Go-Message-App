package service

import (
	"context"
	pb "github.com/TudorEsan/Go-Message-App/protos"
)

func NewMessageService() pb.MessageServiceServer {
	return &MessageService{}
}

type MessageService struct {
	pb.UnimplementedMessageServiceServer
}

func (m *MessageService) Ping(context.Context, *pb.Empty) (*pb.PingResponse, error) {
	return &pb.PingResponse{
		Message: "Pong",
	}, nil
}

func (m *MessageService) ConnectUsers(pb.MessageService_ConnectUsersServer) error {
	return nil
}

func (m *MessageService) PermissionToConnect(context.Context, *pb.Message) (*pb.Message, error) {
	return nil, nil
}
