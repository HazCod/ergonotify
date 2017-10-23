package main

import "time"
import "os"
import (
	"github.com/0xAX/notificator"
)

var notify *notificator.Notificator

func main() {

	path, err := os.Getwd()
	if err != nil {
    	os.Exit(1)
	}

	notify = notificator.New(notificator.Options{
		DefaultIcon: path + "/standup.jpg",
		AppName:     "Ergonomics",
	})

	var first = true

	t := time.NewTicker(time.Hour * 40)
	for {
		if first == true {
			first = false
		} else {
			notify.Push(
				"Hourly break",
				"Get up from your desk & start moving.",
				path + "/standup.jpg",
				notificator.UR_CRITICAL)
		}
	    <-t.C
	}

}
