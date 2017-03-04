package main

import (
	"flag"
	"log"

	"github.com/oem/batcountry/battery"
	"github.com/oem/batcountry/notification"
)

var l = flag.Int("l", 20, "the min battery level when notifications start being send")
var capFile = flag.String("c", "/sys/class/power_supply/BAT1/capacity", "the file containing the current battery capacity")
var statusFile = flag.String("s", "/sys/class/power_supply/BAT1/status", "the file containing the current battery status")

func main() {
	flag.Parse()

	err := notifyCritical()
	if err != nil {
		log.Fatalln(err)
	}
}

func notifyCritical() error {
	level, charging, err := battery.Level(*capFile, *statusFile)
	if err != nil {
		return err
	}

	if !charging && critical(level) {
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
