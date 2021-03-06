package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

const ver string = "1.0.0"

func init() {
	var version = &cobra.Command{
		Use:   "version",
		Short: "Show the current version of gcb-visualizer",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("gcb-visualizer v%s\n", ver)
			long, _ := cmd.Flags().GetBool("long")
			if long {
				fmt.Printf("Compiler: %s\n", runtime.Compiler)
				fmt.Printf("Go version: %s\n", runtime.Version())
			}
		},
	}
	version.Flags().BoolP("long", "l", false, "More detailed version")
	rootCmd.AddCommand(version)
}
