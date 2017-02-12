package main

import (
	"flag"
	"log"

	"github.com/oem/batcountry/battery"
	"github.com/oem/batcountry/notification"
)

var l = flag.Int("l", 20, "the min battery level when notifications start being send")

func main() {
	flag.Parse()

	err := notifyCritical()
	if err != nil {
		log.Fatalln(err)
	}
}

func notifyCritical() error {
	level, err := battery.Level()
	if err != nil {
		return err
	}

	if critical(level) {
		text := notification.OutputText(level)
		err = notification.Send(text, "critical")
		log.Println(text)
		if err != nil {
			return err
		}
	}
	return err
}

func critical(level int) bool {
	return level <= *l
}
