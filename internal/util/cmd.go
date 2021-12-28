package util

import (
	"os/exec"
)

func Execute(cmd ...string) error {
	for _, c := range cmd {
		cmder := exec.Command("/bin/bash", "-c", c)
		_, err := cmder.Output()
		if err != nil {
			return err
		}
	}
	return nil
}
