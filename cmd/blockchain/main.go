package main

import (
	"time"

	"github.com/iagoholekdev/blockchain-go/domains/models"
	"github.com/iagoholekdev/blockchain-go/domains/service"
)

func main() {
	trLocal := service.NewLocalTransport("LOCAL")
	trRemote := service.NewLocalTransport("REMOTE")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("Hello world"))
			time.Sleep(1 * time.Second)
		}

	}()

	opts := models.ServerOpts{
		Transports: []models.Transport{trLocal},
	}

	s := models.NewServer(opts)
	s.Start()
}
