package battery

import (
	"os/exec"
	"regexp"
	"strconv"
)

// AcpiLevel shells out to acpi to get the battery level
func AcpiLevel() (int, error) {
	var level int
	cmd := exec.Command("acpi")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return level, err
	}

	level, err = extractAcpiLevel(output)
	return level, err
}

func extractAcpiLevel(output []byte) (int, error) {
	var level int
	rp, err := regexp.Compile("(?s)Battery 1:.*?(\\d+?)%")
	if err != nil {
		return level, err
	}

	bat := rp.FindSubmatch(output)[1]
	level, err = strconv.Atoi(string(bat))
	return level, err
}
