package delivery

import (
	deliveryGrpc "architecture_go/services/api/internal/delivery/grpc"
	deliveryHttp "architecture_go/services/api/internal/delivery/http"
)

type Delivery struct {
	ListenerHttp *deliveryHttp.Delivery
	ListenerGrpc *deliveryGrpc.Delivery
}

type Options struct{}

func New() Delivery {
	Delivery := Delivery{
		ListenerHttp: deliveryHttp.New(deliveryHttp.Options(Options{})),
		ListenerGrpc: deliveryGrpc.New(deliveryGrpc.Options(Options{})),
	}

	return Delivery
}
