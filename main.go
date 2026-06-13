package main

import (
	"doro/internal/timer"
	"fmt"
	"time"
)

func main() {

	
	fmt.Println("\n------------- Welcome to Doro -------------")

	timer := timer.New(10 * time.Second)

	for !timer.IsFinished() {

		fmt.Println(timer.FormatRemaining())

		time.Sleep(time.Second)

		timer.Tick()

	}

	fmt.Println("Done!!!!!")

}