package grpc

import (
	"context"

	request "architecture_go/services/api/internal/delivery/grpc/interface"
)

func (d *Delivery) CreateGroup(ctx context.Context, code string) (*request.CreateGroupResponse, error) {
	group, err := d.clientAdmin.CreateGroup(ctx, &request.CreateGroupRequest{Code: code})
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (d *Delivery) AddVideo(ctx context.Context, code, content string) (*request.AddVideoResponse, error) {
	video, err := d.clientAdmin.AddVideo(ctx, &request.AddVideoRequest{
		Code:  code,
		Bytes: content,
	})
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (d *Delivery) ShowGroups(ctx context.Context) (request.RequestService_ShowGroupsClient, error) {
	groups, err := d.clientAdmin.ShowGroups(ctx, &request.GroupsQuery{})
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (d *Delivery) RemoveVideo(ctx context.Context, id string) (*request.RemoveVideoResponse, error) {
	res, err := d.clientAdmin.RemoveVideo(ctx, &request.RemoveVideoRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d *Delivery) Login(ctx context.Context, code string) (*request.CreateLoginResponse, error) {
	res, err := d.clientAdmin.Login(ctx, &request.LoginRequest{Code: code})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d *Delivery) DeleteGroup(ctx context.Context, id string) (*request.DeleteGroupResponse, error) {
	res, err := d.clientAdmin.DeleteGroup(ctx, &request.DeleteGroupRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
