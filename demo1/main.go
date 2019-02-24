package main

import (
	"os/exec"
	"fmt"
)

func main() {
	var (
		cmd *exec.Cmd
		err error
	)
	cmd = exec.Command("/bin/bash", "-c", "ls -l")
	err = cmd.Run()
	fmt.Println(err)
}
