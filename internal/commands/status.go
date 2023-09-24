package commands

import (
	"fmt"

	"github.com/rafaelmgr12/git-clone/internal/core"
	"github.com/spf13/cobra"
)

var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the working tree status",
	Run: func(cmd *cobra.Command, args []string) {
		repoPath := ".fit"
		repo := core.NewRepository(repoPath)

		stagedFiles, err := repo.GetStagedFiles()
		if err != nil {
			fmt.Printf("Error fetching staged files: %s\n", err)
			return
		}

		changesNotStaged, err := repo.GetChangesNotStaged()
		if err != nil {
			fmt.Printf("Error fetching changes not staged: %s\n", err)
			return
		}

		if len(stagedFiles) == 0 && len(changesNotStaged) == 0 {
			fmt.Println("Nothing to commit, working tree clean")
			return
		}

		if len(stagedFiles) > 0 {
			fmt.Println("Changes to be committed:")
			for _, file := range stagedFiles {
				fmt.Printf("\tnew file:   %s\n", file)
			}
		}

		if len(changesNotStaged) > 0 {
			fmt.Println("\nChanges not staged for commit:")
			for _, file := range changesNotStaged {
				fmt.Printf("\tmodified:   %s\n", file)
			}
		}
	},
}
