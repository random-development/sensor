package internal

import (
	"flag"
	"path"

	"github.com/spf13/viper"
)

// Config represents sensor configuration
type Config struct {
	Name       string
	Probes     []ProbeConfig
	Publishers []PublisherConfig
}

// ProbeConfig represents configuration of a probe
type ProbeConfig struct {
	Type     string
	Interval int
}

// PublisherConfig represents configuration of a publisher
type PublisherConfig struct {
	Type string
	URL  string
}

// ReadConfig reads configuration from configuration YAML file located at []
// Example config:
//
// name: resource-1
// probes:
// 	- type: memory
// 	  interval: 10
// 	- type: cpu
// 	  interval: 20
// publishers:
// 	- type: websocket
// 	  url: ws://demos.kaazing.com/echo
func ReadConfig() Config {
	configFile := flag.String("config", ".", "config file path")
	flag.Parse()
	p := path.Dir(*configFile)
	n := path.Base(*configFile)
	viper.SetConfigName(n)
	viper.AddConfigPath(p)

	if err := viper.ReadInConfig(); err != nil {
		Log.Fatalf("Error reading config file, %v", err)
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		Log.Fatalf("Unable to decode config into struct, %v", err)
	}
	Log.Debugf("Name: %s", config.Name)
	Log.Debugf("Read configuration, %v", config)
	return config
}
