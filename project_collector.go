package main

//  //  //

type projectCollector struct {
	repositories []string
	dirTags      directoryTagsCmdParam
	patterns     filePatternsCmdParam
}

func (pc projectCollector) find() {

}

//  //  //

func newProjectCollector(repositories []string, dirTags directoryTagsCmdParam, patterns filePatternsCmdParam) *projectCollector {
	return &projectCollector{
		repositories: repositories,
		dirTags:      dirTags,
		patterns:     patterns,
	}
}
