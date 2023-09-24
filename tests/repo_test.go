package test

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/rafaelmgr12/git-clone/internal/core"
)

type MockFileManager struct {
	mockFS map[string][]byte
}

func NewMockFileManager() *MockFileManager {
	return &MockFileManager{
		mockFS: make(map[string][]byte),
	}
}

func (m *MockFileManager) Exists(path string) bool {
	_, exists := m.mockFS[path]
	return exists
}

func (m *MockFileManager) CreateDir(path string) error {
	m.mockFS[path] = []byte{}
	return nil
}

func (m *MockFileManager) ReadFile(path string) ([]byte, error) {
	content, exists := m.mockFS[path]
	if !exists {
		return nil, errors.New("file not found")
	}
	return content, nil
}

func (m *MockFileManager) WriteFile(path string, data []byte) error {
	m.mockFS[path] = data
	return nil
}

func (m *MockFileManager) ListFiles(dir string) ([]string, error) {
	var filenames []string
	for path := range m.mockFS {
		if filepath.Dir(path) == dir {
			filenames = append(filenames, filepath.Base(path))
		}
	}
	return filenames, nil
}

func (m *MockFileManager) CopyFile(src, dest string) error {
	content, exists := m.mockFS[src]
	if !exists {
		return errors.New("source file not found")
	}
	m.mockFS[dest] = content
	return nil
}

func (m *MockFileManager) RemoveFile(path string) error {
	delete(m.mockFS, path)
	return nil
}

func (m *MockFileManager) RemoveDir(path string) error {
	delete(m.mockFS, path)
	return nil
}

func (m *MockFileManager) UpdateIndexFile(filename string) error {
	indexFilePath := ".fit/index.txt"
	content, exists := m.mockFS[indexFilePath]
	if !exists {
		m.mockFS[indexFilePath] = []byte{}
	}
	content = append(content, []byte(filename+"\n")...)
	m.mockFS[indexFilePath] = content
	return nil
}

func TestRepositoryInit(t *testing.T) {
	mockFile := NewMockFileManager()
	repo := &core.Repository{Path: ".fit", File: mockFile}

	err := repo.Init()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestRepositoryAddFiles(t *testing.T) {
	mockFile := NewMockFileManager()
	mockFile.mockFS["file1.txt"] = []byte("content of file1")
	mockFile.mockFS["file2.txt"] = []byte("content of file2")

	repo := &core.Repository{Path: ".fit", File: mockFile}

	err := repo.AddFiles([]string{"file1.txt", "file2.txt"})
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestRepositoryGetStagedFiles(t *testing.T) {
	mockFile := NewMockFileManager()
	mockFile.mockFS[".fit/STAGING"] = []byte("file1.txt\nfile2.txt")

	repo := &core.Repository{Path: ".fit", File: mockFile}

	files, err := repo.GetStagedFiles()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if len(files) == 0 {
		t.Errorf("Expected staged files, but got none")
	}
}

func TestRepositoryGetChangesNotStaged(t *testing.T) {
	mockFile := NewMockFileManager()
	mockFile.mockFS["file3.txt"] = []byte("content of file3")

	repo := &core.Repository{Path: ".fit", File: mockFile}

	changes, err := repo.GetChangesNotStaged()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if len(changes) == 0 {
		t.Errorf("Expected changes not staged, but got none")
	}
}
