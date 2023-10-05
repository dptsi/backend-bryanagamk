package main

import (
	"fmt"
	"os"
	"sort"

	"its.ac.id/base-go/script/internal/app"

	// Commands
	_ "its.ac.id/base-go/script/internal/commands/makecontroller"
	_ "its.ac.id/base-go/script/internal/commands/makemodule"
)

func main() {
	args := os.Args[1:]

	s := app.NewScript()
	app.HookBoot.Dispatch(s)

	if len(args) == 0 || args[0] == "help" {
		cmds := s.Commands()
		sort.Slice(cmds, func(i, j int) bool {
			return cmds[i].Name < cmds[j].Name
		})
		fmt.Println("Available commands:")
		fmt.Println()

		for _, cmd := range cmds {
			fmt.Printf("======== %s ========\n", cmd.Name)
			fmt.Printf("Description\t: %s\n", cmd.Name)
			fmt.Printf("Usage\t\t: %s\n", cmd.Usage)
			fmt.Println()
		}
		return
	}

	command, exist := s.GetCommand(args[0])
	if !exist {
		fmt.Printf("Command %s not found\nRun this script without arguments or with argument \"help\" to show help\n", args[0])
		return
	}
	command.Handler(args[1:])
}
