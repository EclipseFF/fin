package grpc

import (
	contact "architecture_go/services/api/internal/delivery/grpc/interface"
	"architecture_go/services/api/internal/useCase"
	"google.golang.org/grpc"
)

type Delivery struct {
	ucRequest   useCase.Request
	clientAdmin contact.RequestServiceClient

	options Options
}

type Options struct{}

func New( /*ucRequest useCase.Request, */ options Options) *Delivery {

	conn, err := grpc.Dial("localhost:5500")
	if err != nil {
		return nil
	}
	defer conn.Close()
	client := contact.NewRequestServiceClient(conn)
	var d = &Delivery{
		/*ucRequest: ucRequest,*/
		clientAdmin: client,
	}

	d.SetOptions(options)
	return d
}

func (d *Delivery) SetOptions(options Options) {
	if d.options != options {
		d.options = options
	}
}
