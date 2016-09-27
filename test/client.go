package main

import (
	"fmt"
	"net"
	"os"
)

var str string
var msg = make([]byte, 1024)

func main() {
	var (
		host = "127.0.0.1"
		port = "8080"
		remote = host + ":" + port
	)

	con, err := net.Dial("tcp", remote)
	defer con.Close()

	if err != nil {
		fmt.Println("Server not found.")
		os.Exit(-1)
	}
	fmt.Println("Connection OK.")

	for {
		fmt.Printf("Enter a sentence:")
		fmt.Scanf("%s\n", &str)
		if str == "quit" {
			fmt.Println("Communication terminated.")
			os.Exit(1)
		}

		in, err := con.Write([]byte(str))
		if err != nil {
			fmt.Printf("Error when send to server: %d\n", in)
			os.Exit(0)
		}

		length, err := con.Read(msg)
		if err != nil {
			fmt.Printf("Error when read from server.\n")
			os.Exit(0)
		}
		str = string(msg[0:length])
		fmt.Println(str)
	}
}