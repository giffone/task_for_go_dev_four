package service

import (
	"bot_svc/internal/domain"
	"bot_svc/pkg/proto"
	"context"
	"time"
)

type Storage interface {
	Insert(ctx context.Context, msg *domain.Message) (string, error)
}

func New(st Storage) *Service {
	return &Service{storage: st}
}

type Service struct {
	storage Storage
	proto.UnimplementedMessageServiceServer
}

func (s *Service) SendMessage(ctx context.Context, msg *proto.Message) (*proto.Response, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	msgDTO := domain.Message{}
	msgDTO.Fill(msg)
	id, err := s.storage.Insert(ctx, &msgDTO)
	if err != nil {
		return nil, err
	}

	return &proto.Response{
		Id:     id,
		Status: true,
	}, nil
}
