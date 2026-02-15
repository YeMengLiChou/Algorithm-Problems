package main

import (
	"flag"
	"io"
)

func GetLeetcodeCommand() *Command {
	return &Command{
		Name: "leetcode",
		Desc: "",
		Run:  RunLeetcode,
	}
}

func RunLeetcode(args []string, errOut io.Writer) error {
	fs := flag.NewFlagSet("leetcode", flag.ContinueOnError)
	fs.SetOutput(errOut)
}
