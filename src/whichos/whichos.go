package main

import (
	"fmt"
	// "os/exec"
	"runtime"
)

func main() {
	os := runtime.GOOS
	fmt.Printf("Your operating system is: %s\n", os)
}
