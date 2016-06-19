package main

import (
	"log"
	"os/exec"
	"regexp"
	"strconv"
)

func main() {
	err := notifyCritical()
	if err != nil {
		log.Fatalln(err)
	}
}

func notifyCritical() error {
	level, err := batteryLevel()
	if err != nil {
		return err
	}

	if critical(level) {
		err = notify(outputText(level), "critical")
		log.Print(outputText(level))
		if err != nil {
			return err
		}
	}
	return err
}

func notify(message string, level string) error {
	cmd := exec.Command("/usr/bin/notify-send", message, "-u", level)
	return cmd.Run()
}

func batteryLevel() (int, error) {
	var level int
	cmd := exec.Command("acpi")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return level, err
	}

	level, err = extractLevel(output)
	return level, err
}

func extractLevel(output []byte) (int, error) {
	var level int
	rp, err := regexp.Compile("(\\d+?)%")
	if err != nil {
		return level, err
	}

	bat := rp.FindSubmatch(output)[1]
	level, err = strconv.Atoi(string(bat))
	return level, err
}

func outputText(level int) string {
	return "Battery low! <b>" + strconv.Itoa(level) + "</b>%"
}

func critical(level int) bool {
	return level <= 20
}
