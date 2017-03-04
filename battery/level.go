package battery

import "io/ioutil"

// Level reads the files written by the kernel to get status and capacity of the battery
func Level(capFile, statusFile string) (string, bool, error) {
	battery := ""
	charging := false
	cap, err := ioutil.ReadFile(capFile)
	if err != nil {
		return battery, charging, err
	}
	status, err := ioutil.ReadFile(statusFile)
	if err != nil {
		return battery, charging, err
	}
	battery, charging = extractLevel(cap, status)
	return battery, charging, err
}

func extractLevel(cap, status []byte) (string, bool) {
	charging := false
	if string(status) == "Charging" {
		charging = true
	}
	return string(cap), charging
}
