package grpc_api

import (
	"context"
	"log"

	"tgbot/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn    *grpc.ClientConn
	service proto.MessageServiceClient
}

func NewGrpcClient(address string) (*Client, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:    conn,
		service: proto.NewMessageServiceClient(conn),
	}, nil
}

func (c *Client) SendMessage(name string, text string) bool {
	req := &proto.Message{
		UserName: name,
		Text:     text,
	}

	resp, err := c.service.SendMessage(context.Background(), req)
	if err != nil {
		log.Printf("send message: %v", err)
		return false
	}

	return resp.Status
}

func (c *Client) Close() {
	c.conn.Close()
}
