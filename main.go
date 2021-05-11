package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v2"

	homeassist "github.com/user_idle/homeassist_mqtt"
	"github.com/user_idle/macos_idle"
)

const (
	// Config file name
	configF = "config.yaml"
)

// Struct to contain config
type config struct {
	MqttBroker   string `yaml:"mqtt_broker"`
	MqttPort     int    `yaml:"mqtt_port"`
	MqttUser     string `yaml:"mqtt_user"`
	MqttPass     string `yaml:"mqtt_pass"`
	MqttTopic    string `yaml:"mqtt_topic"`
	ActivityTime uint   `yaml:"activity_time"`
	Debug        bool   `yaml:"debug"`
	KeepAlive    uint   `yaml:"keep_alive"`
}

// Read local config file
func (c *config) loadConf() *config {
	configFile, err := ioutil.ReadFile(configF)
	if err != nil {
		log.Fatal("Unable to open config file: ", configF)
	}
	err = yaml.Unmarshal(configFile, c)
	if err != nil {
		log.Fatal("Something is wrong with the config file! Error: ", err)
	}
	return c
}

func main() {
	// Create a config object
	var c config
	frun := true
	var timer uint
	c.loadConf()

	// If the user has specified debug then print program information
	if c.Debug {
		fmt.Println("Version 1.0 - (C) Jim Colderwood\nConfig Array: ", c)
	}

	// Setup MQTT handler
	// TODO Error checking needs to return error
	client := homeassist.Connect(c.MqttBroker, c.MqttPort, c.MqttUser, c.MqttPass, c.Debug)
	client.Connect()

	// Main loop
	for {
		time.Sleep(1 * time.Second)
		timer++
		if macos_idle.Check() < c.ActivityTime && timer >= c.KeepAlive || frun {
			if c.Debug {
				fmt.Println("Sending Keep Alive!")
			}
			// If start up send packet then set first run false
			if frun {
				frun = !frun
			}
			homeassist.Publish(client, c.MqttTopic)
			timer = 0
		}
	}
}
