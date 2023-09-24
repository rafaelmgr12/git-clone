package core

import (
	"crypto/sha1"
	"fmt"
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
	commitPath := repo.Path + "/objects/" + c.Hash
	data := c.Serialize()
	return repo.File.WriteFile(commitPath, []byte(data))
}
