package main

import (
	"log"
	"net"
	"time"
)

func IsOn(addr string) bool {
	timeout := 2 * time.Second
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		log.Println("Could not reach backend, error: ", err)
		return false
	}
	_ = conn.Close()
	return true
}
