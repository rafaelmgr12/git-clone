package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	ErrFileNotFound     = fmt.Errorf("file not found")
	ErrPermissionDenied = fmt.Errorf("permission denied")
)

type FileManagerInterface interface {
	Exists(path string) bool
	CreateDir(path string) error
	ReadFile(path string) ([]byte, error)
	WriteFile(path string, data []byte) error
	CopyFile(src, dest string) error
	RemoveFile(path string) error
	RemoveDir(path string) error
	UpdateIndexFile(filename string) error
}

type FileManager struct{}

func NewFileManager() *FileManager {
	return &FileManager{}
}

func (fm *FileManager) CreateDir(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		if os.IsPermission(err) {
			log.Printf("Permission denied: %s", path)
			return ErrPermissionDenied
		}
		return err
	}
	return nil
}

func (fm *FileManager) Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (fm *FileManager) ReadFile(path string) ([]byte, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsPermission(err) {
			log.Printf("Permission Denied: %s", path)
			return nil, ErrPermissionDenied
		}
		return nil, err
	}
	return content, nil
}

func (fm *FileManager) WriteFile(path string, data []byte) error {
	err := ioutil.WriteFile(path, data, 0644)
	if err != nil {
		if os.IsPermission(err) {
			log.Printf("Denied Permission: %s", path)
			return ErrPermissionDenied
		}
		return fmt.Errorf("error writing to file %s: %v", path, err)
	}
	return nil
}

func (fm *FileManager) CopyFile(src, dest string) error {
	data, err := fm.ReadFile(src)
	if err != nil {
		return err
	}
	return fm.WriteFile(dest, data)
}

func (fm *FileManager) RemoveFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		if os.IsPermission(err) {
			log.Printf("Permission Denied: %s", path)
			return ErrPermissionDenied
		}
		return fmt.Errorf("error removing file %s: %v", path, err)
	}
	return nil
}

func (fm *FileManager) RemoveDir(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		if os.IsPermission(err) {
			log.Printf("Permission Denied: %s", path)
			return ErrPermissionDenied
		}
		return fmt.Errorf("error removing directory %s: %v", path, err)
	}
	return nil
}

func (fm *FileManager) UpdateIndexFile(filename string) error {
	indexFilePath := ".fit/index.txt"
	if !fm.Exists(indexFilePath) {
		if err := fm.WriteFile(indexFilePath, []byte("")); err != nil {
			return err
		}
	}
	content, err := fm.ReadFile(indexFilePath)
	if err != nil {
		return err
	}
	content = append(content, []byte(filename+"\n")...)
	return fm.WriteFile(indexFilePath, content)
}
