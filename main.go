package main

import (
	"log"
	"time"

	"github.com/xiaoheigou/viper-t/config"
)

func main() {

	conf := config.TestConfReader

	log.Println(conf.GetAllKeys())
	log.Println(conf.GetString("test1"))

	log.Println(conf.GetString("test1"))
	go func() {
		for {
			log.Println(conf.GetString("test1"))
			time.Sleep(time.Second)
		}

	}()
	select {}
}
