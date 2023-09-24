package commands

import (
	"fmt"
	"os"

	"github.com/rafaelmgr12/git-clone/internal/core"
	"github.com/spf13/cobra"
)

var message string

var CommitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit changes",
	Long:  `This command stages and commits changes.`,
	Run: func(cmd *cobra.Command, args []string) {
		repoPath := ".fit"
		repo := core.NewRepository(repoPath)

		stagedFiles, err := repo.GetStagedFiles()
		if err != nil {
			fmt.Println("Error fetching staged files:", err)
			return
		}
		if len(stagedFiles) == 0 {
			fmt.Println("No changes to commit.")
			return
		}

		// TODO: You'd want to retrieve these from the user's config or environment
		authorName := "John Doe"
		authorEmail := "john.doe@example.com"

		commit := core.NewCommit(message, authorName, authorEmail, "", stagedFiles) // Assuming no parent for simplicity
		if err := commit.Save(repo); err != nil {
			fmt.Println("Error saving commit:", err)
			return
		}

		// Clear the staging area after a successful commit
		stagingPath := repo.Path + "/" + core.StagingArea
		os.Remove(stagingPath)

		fmt.Printf("Committed with hash: %s\n", commit.Hash)
	},
}

func init() {
	CommitCmd.Flags().StringVarP(&message, "message", "m", "", "Commit message (required)")
	CommitCmd.MarkFlagRequired("message")
}
