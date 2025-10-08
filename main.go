package main

import (
	"fmt"
	"log"

	"wg-monitor/app/internal/runner"
)

func main() {
	if output, err := runner.Run(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(output)
	}
}
