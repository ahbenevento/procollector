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
	} else if len(params.workingDirectories) == 0 {
		return
	}

	if params.printResume {
		printResume(*params)
		return
	}

	projectCollection := findProjects(*params)

	if projectCollection == nil {
		return
	}

	if params.outputCSVFilename != "" {
		writeCSVFile(*projectCollection, params.outputCSVFilename)
	}

	if params.outputJSONFilename != "" {
		writeJSONFile(*projectCollection, params.outputJSONFilename)
	}

	if params.outputCSVFilename == "" && params.outputJSONFilename == "" {
		printProjects(*projectCollection)
	}
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

	fmt.Println("\n  + Salida:")

	if params.outputCSVFilename != "" {
		if params.outputCSVFilename == "-" {
			fmt.Println("    Pantalla (formato CSV)")
		} else {
			fmt.Printf("    Archivo CSV: %s\n", params.outputCSVFilename)
		}
	} else if params.outputJSONFilename != "" {
		if params.outputJSONFilename == "-" {
			fmt.Println("    Pantalla (formato JSON)")
		} else {
			fmt.Printf("    Archivo JSON: %s\n", params.outputJSONFilename)
		}
	} else {
		fmt.Println("    Pantalla")
	}
}

func findProjects(params cmdParams) *projectCollection {
	pf := newProjectFinder(params.workingDirectories, params.tags, params.patterns)

	if params.outputJSONFilename != "" {
		pf.includeDisabledProjects(true)
	}

	if len(params.ignoreFolders) > 0 {
		pf.setIgnoreFolders([]string(params.ignoreFolders))
	}

	count := pf.run()

	if err := pf.Error(); err != nil {
		fmt.Printf("Error al buscar proyectos: %s\n", err)
		os.Exit(1)
	} else if count == 0 {
		return nil
	}

	return newProjectCollection().setProjects(pf.projects)
}

func printProjects(collection projectCollection) {
	fmt.Printf("PROYECTOS ENCONTRADOS: %d\n", len(collection.projects))

	for _, project := range collection.projects {
		fmt.Printf("\n  - %-40s  %s\n", project.Name, project.Path)
	}
}
