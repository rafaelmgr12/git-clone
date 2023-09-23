package core

import (
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
