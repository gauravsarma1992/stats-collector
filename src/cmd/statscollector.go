package main

import (
	"log"
	"os"

	"github.com/gauravsarma1992/statscollector/statscollector"
)

func main() {
	var (
		server *statscollector.Server
		err    error
	)

	if server, err = statscollector.NewServer(); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	if err = server.Run(); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}
