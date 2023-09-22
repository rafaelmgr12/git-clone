package infra

import (
	"fmt"
	"io/ioutil"
	"os"
)

func CreateDit(path string) error {

	if err := os.Mkdir(path, 0755); err != nil {
		return fmt.Errorf("erro ao criar diretório %s: %v", path, err)
	}
	return nil
}

func ExitsDir(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func ReadFile(path string) ([]byte, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler arquivo %s: %v", path, err)
	}
	return content, nil
}

func WriteFile(path string, data []byte) error {
	if err := ioutil.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("erro ao escrever no arquivo %s: %v", path, err)
	}
	return nil
}

func CopyFile(src, dest string) error {
	data, err := ReadFile(src)
	if err != nil {
		return err
	}
	return WriteFile(dest, data)
}

func RemoveFile(path string) error {
	if err := os.Remove(path); err != nil {
		return fmt.Errorf("erro ao remover arquivo %s: %v", path, err)
	}
	return nil
}

func RemoveDir(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return fmt.Errorf("erro ao remover diretório %s: %v", path, err)
	}
	return nil
}
