package grpc

import (
	contact "architecture_go/services/api/internal/delivery/grpc/interface"
	"context"
)

type Client struct {
}

func (c *Client) CreateGroup(ctx context.Context, request *contact.CreateGroupRequest) {

}
