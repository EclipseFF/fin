package grpc

import (
	admin "architecture_go/services/admin/internal/delivery/grpc/interface"
)
import "architecture_go/services/admin/internal/useCase"

type Delivery struct {
	admin.UnimplementedAdminPanelServiceServer
	ucVideo useCase.Video
	ucGroup useCase.Group

	options Options
}

type Options struct{}

func New(ucVideo useCase.Video, ucGroup useCase.Group, o Options) *Delivery {
	var d = &Delivery{
		ucVideo: ucVideo,
		ucGroup: ucGroup,
	}

	d.SetOptions(o)
	return d
}

func (d *Delivery) SetOptions(options Options) {
	if d.options != options {
		d.options = options
	}
}
