package app

import (
	"fmt"

	"github.com/mikestefanello/hooks"
)

type Command struct {
	Name        string
	Description string
	Usage       string
	Handler     func([]string)
}

type Script struct {
	commands []Command
}

func NewScript() *Script {
	return &Script{
		commands: []Command{},
	}
}

func (s *Script) AddCommand(c Command) {
	for _, command := range s.commands {
		if command.Name == c.Name {
			panic(fmt.Sprintf("Command %s already exist", c.Name))
		}
	}
	s.commands = append(s.commands, c)
}

func (s *Script) GetCommand(name string) (Command, bool) {
	for _, command := range s.commands {
		if command.Name == name {
			return command, true
		}
	}
	return Command{}, false
}

func (s *Script) Commands() []Command {
	return s.commands
}

var HookBoot = hooks.NewHook[*Script]("boot")
