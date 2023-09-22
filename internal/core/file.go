package core

import (
	"fmt"
	"io/ioutil"
	"os"
)

type FileManager struct{}

func NewFileManager() *FileManager {
	return &FileManager{}
}

func (fm *FileManager) CreateDir(path string) error {
	if err := os.Mkdir(path, 0755); err != nil {
		return fmt.Errorf("erro ao criar diretório %s: %v", path, err)
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
		return nil, fmt.Errorf("erro ao ler arquivo %s: %v", path, err)
	}
	return content, nil
}

func (fm *FileManager) WriteFile(path string, data []byte) error {
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
