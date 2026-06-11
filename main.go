package main

import (
	"doro/internal/timer"
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome to Doro")

	timer := timer.New(25 * time.Minute)

	fmt.Println(timer)

}