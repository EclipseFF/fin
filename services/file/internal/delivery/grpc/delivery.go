package grpc

type Delivery struct {
	route routeGuideServer

	options Options
}

type Options struct{}

func New(o Options) *Delivery {
	var d = &Delivery{
		route: routeGuideServer{},
	}

	d.SetOptions(o)
	return d
}

func (d *Delivery) SetOptions(options Options) {
	if d.options != options {
		d.options = options
	}
}
