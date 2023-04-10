package service

import (
	"bot_svc/internal/domain"
	"bot_svc/pkg/proto"
	"context"
	"time"
)

type Storage interface {
	Insert(ctx context.Context, msg *domain.Note) (string, error)
}

func New(st Storage) *Service {
	return &Service{storage: st}
}

type Service struct {
	storage Storage
	proto.UnimplementedMessageServiceServer
}

func (s *Service) InsertNote(ctx context.Context, msg *proto.NewNote) (*proto.Response, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	msgDTO := domain.Note{}
	msgDTO.Fill(msg)
	if err := msgDTO.Validate(); err != nil {
		return nil, err
	}
	// insert to db
	id, err := s.storage.Insert(ctx, &msgDTO)
	if err != nil {
		return nil, err
	}

	return &proto.Response{
		IdNote: id,
		Status: true,
	}, nil
}
