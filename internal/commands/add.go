// commands/add.go

package commands

import (
	"fmt"

	"github.com/rafaelmgr12/git-clone/internal/core"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "add file to fit staging area",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repoPath := ".fit"
		repo := core.NewRepository(repoPath)

		err := repo.AddFiles(args)
		if err != nil {
			fmt.Printf("Cannot add files : %s\n", err)
			return
		}
	},
}
