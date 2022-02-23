package mqtt_client

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
)

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func MQTTInit(pathFile string, nameFile string) {
	// viper.SetConfigName("global")

	// viper.AddConfigPath("./../config")

	viper.SetConfigName(nameFile)
	viper.AddConfigPath(pathFile)

	configError := viper.ReadInConfig()

	if configError != nil {
		fmt.Println("Fatal error config file: %s \n", configError)
		panic("Fatal error")
	}

	password := viper.GetString("mqtt_client.password")
	username := viper.GetString("mqtt_client.username")
	topic := viper.GetString("mqtt_client.topic")

	broker_connect_type := viper.GetString("mqtt_broker.connect_type")
	host := viper.GetString("mqtt_broker.host")
	port := viper.GetString("mqtt_broker.port")

	// opts := MQTT.NewClientOptions().AddBroker("ws://127.0.0.1:9001/mqtt") //for websocket
	opts := MQTT.NewClientOptions().AddBroker(broker_connect_type + "://" + host + port) //for tcp connection

	opts.SetPassword(password)
	opts.SetUsername(username)
	opts.SetDefaultPublishHandler(f)

	opts.OnConnect = func(c MQTT.Client) {
		if token := c.Subscribe(topic, 0, f); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	client := MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Error", token)
		panic(token.Error())
	} else {
		fmt.Printf("Connected to server\n")
	}

}
