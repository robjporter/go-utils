package main

import (
	"fmt"
)

func main() {
	message := "This is a simple announcement !"
	mes := announcement(message)
	fmt.Println(mes)
	fmt.Println()
	makeannouncement(message)
}
