package main

import (
	"fmt"
	"time"
)

func HandleCollapse(l *Library) {
	t := time.NewTicker(time.Minute * 1)
	for {
		select {
		case <-t.C:
			fmt.Println("Checking if servers have collapsed")
			for __, w := range l.workers {
				status := "up"
				// avoiding race conditions here~
				SetOn(w, IsOn(w.address))
				if on != true {
					status = "down"
				}
				fmt.Println("%s is %s", w.address, status)
			}
			fmt.Println("Check complete")
		}
	}
}
