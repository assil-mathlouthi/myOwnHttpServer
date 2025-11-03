package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:42069")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	for {
		var input string
		fmt.Print("Enter message (or 'exit'): ")
		fmt.Scanln(&input)
		if input == "exit" {
			break
		}

		_, err := conn.Write([]byte(input + "\n"))
		if err != nil {
			fmt.Println("Error writing:", err)
			os.Exit(1)
		}
	}
}
