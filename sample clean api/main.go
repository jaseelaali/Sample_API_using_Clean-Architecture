package main

import (
	"log"

	dependency "github.com/jaseelaali/Sample_API_using_Clean-Architecture/dep"
	config "github.com/jaseelaali/Sample_API_using_Clean-Architecture/pkg/config"
)

func main() {
	//	LOAD CONFIG FILE
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("configuration error : %v", err)
	}
	//INITIALIZE API
	server, dberr := dependency.InitializeApi(config)
	if dberr != nil {
		log.Fatal(dberr)
	}
	server.Start()
}
