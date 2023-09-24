package core

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Repository struct {
	Path string
	File FileManagerInterface
}

const (
	ObjectsDir  = "objects"
	RefsDir     = "refs"
	ConfigFile  = "config"
	StagingArea = "STAGING"
)

func NewRepository(path string) *Repository {
	return &Repository{Path: path, File: NewFileManager()}
}

func (r *Repository) Init() error {
	if r.File.Exists(r.Path) {
		return errors.New(".fit directory already exists")
	}

	if err := r.File.CreateDir(r.Path); err != nil {
		return err
	}

	for _, dir := range []string{ObjectsDir, RefsDir} {
		if err := r.File.CreateDir(r.Path + "/" + dir); err != nil {
			return err
		}
	}

	return r.CreateDefaultConfig()

}

func (r *Repository) CreateDefaultConfig() error {
	defaultConfig := `{
	"author": "",
	"email": ""
}`

	configPath := filepath.Join(r.Path, ConfigFile)
	return r.File.WriteFile(configPath, []byte(defaultConfig))
}

func (r *Repository) AddFiles(filePaths []string) error {
	for _, filePath := range filePaths {
		if err := r.addSingleFile(filePath); err != nil {
			return err
		}
	}
	return r.UpdateStagingArea(filePaths)
}

func (r *Repository) addSingleFile(filePath string) error {
	content, err := r.File.ReadFile(filePath)
	if err != nil {
		return err
	}

	hash := sha1.Sum(content)
	hastToString := hex.EncodeToString(hash[:])

	objectPath := filepath.Join(r.Path, ObjectsDir, hastToString)
	return r.File.WriteFile(objectPath, content)
}

func (r *Repository) GetStagedFiles() ([]string, error) {
	stagingPath := filepath.Join(r.Path, StagingArea)
	if !r.File.Exists(stagingPath) {
		return nil, nil
	}
	content, err := r.File.ReadFile(stagingPath)
	if err != nil {
		return nil, err
	}

	files := strings.Split(string(content), "\n")
	return files, nil
}

func (r *Repository) GetChangesNotStaged() ([]string, error) {
	workingDir := r.Path // Use the repository path as the working directory
	files, err := ioutil.ReadDir(workingDir)
	if err != nil {
		return nil, err
	}

	stagedFiles, err := r.GetStagedFiles()
	if err != nil {
		return nil, err
	}
	var changesNotStaged []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if !contains(stagedFiles, file.Name()) {
			changesNotStaged = append(changesNotStaged, file.Name())
		}
	}
	return changesNotStaged, nil
}

func (r *Repository) UpdateStagingArea(files []string) error {
	stagingPath := filepath.Join(r.Path, StagingArea)
	content := strings.Join(files, "\n")
	return r.File.WriteFile(stagingPath, []byte(content))

}

func contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}
