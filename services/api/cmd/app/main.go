package main

import (
	"architecture_go/services/api/internal/delivery"
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetDefault("SERVICE_NAME", "apiService")
}

func main() {

	var listener = delivery.New()

	go func() {
		fmt.Printf("service started successfully on http port: %d", viper.GetUint("HTTP_PORT"))
		if err := listener.ListenerHttp.Run(); err != nil {
			panic(err)
		}

	}()

}
