package main

import "fmt"

func main() {

	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	for range "go" {
		
		fmt.Println(<-messages)

	}

}