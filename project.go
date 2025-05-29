package main

import "slices"

//  //  //

type project struct {
	name string
	path string
	tags []string
}

func (p *project) addTag(name string) {
	p.tags = append(p.tags, name)
}

func (p project) hasTag(name string) bool {
	return slices.Contains(p.tags, name)
}

//  //  //

func newProject(name, path string) *project {
	return &project{
		name: name,
		path: path,
	}
}
