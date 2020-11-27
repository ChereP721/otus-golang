package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		log.Fatalf(err.Error())
	}

	timeServer := time.Now().In(loc).Round(time.Second)
	fmt.Println("current time:", timeServer)

	timeExact, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatalf(err.Error())
	}

	timeExact = timeExact.In(loc).Round(time.Second)
	fmt.Println("exact time:", timeExact)
}
