package main

import (
	"fmt"
	"os"

	"github.com/rafaelmgr12/git-clone/internal/commands"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "fit"}

	rootCmd.AddCommand(commands.InitCmd)
	rootCmd.AddCommand(commands.AddCmd)
	rootCmd.AddCommand(commands.LogCmd)
	rootCmd.AddCommand(commands.StatusCmd)
	rootCmd.AddCommand(commands.CommitCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
