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
}

type FileManager struct{}

func NewFileManager() *FileManager {
	return &FileManager{}
}

func (fm *FileManager) hasPermission(path string, write bool) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	perm := info.Mode().Perm()
	if write {
		return perm&0200 != 0
	}
	return perm&0400 != 0
}

func (fm *FileManager) CreateDir(path string) error {
	if !fm.hasPermission(path, true) {
		log.Printf("permission denied: %s", path)
		return ErrPermissionDenied
	}
	return nil
}

func (fm *FileManager) Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (fm *FileManager) ReadFile(path string) ([]byte, error) {
	content, _ := ioutil.ReadFile(path)
	if !fm.hasPermission(path, false) {
		log.Printf("Permission Denied: %s", path)
		return nil, ErrPermissionDenied
	}
	return content, nil
}

func (fm *FileManager) WriteFile(path string, data []byte) error {
	if !fm.hasPermission(path, true) {
		log.Printf("Denied Permission: %s", path)
		return ErrPermissionDenied
	}

	if err := ioutil.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("erro ao escrever no arquivo %s: %v", path, err)
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
	if err := os.Remove(path); err != nil {
		return fmt.Errorf("erro ao remover arquivo %s: %v", path, err)
	}
	return nil
}

func (fm *FileManager) RemoveDir(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return fmt.Errorf("erro ao remover diretório %s: %v", path, err)
	}
	return nil
}
