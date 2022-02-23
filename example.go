package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	mqtt_client "github.com/tinobt/mqtt_client"
	
	"github.com/spf13/viper"
)

var configPort string

var f  = func(payload []byte) {
	fmt.Printf("MSG: %s\n", payload)
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	viper.SetConfigName("global")
	viper.AddConfigPath("./config")
	configError := viper.ReadInConfig()

	if configError != nil {
		fmt.Println("Fatal error config file: %s \n", configError)
		panic("Fatal error")
	}

	MQTTInit("./config", "global", f)
	<-c
}
