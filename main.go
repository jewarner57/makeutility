package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("tesseract", "./eurotext.png", "-")
	res, _ := cmd.Output()
	fmt.Print(string(res))
}
