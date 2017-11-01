package cli

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Main is the entry point for the cli pkg.
func Main() {
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "help")
	}

	errCode := 0

	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "Failed running %q\n", os.Args[1])
		errCode = 1
		if ec, ok := errors.Cause(err).(*cliError); ok {
			errCode = ec.exitCode
		}
	}

	os.Exit(errCode)
}

type cliError struct {
	exitCode int
	cause    error
}

func (e *cliError) Error() string {
	return e.cause.Error()
}

var mainCmd = &cobra.Command{
	Use:          "smoke [command] (flags)",
	Short:        "Smoke CLI",
	Long:         "Smoke CLI",
	SilenceUsage: true,
}

func init() {
	cobra.EnableCommandSorting = false

	mainCmd.SetFlagErrorFunc(func(c *cobra.Command, err error) error {
		if err := c.Usage(); err != nil {
			return err
		}

		fmt.Fprintln(c.OutOrStderr())
		return err
	})

	mainCmd.AddCommand(
		initCmd,
		serverCmd,
	)
}

func run(args []string) error {
	mainCmd.SetArgs(args)
	return mainCmd.Execute()
}
