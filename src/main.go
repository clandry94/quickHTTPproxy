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
	proxySpec, err := loadConfig(config)
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	logger.Info("Spec loaded")
	p := proxy.New(proxySpec)
	logger.Info("Proxy loaded", p)
	p.Listen()
}

func loadConfig(config string) (*spec.ProxySpec, error) {
	var ps *spec.ProxySpec

	j, err := ioutil.ReadFile(config)
	if err != nil {
		logger.Error("Unable to read configuration file", config)
		return ps, err
	}

	err = json.Unmarshal(j, &ps)
	if err != nil {
		logger.Error("Unable to unmarshal config")
		return ps, err
	}
	return ps, err
}
