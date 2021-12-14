package main

import (
	"flag"

	"fmt"

	"github.com/dineshkumar181094/key-value-store/pkg/server"
	log "github.com/sirupsen/logrus"
)

var (
	port     = flag.Int("port", 8080, "Server port")
	loglevel = flag.String("level", "info", "loglevel `flags`, serveral allowed [debug,info,warn,error,fatal]")
)

func main() {
	flag.Parse()
	lvl, err := log.ParseLevel(*loglevel)
	if err != nil {
		fmt.Println("log version is not in [debug,info,warn,error,fatal] setting default to info")
		*loglevel = "info"
		lvl, _ = log.ParseLevel(*loglevel)
	}
	//
	log.SetLevel(lvl)
	log.Info("starting server!!!!")
	server.Start(*port)
}
