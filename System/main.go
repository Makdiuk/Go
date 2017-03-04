package main

import (
	"os"
	"syscall"
)

func main() {

	ret, _, err := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	if err != 0 {
		os.Exit(2)
	}

	if ret > 0 {
		os.Exit(0)
	}
}
