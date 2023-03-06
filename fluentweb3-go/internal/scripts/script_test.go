package scripts

import (
	"fmt"
	"io"
	"testing"
)

// https://github.com/bitfield/script
func TestExecCommand(t *testing.T) {
	Exec("git -h").FilterScan(func(s string, writer io.Writer) {
		fmt.Println(s)
	}).Stdout()
}

func TestExecCommandChain(t *testing.T) {
	var commands = []string{}
	commands = append(commands, "ls", "git -v", "echo test")
	fmt.Println(commands)
	for i, command := range commands {
		fmt.Println(i)
		Exec(command).FilterScan(func(s string, writer io.Writer) {
			fmt.Println(s)
		}).Stdout()
	}
}

//需要完成内容:
