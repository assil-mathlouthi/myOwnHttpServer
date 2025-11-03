package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	channel := make(chan string)

	buffer := make([]byte, 8)
	line := ""
	go func() {
		defer close(channel)
		for {
			n, err := f.Read(buffer)
			if err != nil {
				break
			}
			data := string(buffer[:n])

			if p := strings.Index(data, "\n"); p != -1 {
				line += data[:p]
				channel <- line
				line = data[p+1:]
			} else {
				line += data
			}
		}
		if len(line) != 0 {
			channel <- line
		}
	}()
	return channel

}

func main() {
	myFile, err := os.Open("message.txt")
	if err != nil {
		fmt.Println(err)
	}
	getLinesChannel := getLinesChannel(myFile)
	defer myFile.Close()

	for line := range getLinesChannel {
		fmt.Printf("read:%s\n", line)
	}

}
