package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
			// 实际上是一个位操作是: 67108864|134217728
			// refer: https://man7.org/linux/man-pages/man2/clone.2.html
			syscall.CLONE_NEWIPC |
			// 运行后看到当前pid为1
			// # echo $$
			// 1
			syscall.CLONE_NEWPID,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
