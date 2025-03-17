package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type ResolvedExecutable struct {
	Name string
	Path string
}

var ErrNotFound = errors.New("executable not found")

func ResolveExecutable(name string) (ResolvedExecutable, error) {
	searchPaths := parsePathEnv()
	for _, path := range searchPaths {
		resolved, err := resolveFromPath(name, path)
		if err != nil {
			if errors.Is(err, ErrNotFound) {
				continue
			}
			return ResolvedExecutable{}, err
		}

		return resolved, nil
	}

	return ResolvedExecutable{}, ErrNotFound
}

func parsePathEnv() []string {
	pathEnv := os.Getenv("PATH")
	return strings.Split(pathEnv, ":")
}

func resolveFromPath(toResolve string, path string) (ResolvedExecutable, error) {
	fullPath := filepath.Join(path, toResolve)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return ResolvedExecutable{}, ErrNotFound
	}

	return ResolvedExecutable{Name: toResolve, Path: fullPath}, nil
}
