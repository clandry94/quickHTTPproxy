package main

import (
	"encoding/json"
	"github.com/clandry94/quickHTTPproxy/src/spec"
	"github.com/golang/glog"
	"io/ioutil"
	"os"
)

func main() {
	config := os.Args[1]
	glog.Info("Starting...")
	glog.Infof("Unmarshalling %v", config)
	handlerSpec, err := loadConfig(config)
	if err != nil {
		glog.Fatal(err)
		panic(err)
	}
	glog.Info(handlerSpec)

}

func loadConfig(config string) (*spec.HandlerSpec, error) {
	var hs *spec.HandlerSpec

	j, err := ioutil.ReadFile(config)
	if err != nil {
		glog.Fatalf("Unable to read configuration file, %v", config)
		return hs, err
	}

	err = json.Unmarshal(j, &hs)
	if err != nil {
		glog.Fatal("Unable to unmarshal config")
		return hs, err
	}
	return hs, err
}
