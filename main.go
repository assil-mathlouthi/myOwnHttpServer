package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	myFile, err := os.Open("message.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer myFile.Close()

	buffer := make([]byte, 8)

	line := ""

	for {
		n, err := myFile.Read(buffer)
		if err != nil {
			break
		}

		data := string(buffer[:n])

		if p := strings.Index(data, "\n"); p != -1 {
			line += data[:p]
			fmt.Printf("read:%s\n", line)
			line = data[p+1:]
		} else {
			line += data
		}
	}
	if len(line) != 0 {
		fmt.Println(line)
	}

}
