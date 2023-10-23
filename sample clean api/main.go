package main

import (
	"log"

	"github.com/jaseelaali/Sample_API_using_Clean-Architecture/dep"
	c "github.com/jaseelaali/Sample_API_using_Clean-Architecture/pkg/config"
)

func main() {
	//	LOAD CONFIG FILE
	config, err := c.LoadConfig()
	if err != nil {
		log.Fatal("configuration error : %v", err)
	}
	//INITIALIZE API
	server, dberr := dep.InitializeApi(config)
	if dberr != nil {
		log.Fatal(dberr)
	}
	server.Start()
}
