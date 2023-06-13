package http

import (
	"architecture_go/services/api/internal/delivery/grpc"
	"architecture_go/services/api/internal/useCase"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetDefault("HTTP_PORT", 80)
}

type Delivery struct {
	ucRequest    useCase.Request
	router       *gin.Engine
	grpcDelivery grpc.Delivery
	options      Options
}

type Options struct{}

func New(options Options) *Delivery {
	var d = &Delivery{}

	d.SetOptions(options)

	d.router = d.initRouter()
	return d
}

func (d *Delivery) SetOptions(options Options) {
	if d.options != options {
		d.options = options
	}
}

func (d *Delivery) Run() error {
	return d.router.Run(fmt.Sprintf(":%d", uint16(viper.GetUint("HTTP_PORT"))))
}
