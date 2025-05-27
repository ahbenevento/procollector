package main

import (
	"fmt"
	"os"
)

//  //  //

func main() {
	params := newCmdParams()

	if err := params.parse(); err != nil {
		fmt.Printf("Error en los parámetros utilizados: %s\n", err)
		os.Exit(1)
	}

	if params.printResume {
		printResume(*params)
		return
	}

	findProjects(*params)
}

func printResume(params cmdParams) {
	fmt.Println("RESUMEN:")
	fmt.Println("\n  + Buscar proyectos en:")

	for _, path := range params.workingDirectories {
		fmt.Printf("    %s\n", path)
	}

	if len(params.tags) > 0 {
		fmt.Println("\n  + Etiquetar según directorios:")

		for tag, dir := range params.tags {
			fmt.Printf("    %-20s  %s\n", "\""+tag+"\"", dir)
		}
	}

	if len(params.patterns) > 0 {
		fmt.Println("\n  + Archivos con información de proyectos:")

		for _, pattern := range params.patterns {
			fmt.Printf("    %s\n", pattern)
		}
	}
}

func findProjects(params cmdParams) {
	printError := func(path string, err error) {
		fmt.Fprintf(os.Stderr, "* Ignorando \"%s\" por error:\n\t%s\n", path, err)
	}
	checkProjectFile := func(filename string) bool {
		fmt.Println(filename)
		return true
	}

	fmt.Println(getSanitizedPathList(params.workingDirectories))

	for _, repo := range params.workingDirectories {
		ffinder := newFilesFinder(repo, params.patterns, checkProjectFile)

		ffinder.setErrorCallback(printError)

		err := ffinder.find()

		if err != nil {
			fmt.Printf("Error al buscar proyectos: %s\n", err)
			os.Exit(1)
		}
	}
}
