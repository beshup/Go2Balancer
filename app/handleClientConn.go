package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func HandleClientConn(conn net.Conn) {
	fmt.Printf("Serving %s \n", conn.RemoteAddr().String())
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(data))
		if temp == "STOP" {
			break
		}

		// here forward data to chosen backend server
		result := "jfc"
		conn.Write([]byte(string(result)))
	}
	conn.Close()
}
