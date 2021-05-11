package homeassist

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	//fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	//fmt.Printf("Connect lost: %v", err)
}

func Publish(client mqtt.Client, topic string) {
	// Publish the ACK message on chosen topic
	text := "ACK"
	token := client.Publish(topic, 0, false, text)
	token.Wait()
	time.Sleep(time.Second)
}

func Connect(broker string, port int, user string, pass string, debug bool) mqtt.Client {
	// Load MQTT settings
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("user_idle")
	opts.SetUsername(user)
	opts.SetPassword(pass)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	// Create client and return client to caller
	client := mqtt.NewClient(opts)
	return client
}
