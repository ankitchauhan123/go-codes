package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	balls := make(chan int)
	fmt.Println("Tennis Game using channels")
	go playTennis("P1", balls)
	go playTennis("P2", balls)
	balls <- 1
	wg.Wait()
}

func playTennis(playerName string, balls chan int) {
	defer wg.Done()
	for {
		ball, ok := <-balls
		if !ok {
			fmt.Println("Player has won:", playerName)
			return
		}

		n := rand.Intn(100)
		if n != 0 && ball%n == 0 {
			fmt.Println("Missed the ball:", playerName)
			close(balls)
			return
		}
		fmt.Println(playerName, ":", ball)
		balls <- ball + 1
	}

}
