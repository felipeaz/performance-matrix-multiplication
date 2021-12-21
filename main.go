package main

import (
	"fmt"
	"log"
	"matrix-multiplication/matrix"
	"strings"
)

func main() {
	var performance string
	fmt.Println("Performance mode? (Y/N)")
	_, err := fmt.Scanln(&performance)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting Process")
	switch strings.ToUpper(performance) {
	case "Y":
		matrix.StartProcess()
	default:
		matrix.Start()
	}
}
