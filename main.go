package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v2"

	lockfile "github.com/nightlyone/lockfile"

	homeassist "github.com/user_idle/homeassist_mqtt"
	"github.com/user_idle/macos_idle"
)

const (
	// Config file name
	configF = "config.yaml"
	lockF   = ".user_idle"
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
	KeepAlive    int    `yaml:"keep_alive"`
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
	var timer = -1
	c.loadConf()

	// is there another instance of user idle running?

	lock, err := lockfile.New("/tmp/" + lockF)
	if err != nil {
		log.Fatal("Unable to create lock file", err)
	}

	if err = lock.TryLock(); err != nil {
		log.Fatal("Cannot lock, reason:", err)
	}

	// If the user has specified debug then print program information
	if c.Debug {
		fmt.Println("Version 1.0 - (C) Jim Colderwood\nConfig Array: ", c)
	}

	// Setup MQTT handler

	client := homeassist.Connect(c.MqttBroker, c.MqttPort, c.MqttUser, c.MqttPass, c.Debug)
	client.Connect()

	// Main loop
	for {
		time.Sleep(1 * time.Second)
		if macos_idle.Check() < c.ActivityTime && timer >= c.KeepAlive || timer == -1 {
			if c.Debug {
				fmt.Println("Sending Keep Alive!")
			}
			homeassist.Publish(client, c.MqttTopic)
			timer = 0
		}
		timer++
	}
}
