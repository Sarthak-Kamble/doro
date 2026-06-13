package main

import (
	"doro/internal/timer"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: doro <minutes>")
		return
	}

	minutes, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("Please enter a valid number")
		return
	}

	fmt.Println("\n------------- Welcome to Doro -------------")

	timer := timer.New(time.Duration(minutes) * time.Minute)

	for !timer.IsFinished() {

		fmt.Println(timer.FormatRemaining())

		time.Sleep(time.Second)

		timer.Tick()

	}

	fmt.Println("Done!!!!!")

	fmt.Println(os.Args[1])


}