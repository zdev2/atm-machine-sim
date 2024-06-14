package helpers

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func WaitForEnter() {
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func Clear(){
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux", "darwin":
		cmd = exec.Command("clear") // Linux or macOS
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls") // Windows
	default:
		fmt.Println("Unsupported platform")
		return
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}