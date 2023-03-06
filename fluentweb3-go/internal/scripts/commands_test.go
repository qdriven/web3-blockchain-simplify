package scripts

import "testing"

func TestExecuteCommand(t *testing.T) {
	ExecShellCommands("commands.json", nil)
}
