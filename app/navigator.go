package main

import (
	"errors"
	"fmt"
	"os"
	"path"
)

type Navigator struct {
	CurrentDir string
}

var ErrNotADirectory = errors.New("not a directory")

func (n *Navigator) PrintWorkingDirectory() {
	fmt.Println(n.CurrentDir)
}

func (n *Navigator) ChangeDirectory(dir string) error {
	if dir == "" {
		return fmt.Errorf("path is empty")
	}

	if dir == "~" {
		dir = os.Getenv("HOME")
	} else if !path.IsAbs(dir) {
		dir = path.Join(n.CurrentDir, dir)
	}

	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return ErrNotFound
	}
	if err != nil {
		return err
	}

	if !info.IsDir() {
		return ErrNotADirectory
	}

	n.CurrentDir = dir
	return nil
}

func MakeNavigator() *Navigator {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("error getting working directory: %v\n", err)
		os.Exit(1)
	}

	return &Navigator{
		CurrentDir: dir,
	}
}
