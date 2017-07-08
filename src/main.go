package main

import (
	"encoding/json"
	"github.com/clandry94/quickHTTPproxy/src/proxy"
	"github.com/clandry94/quickHTTPproxy/src/spec"
	"github.com/ivahaev/go-logger"
	"io/ioutil"
	"os"
)

func main() {
	logger.SetLevel("DEBUG")
	logger.Info("Initializing...")
	config := os.Args[1]
	handlerSpec, err := loadConfig(config)
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	logger.Info("Spec loaded")
	proxy := proxy.New(handlerSpec)
	logger.Info("Proxy loaded", proxy)
}

func loadConfig(config string) (*spec.HandlerSpec, error) {
	var hs *spec.HandlerSpec

	j, err := ioutil.ReadFile(config)
	if err != nil {
		logger.Error("Unable to read configuration file", config)
		return hs, err
	}

	err = json.Unmarshal(j, &hs)
	if err != nil {
		logger.Error("Unable to unmarshal config")
		return hs, err
	}
	return hs, err
}
