package main

import (
	"log"
	"os"
	"os/exec"
)

var cmdChain = []*exec.Cmd{
	exec.Command("synonyms"),
	exec.Command("sprinkle"),
	exec.Command("coolify"),
	exec.Command("domainify"),
	exec.Command("available"),
}

func main() {
	cmdChain[0].Stdin = os.Stdin
	cmdChain[len(cmdChain)-1].Stdout = os.Stdout

	for i := 0; i < len(cmdChain)-1; i++ {
		thisCmd := cmdChain[i]
		nextCmd := cmdChain[i+1]
		stdOut, err := thisCmd.StdoutPipe()
		if err != nil {
			log.Panicln(err)
		}

		nextCmd.Stdin = stdOut
	}

	for _, cmd := range cmdChain {
		if err := cmd.Start(); err != nil {
			log.Panicln(err)
		} else {
			defer cmd.Process.Kill()
		}
	}

	for _, cmd := range cmdChain {
		if err := cmd.Wait(); err != nil {
			log.Panicln(err)
		}
	}

}
