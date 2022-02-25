package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/spf13/viper"
	sourse "./../"
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

	sourse.MQTTInit("./config", "global", f)
	api := sourse.MQTTClientAPI{}
	api.Publish("This is my message for you right now")
	<-c
}
