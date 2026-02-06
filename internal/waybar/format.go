package waybar

import "fmt"

func format(sec int) string {
	return fmt.Sprintf("%02d:%02d", sec/60, sec%60)
}
