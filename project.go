package main

import (
	"slices"
)

//  //  //

type project struct {
	Name    string   `json:"name"`
	Path    string   `json:"rootPath"`
	Enabled bool     `json:"enabled"`
	Tags    []string `json:"tags,omitempty"`
}

func (p project) hasTag(name string) bool {
	return slices.Contains(p.Tags, name)
}

//  //  //

func newProject(name, path string) *project {
	return &project{
		Name:    name,
		Path:    path,
		Enabled: true,
	}
}
