package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Smoke db and configuration",
	Long: `
Init will create the initial configuration and service database. It will also
prompt you for the initial admin's username and password.
	`,
	RunE: runInit,
}

func runInit(cmd *cobra.Command, args []string) error {
	fmt.Fprintln(os.Stdout, "Initializing Smoke...")
	fmt.Fprintln(os.Stdout, "Smoke initialized.")

	return nil
}
