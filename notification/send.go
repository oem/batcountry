package notification

import "os/exec"

// Send is shelling out to notify-send
func Send(message string, level string) error {
	cmd := exec.Command("/usr/bin/notify-send", message, "-u", level)
	return cmd.Run()
}
