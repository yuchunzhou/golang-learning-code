package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func main() {
	mod_name := os.Args[1]
	if err := os.Mkdir(mod_name, 0755); err != nil {
		fmt.Println(err.Error())
		return
	}

	if _, err := os.Create(path.Join(mod_name, "main.go")); err != nil {
		fmt.Println(err.Error())
		return
	}

	cmd := exec.Command("go", "mod", "init", mod_name)
	cmd.Dir = mod_name

	if output, err := cmd.CombinedOutput(); err == nil {
		fmt.Println(string(output))
	} else {
		fmt.Println(err.Error())
	}
}
