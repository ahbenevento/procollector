package main

import (
	"maps"
	"slices"
)

//  //  //

type projectCollection struct {
	projects []project
	tags     map[string][]int
}

func (pc *projectCollection) setProjects(projects []project) *projectCollection {
	pc.projects = projects

	for ip, project := range pc.projects {
		if len(project.tags) > 0 {
			for _, tag := range project.tags {
				if collectionTag, ok := pc.tags[tag]; ok {
					collectionTag = append(collectionTag, ip)
				} else {
					pc.tags[tag] = []int{ip}
				}
			}
		}
	}

	return pc
}

func (pc projectCollection) getAllTags() []string {
	return slices.Collect(maps.Keys(pc.tags))
}

func (pc projectCollection) findTagByProjectItem(tag string, projectItem int) bool {
	if tags, ok := pc.tags[tag]; ok {
		return slices.Contains(tags, projectItem)
	}

	return false
}

//  //  //

func newProjectCollection() *projectCollection {
	return &projectCollection{
		tags: make(map[string][]int),
	}
}
