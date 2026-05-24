package main

import (
	"fmt"
	"os"
)

func main() {
	if err := ScrapeScripture(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
