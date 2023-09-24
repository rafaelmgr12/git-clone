package commands

import (
	"fmt"

	"github.com/rafaelmgr12/git-clone/internal/core"
	"github.com/spf13/cobra"
)

var LogCmd = &cobra.Command{
	Use:   "log",
	Short: "Show commit logs",
	Run: func(cmd *cobra.Command, args []string) {
		repoPath := ".fit"
		repo := core.NewRepository(repoPath)

		commits, err := repo.GetCommits()
		if err != nil {
			fmt.Printf("Error fetching commits: %s\n", err)
			return
		}

		if len(commits) == 0 {
			fmt.Println("No commits found.")
			return
		}

		for _, commit := range commits {
			fmt.Println("------------------------------------------------")
			fmt.Printf("Hash: %s\n", commit.Hash)
			fmt.Printf("Date: %s\n", commit.Date)
			fmt.Printf("Author: %s <%s>\n", commit.AuthorName, commit.AuthorEmail)
			fmt.Printf("Commit: %s\n", commit.Message)
			if commit.Parent != "" {
				fmt.Printf("Parent: %s\n", commit.Parent)
			}
			fmt.Println("------------------------------------------------")
		}
	},
}
