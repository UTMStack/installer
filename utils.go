package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func runEnvCmd(env []string, command string, arg ...string) error {
	cmd := exec.Command(command, arg...)
	cmd.Env = append(os.Environ(), env...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Println("Executing command:", command, "with args:", arg)
	return cmd.Run()
}

func runCmd(command string, arg ...string) error {
	return runEnvCmd([]string{}, command, arg...)
}

func checkCmd(command string, arg ...string) {
	if err := runCmd(command, arg...); err != nil {
		log.Fatal(err)
	}
}

func checkOutput(command string, arg ...string) string {
	log.Println("Executing command:", command, "with args:", arg)
	out, err := exec.Command(command, arg...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(out))
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}