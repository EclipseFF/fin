package main

import (
	"fmt"
	"github.com/spf13/viper"

	deliveryHttp "architecture_go/services/api/internal/delivery/http"
)

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetDefault("SERVICE_NAME", "apiService")
}

func main() {

	var (
		listenerHttp = deliveryHttp.New
	)

	go func() {
		fmt.Printf("service started successfully on http port: %d", viper.GetUint("HTTP_PORT"))
		if err = listenerHttp.Run(); err != nil {
			panic(err)
		}
	}()
}
