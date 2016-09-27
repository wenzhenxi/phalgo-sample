package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	var (
		host = "127.0.0.1"
		port = "8080"
		remote = host + ":" + port
		data = make([]byte, 1024)
	)
	fmt.Println("Initiating server... (Ctrl-C to stop)")

	lis, err := net.Listen("tcp", remote)
	defer lis.Close()

	if err != nil {
		fmt.Println("Error when listen: ", remote)
		os.Exit(-1)
	}

	for {
		var res string
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("Error accepting client: ", err.Error())
			os.Exit(0)
		}

		go func(con net.Conn) {
			fmt.Println("New connection: ", con.RemoteAddr())
			for {
				length, err := con.Read(data)
				if err != nil {
					fmt.Printf("Client %v quit.\n", con.RemoteAddr())
					con.Close()
					return
				}
				res = string(data[0:length])
				fmt.Printf("%q said: %s\n", strings.Split(con.RemoteAddr().String(), "\r\n"), res)
				res = "You said:" + res
				con.Write([]byte(res))
			}
		}(conn)
	}

}