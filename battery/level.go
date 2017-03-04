package battery

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// Level reads the files written by the kernel to get status and capacity of the battery
func Level(capFile, statusFile string) (int, bool, error) {
	battery := 0
	charging := false
	cap, err := ioutil.ReadFile(capFile)
	if err != nil {
		return battery, charging, err
	}
	status, err := ioutil.ReadFile(statusFile)
	if err != nil {
		return battery, charging, err
	}
	battery, charging, err = extractLevel(cap, status)
	return battery, charging, err
}

func extractLevel(cap, status []byte) (int, bool, error) {
	charging := false
	if strings.TrimSpace(string(status)) == "Charging" {
		charging = true
	}

	level, err := strconv.Atoi(strings.TrimSpace(string(cap)))
	return level, charging, err
}
