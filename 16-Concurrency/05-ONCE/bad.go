package main

import "sync"

var (
	mx     = sync.Mutex{}
	config *Config
)

func GetConfigWithSingleTon() *Config {
	if config == nil {
		mx.Lock()
		defer mx.Unlock()
		config = &Config{}
		println("creating new config")
	}
	println("got config")
	return config

}
