package main

import "sync"

var once sync.Once

func GetConfigWithOnce() *Config {
	once.Do(func() {
		println("Creating config")
	})
	println("got config")
	return config
}
