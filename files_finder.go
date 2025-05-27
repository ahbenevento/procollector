package main

import (
	"os"
	"path"
	"path/filepath"
)

//  //  //

type errorCallback func(string, error)

type foundCallback func(string) bool

type filesFinder struct {
	path     string
	patterns []string
	onError  errorCallback
	onFound  foundCallback
}

func (ff filesFinder) find() error {
	for _, pattern := range ff.patterns {
		files, err := ff.files(ff.path, pattern)

		if err != nil {
			return err
		}

		for _, filename := range files {
			if ff.onFound(filename) {
				return nil
			}
		}
	}

	// Buscar en los subdirectorios
	subdirs, err := ff.dirs(ff.path)

	if err != nil {
		return err
	}

	for _, subdir := range subdirs {
		newff := newFilesFinder(subdir, ff.patterns, ff.onFound)

		if ff.onError != nil {
			newff.setErrorCallback(ff.onError)
		}

		if err := newff.find(); err != nil && ff.onError != nil {
			ff.onError(subdir, err)
		}
	}

	return nil
}

func (ff filesFinder) files(dir, pattern string) ([]string, error) {
	files, err := filepath.Glob(path.Join(dir, pattern))

	if err != nil {
		return nil, err
	}

	result := []string{}

	for _, name := range files {
		finfo, err := os.Stat(name)

		if err != nil {
			return nil, err
		}

		if !finfo.IsDir() && finfo.Size() > 0 {
			result = append(result, name)
		}
	}

	return result, nil
}

func (ff filesFinder) dirs(dir string) ([]string, error) {
	files, err := filepath.Glob(path.Join(dir, "*"))

	if err != nil {
		return nil, err
	}

	result := []string{}

	for _, name := range files {
		finfo, err := os.Stat(name)

		if err != nil {
			return nil, err
		}

		if finfo.IsDir() && finfo.Name()[0] != '.' {
			result = append(result, name)
		}
	}

	return result, nil
}

func (ff *filesFinder) setErrorCallback(callback errorCallback) {
	ff.onError = callback
}

//  //  //

func newFilesFinder(path string, patterns []string, callback foundCallback) *filesFinder {
	return &filesFinder{
		path:     path,
		patterns: patterns,
		onFound:  callback,
	}
}
