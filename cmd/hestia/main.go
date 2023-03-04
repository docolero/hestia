package main

import (
	"flag"
	"log"

	"github.com/docolero/hestia/api"
	"github.com/gin-gonic/gin"
)

func main() {

	listenAddr := flag.String("listenaddr", ":3000", "server address")
	flag.Parse()

	server := api.NewServer(*listenAddr, gin.Default())
	log.Printf("starting server on port %s", *listenAddr)
	log.Fatal(server.Start())
}
