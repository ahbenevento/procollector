package main

import (
	"fmt"
	"os"
	"sync"
)

//  //  //

type projectFinder struct {
	mu sync.Mutex

	sanitizedPaths []string
	tags           directoryTagsCmdParam
	patterns       filePatternsCmdParam
	err            error
	foundFiles     int
}

func (pf *projectFinder) printError(path string, err error) {
	pf.mu.Lock()
	defer pf.mu.Unlock()

	fmt.Fprintf(os.Stderr, "* Ignorando \"%s\" por error:\n\t%s\n", path, err)
}

func (pf *projectFinder) checkProjectFile(path string) bool {
	pf.mu.Lock()
	defer pf.mu.Unlock()

	// TODO: revisar archivo de proyecto
	fmt.Println(path)

	pf.foundFiles++

	return true
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
	return pf.foundFiles
}

//  //  //

func newProjectFinder(sanitizedPaths []string, tags directoryTagsCmdParam, patterns filePatternsCmdParam) *projectFinder {
	return &projectFinder{
		sanitizedPaths: sanitizedPaths,
		tags:           tags,
		patterns:       patterns,
	}
}
