package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

//  //  //

type projectFinder struct {
	mu sync.Mutex

	sanitizedPaths []string
	tags           directoryTagsCmdParam
	patterns       filePatternsCmdParam
	err            error
	projects       []project
	all            bool
}

func (pf *projectFinder) includeDisabledProjects(include bool) {
	pf.all = include
}

func (pf *projectFinder) printError(path string, err error) {
	pf.mu.Lock()
	defer pf.mu.Unlock()

	fmt.Fprintf(os.Stderr, "* Error al intentar acceder a \"%s\" (ignorando):\n\t%s\n", path, err)
}

func (pf *projectFinder) checkProjectFile(path string) bool {
	pf.mu.Lock()
	defer pf.mu.Unlock()

	project, err := loadProjectFromIniFile(path, pf.all)

	if project != nil {
		if len(pf.tags) > 0 {
			for tag, subdir := range pf.tags {
				if strings.Contains(path, subdir) && !project.hasTag(tag) {
					project.Tags = append(project.Tags, tag)
				}
			}
		}

		pf.projects = append(pf.projects, *project)

		return true
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "* Fallo al leer el archivo INI \"%s\":\n\t%s\n", path, err)
	}

	return false
}

func (pf *projectFinder) Error() error {
	return pf.err
}

func (pf *projectFinder) run() int {
	var wg sync.WaitGroup

	wg.Add(len(pf.sanitizedPaths))

	for _, repo := range pf.sanitizedPaths {
		go func() {
			defer wg.Done()

			ffinder := newFilesFinder(repo, pf.patterns, pf.checkProjectFile)

			ffinder.setErrorCallback(pf.printError)

			if err := ffinder.find(); err != nil && pf.err == nil {
				pf.mu.Lock()
				defer pf.mu.Unlock()

				pf.err = err
			}
		}()
	}

	wg.Wait()
	return len(pf.projects)
}

//  //  //

func newProjectFinder(sanitizedPaths []string, tags directoryTagsCmdParam, patterns filePatternsCmdParam) *projectFinder {
	return &projectFinder{
		sanitizedPaths: sanitizedPaths,
		tags:           tags,
		patterns:       patterns,
	}
}
