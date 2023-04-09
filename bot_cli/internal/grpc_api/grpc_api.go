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

func (c *Client) InsertNote(name string, text string) bool {
	req := proto.NewNote{
		UserName: name,
		Body:     text,
	}

	resp, err := c.service.InsertNote(context.Background(), &req)
	if err != nil {
		log.Printf("send message: %v", err)
		return false
	}

	return resp.Status
}

func (c *Client) Close() {
	c.conn.Close()
}
