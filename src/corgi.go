package main

import (
	"encoding/json"
	"fmt"
	"github.com/clandry94/quickHTTPproxy/src/proxy"
	"github.com/clandry94/quickHTTPproxy/src/spec"
	"github.com/ivahaev/go-logger"
	"io/ioutil"
	"os"
)

func main() {
	logger.SetLevel("DEBUG")
	logger.Info("Initializing...")
	config, err := checkArgs()
	if err != nil {
		logger.Error(err)
		os.Exit(2)
	}

	proxySpec, err := loadConfig(config)
	if err != nil {
		logger.Errorf("Could not load spec: %v", err)
		panic(err)
	}
	logger.Info("Spec loaded")
	p := proxy.New(proxySpec)
	logger.Info("Proxy loaded")
	p.Listen()
}

func checkArgs() (string, error) {
	if len(os.Args) < 2 {
		return "", fmt.Errorf("No configuration file provided")
	}

	return os.Args[1], nil
}

func loadConfig(config string) (*spec.ProxySpec, error) {
	var ps *spec.ProxySpec

	j, err := ioutil.ReadFile(config)
	if err != nil {
		logger.Errorf("Unable to read configuration file: %v", config)
		return ps, err
	}

	err = json.Unmarshal(j, &ps)
	if err != nil {
		logger.Error("Unable to unmarshal config")
		return ps, err
	}
	return ps, err
}
