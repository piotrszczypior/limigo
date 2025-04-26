package main

import (
	"fmt"
	"github.com/piotrszczypior/limigo/limiter"
	"time"
)

func main() {
	var lim = limiter.CreateTokenBucket(10, 12*time.Second)

	for i := 0; i < 12; i++ {
		time.Sleep(3 * time.Second)
		fmt.Printf("Request %d - limiter - %t\n", i, lim.Allow())
	}
}
