package tmp

import (
	"fmt"
	"os"
)

type Tmp struct {
	files []*os.File
	dirs  []string
}

func New() *Tmp {
	return &Tmp{
		files: make([]*os.File, 0),
		dirs:  make([]string, 0),
	}
}

func (t *Tmp) Close() error {
	var e1, e2 error

	for _, file := range t.files {
		if err := os.Remove(file.Name()); err != nil {
			e1 = err
		}
	}

	for _, dir := range t.dirs {
		if err := os.RemoveAll(dir); err != nil {
			e2 = err
		}
	}

	switch {
	case e1 != nil:
		return fmt.Errorf("failed to remove temp file: %w", e1)
	case e2 != nil:
		return fmt.Errorf("failed to remove temp dir: %w", e2)
	default:
		return nil
	}
}

func (t *Tmp) Dir(name string) (dir string, err error) {
	dir, err = os.MkdirTemp("", fmt.Sprintf("*.%s", name))
	if err != nil {
		return dir, fmt.Errorf("failed to create temp dir: %w", err)
	}

	t.dirs = append(t.dirs, dir)
	return dir, nil
}

func (t *Tmp) File(name string) (file *os.File, err error) {
	file, err = os.CreateTemp("", fmt.Sprintf("*.%s", name))
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %w", err)
	}

	t.files = append(t.files, file)
	return file, nil
}

func (t *Tmp) MustDir(name string) string {
	dir, err := t.Dir(name)
	if err != nil {
		panic(err)
	}

	return dir
}

func (t *Tmp) MustFile(name string) *os.File {
	file, err := t.File(name)
	if err != nil {
		panic(err)
	}

	return file
}
