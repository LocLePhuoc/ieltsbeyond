package file

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const MaxFileSize = 1 << 20

type Loader struct {
	root *os.Root
}

var (
	ErrNotFound    = errors.New("content: file not found")
	ErrInvalidPath = errors.New("content: invalid path")
	ErrTooLarge    = errors.New("content: file too large")
)

func NewLoader(dir string) (*Loader, error) {
	root, err := os.OpenRoot(dir)
	if err != nil {
		return nil, fmt.Errorf("content: open root %q: %w", dir, err)
	}
	return &Loader{root: root}, nil
}

func (l *Loader) Close() error {
	return l.root.Close()
}

func (l *Loader) LoadMD(path string) ([]byte, error) {
	path = filepath.Clean(path)
	if path == "." || filepath.IsAbs(path) || !strings.HasSuffix(path, ".md") {
		return nil, fmt.Errorf("%w: %q", ErrInvalidPath, path)
	}

	f, err := l.root.Open(path)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, fmt.Errorf("%w: %q", ErrNotFound, path)
		}
		// error when trying to get outside root folder
		return nil, fmt.Errorf("%w: %q: %v", ErrInvalidPath, path, err)
	}

	defer f.Close()

	data, err := io.ReadAll(io.LimitReader(f, MaxFileSize+1))
	if err != nil {
		return nil, fmt.Errorf("content: read %q: %w", path, err)
	}
	if len(data) > MaxFileSize {
		return nil, fmt.Errorf("%w: %q", ErrTooLarge, path)
	}
	return data, nil
}
