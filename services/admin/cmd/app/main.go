package main

import (
	"architecture_go/pkg/store/postgres"
	"architecture_go/pkg/tracing"
	"architecture_go/pkg/type/context"
	log "architecture_go/pkg/type/logger"
	deliveryGrpc "architecture_go/services/admin/internal/delivery/grpc"
	deliveryHttp "architecture_go/services/admin/internal/delivery/http"
	repositoryStorage "architecture_go/services/admin/internal/repository/storage/postgres"
	useCaseGroup "architecture_go/services/admin/internal/useCase/group"
	useCaseVideo "architecture_go/services/admin/internal/useCase/video"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetDefault("SERVICE_NAME", "adminPanel")
}

func main() {
	conn, err := postgres.New(postgres.Settings{})
	if err != nil {
		panic(err)
	}
	defer conn.Pool.Close()

	closer, err := tracing.New(context.Empty())
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = closer.Close(); err != nil {
			log.Error(err)
		}
	}()

	repoStorage, err := repositoryStorage.New(conn.Pool, repositoryStorage.Options{})
	if err != nil {
		panic(err)
	}
	var (
		ucVideo      = useCaseVideo.New(repoStorage, useCaseVideo.Options{})
		ucGroup      = useCaseGroup.New(repoStorage, useCaseGroup.Options{})
		_            = deliveryGrpc.New(ucVideo, ucGroup, deliveryGrpc.Options)
		listenerHttp = deliveryHttp.New(ucVideo, ucGroup, deliveryHttp.Options{})
	)

	go func() {
		fmt.Printf("service started successfully on http port: %d", viper.GetUint("HTTP_PORT"))
		if err = listenerHttp.Run(); err != nil {
			panic(err)
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh

}
