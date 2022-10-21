package service

import (
	"context"
	pb "github.com/TudorEsan/Go-Message-App/protos"
)

func NewMessageService() pb.MessageServiceServer {
	return &MessageService{
		onlineUsers:  make(map[string]pb.User),
		messageStream: make(map[string]pb.MessageService_ConnectToMessageChannelServer),
	}
}

type MessageService struct {
	pb.UnimplementedMessageServiceServer
	onlineUsers map[string]pb.User
	userMessageChannel map[string]chan pb.Message
	messageStream map[string]pb.MessageService_ConnectToMessageChannelServer
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

func (m *MessageService) SendMessage(context.Context, *pb.Message) (*pb.SendMessageResp, error) {
	return nil, nil
}

func (m *MessageService) ConnectToMessageChannel(empty *pb.User, messageService pb.MessageService_ConnectToMessageChannelServer) error {
	go func() {
		messageService.R
	}()
	return nil
}

