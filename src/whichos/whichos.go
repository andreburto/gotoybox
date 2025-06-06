package main

import (
	"fmt"
	// "os/exec"
	"runtime"
)

func main() {
	os := runtime.GOOS
	// TODO: Test on Windows
	fmt.Printf("Your operating system is: %s\n", os)
}
