package commands

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

type Command struct {
	Name string
	Description string
	ID CommandID
}

type CommandID int

const (
	Install CommandID = iota
)

func newCommand(name string, description string, id CommandID) *Command {
	c := Command{Name: name, Description: description, ID: id}

	return &c
}

func GetCommands() []*Command {
	return []*Command{
		newCommand("install", "installs", Install),
	}
}

func RunCommand(command CommandID, args []string) {
	switch(command){
		case Install:
			install(args)
		default:
			fmt.Println("UNKNOWN COMMAND")
	}
}

func install(args []string) {
	if(len(args) <= 0){
		fmt.Println("No install target given, cancelling")

		return
	}

	var packages string = ""

	for _, arg := range args {
		packages = packages + arg
	}

	fmt.Println("running emerge -av --pretend " + packages)

	var portageArgs string = "emerge -av --pretend " + packages
	cmd := exec.Command("emerge", strings.Split(portageArgs, " ")...)

	stderr, _ := cmd.StderrPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	cmd.Wait()
}
