package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

type Worker struct {
	address string
	on      bool
	mux     sync.RWMutex
	backlog int
}

type Library struct {
	workers []*Worker
	num     int
}

/*
func balancer(w http.ResponseWriter, r *http.Request) {
	peer := Library.GetNextPeer()
	if peer != nil {
		peer.ReverseProxy.ServeHTTP(w, r)
		return
	}
	http.Error(w, "Service not available", http.StatusServiceUnavailable)
}
*/

func main() {
	var portServerList string
	var urlServerList string
	var ipServerList string
	var port string
	flag.StringVar(&portServerList, "portBackends", "", "Use commas between backends to add.")
	flag.StringVar(&urlServerList, "urlBackends", "", "Use commas between backends to add.")
	flag.StringVar(&ipServerList, "ipBackends", "", "Use commas between backends to add.")
	flag.StringVar(&port, "serveOn", "3000", "Port to serve load balancer")
	flag.Parse()

	if len(portServerList) == 0 && len(urlServerList) == 0 && len(ipServerList) == 0 {
		log.Fatal("Please provide one or more backends to load balance")
	}

	// for readability
	var library *Library
	library = new(Library)

	// create pool of backends
	portTokens := strings.Split(portServerList, "/")
	for _, tok := range portTokens {
		AddWorker(tok, library, "port")
	}

	urlTokens := strings.Split(urlServerList, "/")
	if urlTokens[0] != "" {
		for _, tok := range urlTokens {
			AddWorker(tok, library, "url")
		}
	}

	ipTokens := strings.Split(ipServerList, "/")
	if ipTokens[0] != "" {
		for _, tok := range ipTokens {
			AddWorker(tok, library, "ip")
		}
	}

	PORT := ":" + port
	server, err := net.Listen("tcp4", PORT)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Serving on port %s\n", port)

	defer server.Close()

	go HandleCollapse()

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go HandleClientConn(conn, library)
	}
}
