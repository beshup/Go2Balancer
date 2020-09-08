package main

import (
	"fmt"
	"net"
)

func AddWorker(addr string, l *Library, addrType string) {
	if addrType == "port" {
		addr = ":" + addr
	} else {
		addr = addr + ":http"
	}
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		fmt.Printf("%s is not alive\n", addr)
		return
	}

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	//	status, err := bufio.NewReader(conn).ReadString('\n')

	//	fmt.Println(status)
	// if here, backend is alive
	var w *Worker
	w = new(Worker)
	w.address = addr
	w.on = true
	w.backlog = 0

	l.workers = append(l.workers, w)
	l.num += 1
	fmt.Printf("%s added to pool.\n", addr)
}
