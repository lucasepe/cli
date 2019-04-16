package cli

import (
	"fmt"
	"io"
	"os"
)

func createHelpCommand(tool *Tool, w io.Writer) *Command {
	action := func() error {
		// The first parameter passed to the help command should be the
		// command for requested documentation.
		if len(os.Args) < 3 {
			tool.printDefaultHelp(w)
		} else {
			cmdName := os.Args[2]
			cmd := tool.getCommand(cmdName)

			if cmd != nil && cmdName != "help" {
				//fmt.Fprintf(w, "%s %s\n%s", tool.binaryName, cmdName, cmd.Documentation)
				fmt.Fprintf(w, "%s", cmd.Documentation)

				if cmd.FlagInit != nil {
					cmd.FlagInit(&cmd.flagSet)

					fmt.Fprintf(w, "\nOptions:\n")

					cmd.flagSet.SetOutput(w)
					cmd.flagSet.PrintDefaults()
				}
			} else {
				tool.printDefaultHelp(w)
			}
		}

		return nil
	}

	cmd := NewCommand("help", "get more information about a command", action)
	return cmd
}
