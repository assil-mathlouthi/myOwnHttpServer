package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	buffer := make([]byte, 8)
	line := ""
	go func() {
		defer close(out)
		for {
			n, err := f.Read(buffer)
			if err != nil {
				break
			}
			data := string(buffer[:n])

			if p := strings.Index(data, "\n"); p != -1 {
				line += data[:p]
				out <- line
				line = data[p+1:]
			} else {
				line += data
			}
		}
		if len(line) != 0 {
			out <- line
		}
	}()
	return out

}

func main() {

	listner, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listner.Accept()

		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Connection Accepted ✅")
		}

		getLinesChannel := getLinesChannel(conn)

		for line := range getLinesChannel {
			fmt.Println(line)
		}
		fmt.Println("the connection has been closed ❌")
	}

}
