package scripts

import (
	"fluentweb3-go/internal/jsonutil"
	"fluentweb3-go/internal/tpl"
	"io"
	"log"
	"os"
)

type NamedCommand struct {
	Name    string `json:"name "`
	Command string `json:"command"`
}

type NamedCommands struct {
	Commands []NamedCommand `json:"commands"`
}

func ExecShellCommand(cmd string) (int, error) {
	return Exec(cmd).
		FilterScan(func(s string, writer io.Writer) {
			log.Println(s)
		}).Stdout()
}

func ExecShellCommands(jsonFile string, data any) error {
	jsonBytes, _ := os.ReadFile(jsonFile)
	var commands = NamedCommands{}
	jsonutil.ToStruct(string(jsonBytes), &commands)
	return ExecCommands(commands, data)
}

func ExecCommands(commands NamedCommands, data any) error {
	for _, namedCommand := range commands.Commands {
		log.Printf("start execute command: %s,%s", namedCommand.Name, namedCommand.Command)
		output := tpl.RenderTemplate(namedCommand.Command, data)
		_, err := ExecShellCommand(output)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}
