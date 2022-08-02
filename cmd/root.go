/*
Copyright © 2022 Darkhound-org

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var version = "20220802"
var rootCmd = &cobra.Command{
	Use:     "IMDBot",
	Version: version,
	Short:   "Movie and series search bot",
	Long: `IMDBot can search information about your favourite movies and tv series right from the terminal.
	`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
