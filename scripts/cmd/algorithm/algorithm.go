package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type Command struct {
	Name  string
	Desc  string
	Run   func(args []string) error
}

var help = ""

func main() {
	if err := run(os.Args, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}
}

func run(args []string, out, errOut io.Writer) error {
	if len(args) < 2 {
		return errors.New("missing subcommand.")
	}
	commands := collectCommands()
	command := commands[args[1]]
	if command == nil {
		return errors.New("invalid subcommand.")
	}
}

func collectCommands() map[string]*Command {
	commands := map[string]*Command{}
	
	// env
	commands["env"] = &Command{
		Name: "env",
		Desc: "设置 input 环境变量",
		Run
	}

	// leetcode 
	commands["leetcode"] = GetLeetcodeCommand(errOut io.Writer)

	return commands
}




