package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

// executePython executes a python script passing cmdline args to python
func executePython(response chan [2]string, count int) {
	cmd := exec.Command("python3", "main.py", fmt.Sprintf("%d", count))
	stderr, err := cmd.StderrPipe()
	if err != nil {
		response <- [2]string{"", err.Error()}
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		response <- [2]string{"", err.Error()}
	}
	if err := cmd.Start(); err != nil {
		log.Fatalln("ERROR FROM START", err)
	}
	fmt.Println("PROCESS ID: ", cmd.Process.Pid)
	progError, _ := ioutil.ReadAll(stderr)
	progOutput, _ := ioutil.ReadAll(stdout)

	response <- [2]string{string(progOutput), string(progError)}

	_ = cmd.Wait()

}

func main() {
	response := make(chan [2]string)
	go executePython(response, 30)
	go executePython(response, 40)
	go executePython(response, 20)

	for i := 0; i <= 2; i++ {
		log.Println(<-response)
	}
	fmt.Println("ALL PROGRAM EXECUTED...")
}
