// shellexec.go
package goutils

import (
	"os/exec"
)

func ShellExec(cmdStr string) (err error) {

	cmd := exec.Command("bash", "-c", cmdStr)

	_, err = cmd.Output()
	if err != nil {
		return
	}

	return

}
