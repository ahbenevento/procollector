package main

//  //  //

type projectCollection struct {
	projects []project
}

func (pc *projectCollection) add(name, path string) {
	pc.projects = append(pc.projects, *newProject(name, path))
}

func (pc *projectCollection) addProject(p project) {
	pc.projects = append(pc.projects, p)
}

//  //  //

func newProjectCollection() *projectCollection {
	return &projectCollection{}
}
