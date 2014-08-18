package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := 0

	for {
		n++

		i := rand.Intn(4200)
		fmt.Println(i)

		if n%42 == 0 {
			break
		}
	}

	fmt.Println("Exit after %d actions", n)
}
