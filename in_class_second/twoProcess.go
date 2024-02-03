package main

import(
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main(){
	currentDir, err := os.Getwd()
	if err != nil{
		os.Exit(1)
	}
	prog := filepath.Join(currentDir, "secondMain")
	command := exec.Command(prog)
	command.Stdout = os.Stdout
	err = command.Run()
	if err != nil{
		log.Fatal("couldn't start process", err)
	
	}
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("====================================================")

	err = command.Wait()
	if err != nil{
	fmt.Println("huh how did that happen:")	
}	


}
