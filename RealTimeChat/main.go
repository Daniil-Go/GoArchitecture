package main

import (
	"GoArchitecture/RealTimeChat/server"
	"fmt"
	"log"
)

func main() {

	serv := server.New()
	fmt.Println("running ...")
	if err := serv.Start(); err != nil {
		log.Fatalf("start error: %v", err)
	}

}
