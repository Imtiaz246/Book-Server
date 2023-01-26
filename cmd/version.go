package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var (
	version    = "v1.0.0"
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Prints the version no of the book-api-server",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("api-server version", version)
		},
	}
)
