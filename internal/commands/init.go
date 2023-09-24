package commands

import (
	"fmt"

	"github.com/rafaelmgr12/git-clone/internal/core"
	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new fit repository",
	Run: func(cmd *cobra.Command, args []string) {
		repoPath := ".fit"
		repo := core.NewRepository(repoPath)
		err := repo.Init()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		fmt.Println("Initialized empty fit repository in .fit")
	},
}
