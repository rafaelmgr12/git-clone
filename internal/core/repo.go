package core

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"path/filepath"
)

type Repository struct {
	Path string
	File *FileManager
}

const (
	objectsDir = "objects"
	refsDir    = "refs"
	configFile = "config"
)

func NewRepository(path string) *Repository {
	return &Repository{Path: path}
}

func (r *Repository) Init() error {
	if r.File.Exists(r.Path) {
		return errors.New(".fit já existe neste diretório")
	}

	if err := r.File.CreateDir(r.Path); err != nil {
		return err
	}

	for _, dir := range []string{objectsDir, refsDir} {
		if err := r.File.CreateDir(r.Path + "/" + dir); err != nil {
			return err
		}
	}

	return r.createDefaultConfig()

}

func (r *Repository) createDefaultConfig() error {
	defaultConfig := `{
	"author": "",
	"email": ""
}`

	configPath := filepath.Join(r.Path, configFile)
	return r.File.WriteFile(configPath, []byte(defaultConfig))
}

func (r *Repository) AddFile(filePath string) error {
	content, err := r.File.ReadFile(filePath)
	if err != nil {
		return err
	}

	hash := sha1.Sum(content)
	hastToString := hex.EncodeToString(hash[:])

	objectPath := filepath.Join(r.Path, objectsDir, hastToString)
	return r.File.WriteFile(objectPath, content)
}
