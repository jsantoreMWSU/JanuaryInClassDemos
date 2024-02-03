package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
	"fmt"
)

func main() {
	binary, lookerr := exec.LookPath("ls")
	if lookerr != nil {
		log.Fatal("couldn't find commad", lookerr)
	}

	args := []string{"ls", "-a", "-l", "-h"}
	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		log.Fatal("Help! no ls", execErr)
	}
	fmt.Println("GoodBye!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
}
