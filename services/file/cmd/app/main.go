package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"architecture_go/pkg/tracing"
	"architecture_go/pkg/type/context"
	log "architecture_go/pkg/type/logger"
	deliveryGrpc "architecture_go/services/file/internal/delivery/grpc"
)

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetDefault("SERVICE_NAME", "contactService")
}

func main() {
	closer, err := tracing.New(context.Empty())
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = closer.Close(); err != nil {
			log.Error(err)
		}
	}()

	if err != nil {
		panic(err)
	}
	var (
		_ = deliveryGrpc.New( /*ucContact, ucGroup, */ deliveryGrpc.Options{})
	)

	go func() {
		fmt.Printf("service started successfully on http port: %d", viper.GetUint("HTTP_PORT"))
		/*		if err = listenerHttp.Run(); err != nil {
				panic(err)
			}*/
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh

}
