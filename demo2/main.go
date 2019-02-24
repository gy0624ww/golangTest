package main

import (
	"os/exec"
	"fmt"
)

func main() {
	var (
		cmd *exec.Cmd
		output []byte
		err error
	)
	cmd = exec.Command("/bin/bash", "-c", "sleep 5;ls -a")

	if output, err = cmd.CombinedOutput();err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(output))
}