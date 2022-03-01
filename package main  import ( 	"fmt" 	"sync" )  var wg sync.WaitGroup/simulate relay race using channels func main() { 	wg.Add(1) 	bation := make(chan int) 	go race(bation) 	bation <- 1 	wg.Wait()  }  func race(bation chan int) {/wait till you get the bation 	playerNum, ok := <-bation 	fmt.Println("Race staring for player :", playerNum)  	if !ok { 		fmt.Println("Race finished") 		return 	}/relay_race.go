package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

//simulate relay race using channels
func main() {
	wg.Add(1)
	bation := make(chan int)
	go race(bation)
	bation <- 1
	wg.Wait()
}

func race(bation chan int) {

	//wait till you get the bation
	playerNum, ok := <-bation
	fmt.Println("Race staring for player :", playerNum)

	if !ok {
		fmt.Println("Race finished")
		return
	}
	//racing
	for i := 0; i <= 10; i++ {
		fmt.Print(playerNum)
	}
	fmt.Println()

	if playerNum < 4 {
		go race(bation)
	}

	if playerNum == 4 {
		close(bation)
		wg.Done()
		return
	}

	bation <- playerNum + 1
}
