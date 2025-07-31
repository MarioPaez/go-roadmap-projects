package utils

import "fmt"

func CommandUnknown(cmd1, cmd2 string) error {
	return fmt.Errorf("command unknown. please see '%s %s' for more information", cmd1, cmd2)
}
