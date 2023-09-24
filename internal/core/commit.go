package core

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type Commit struct {
	Hash        string
	Message     string
	Date        time.Time
	AuthorName  string
	AuthorEmail string
	Parent      string
	Files       []string
}

func NewCommit(message, authorName, authorEmail string, parent string, files []string) *Commit {
	c := &Commit{
		Message:     message,
		Date:        time.Now(),
		AuthorName:  authorName,
		AuthorEmail: authorEmail,
		Parent:      parent,
		Files:       files,
	}
	c.Hash = c.generateHash()
	return c
}

func (c *Commit) generateHash() string {
	hashData := fmt.Sprintf("%s%s%s%s", c.Message, c.Date, c.AuthorName, c.AuthorEmail)
	hasher := sha1.New()
	hasher.Write([]byte(hashData))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func (c *Commit) Serialize() string {
	files := strings.Join(c.Files, "\n")
	return fmt.Sprintf(
		"Hash: %s\nMessage: %s\nDate: %s\nAuthor: %s <%s>\nParent: %s\nFiles:\n%s",
		c.Hash, c.Message, c.Date, c.AuthorName, c.AuthorEmail, c.Parent, files,
	)
}

func (c *Commit) Save(repo *Repository) error {
	commitPath := repo.Path + "/objects/commits/" + c.Hash
	data := c.Serialize()
	return repo.File.WriteFile(commitPath, []byte(data))
}

func (r *Repository) GetCommits() ([]Commit, error) {
	objectFiles, err := ioutil.ReadDir(r.Path + "/objects/commits/")
	if err != nil {
		return nil, fmt.Errorf("error reading commits directory: %v", err)
	}

	var commits []Commit
	for _, objectFile := range objectFiles {
		commitContent, err := r.File.ReadFile(r.Path + "/objects/commits/" + objectFile.Name())
		if err != nil {
			return nil, fmt.Errorf("error reading commit object: %v", err)
		}

		commit := DeserializeCommit(string(commitContent))
		commits = append(commits, commit)
	}

	return commits, nil
}

func DeserializeCommit(data string) Commit {
	lines := strings.Split(data, "\n")

	var commit Commit
	for _, line := range lines {
		switch {
		case strings.HasPrefix(line, "Hash:"):
			commit.Hash = strings.TrimPrefix(line, "Hash: ")
		case strings.HasPrefix(line, "Message:"):
			commit.Message = strings.TrimPrefix(line, "Message: ")
		case strings.HasPrefix(line, "Date:"):
			date, _ := time.Parse(time.RFC3339, strings.TrimPrefix(line, "Date: "))
			commit.Date = date
		case strings.HasPrefix(line, "Author:"):
			authorData := strings.TrimPrefix(line, "Author: ")
			parts := strings.SplitN(authorData, " <", 2)
			commit.AuthorName = parts[0]
			commit.AuthorEmail = strings.TrimRight(parts[1], ">")
		case strings.HasPrefix(line, "Parent:"):
			commit.Parent = strings.TrimPrefix(line, "Parent: ")
		case strings.HasPrefix(line, "Files:"):
			i := strings.Index(data, "Files:\n")
			commit.Files = strings.Split(data[i+7:], "\n")
		}
	}
	return commit
}
