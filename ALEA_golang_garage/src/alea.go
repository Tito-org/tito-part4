/***********
Create by Hugo Janasik
Intern Developer
VMware
************/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min int, max int) int {
	return (rand.Intn(max-min) + min)
}

func defineAlea(nb int) {
	rand.Seed(time.Now().UnixNano())
	randomNum := random(1, nb+1)
	nbAlea = randomNum
}

func alea() {
	ticker := time.NewTicker(time.Second * 20)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
			if sqlQuery() != 84 {
				selectCar()
			}
		}
	}()
}
