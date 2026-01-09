package main

import (
	"fmt"
	"os"

	"fastflag.com/fastflag/commands"
)

func main() {
	var args []string = os.Args[1:]

	var foundCommand bool = false

	for _, command := range commands.GetCommands() {
		if (command.Name == args[0]) {
			commands.RunCommand(command.ID, args[1:])

			foundCommand = true

			break
		}
	}

	if(!foundCommand) {
		fmt.Println("UNKNOWN COMMAND")
	}
}

