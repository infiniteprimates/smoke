package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Short: "Start the Smoke service",
	Long: `
Start the smoke service
	`,
	RunE: runServer,
}

func runServer(cmd *cobra.Command, args []string) error {
	fmt.Fprintln(os.Stdout, "Starting Smoke...")
	fmt.Fprintln(os.Stdout, "Smoke started.")

	fmt.Fprintln(os.Stdout, "Smoke stopped.")

	return nil
}