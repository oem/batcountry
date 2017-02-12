package notification

import "strconv"

// OutputText is creating text with minimal markup that could be used by dunst or something similar
func OutputText(level int) string {
	return "Battery low! <b>" + strconv.Itoa(level) + "%</b>"
}
