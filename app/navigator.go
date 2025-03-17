package main

import (
	"fmt"
	"os"
)

type Navigator struct {
	CurrentDir string
}

func (n *Navigator) PrintWorkingDirectory() {
	fmt.Println(n.CurrentDir)
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
