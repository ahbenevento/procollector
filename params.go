package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

//  //  //

// La lista de archivos que pueden contener información sobre un proyecto.
const default_projectFilePatterns string = ".project:project.ini"

//  //  //

// Cada una de las etiquetas que pueden definirse mediante los formatos:
// "tag=dir/subdir"
// "tag:dir/subdir"
type directoryTagsCmdParam map[string]string

// Permite definir una etiqueta y su correspondiente carpeta utilizando los
// separadores "=" o ":".
func (dt directoryTagsCmdParam) Set(value string) error {
	parts := 1
	valueParts := strings.FieldsFunc(value, func(r rune) bool {
		if parts == 0 {
			return false
		}

		if r == ':' || r == '=' {
			parts--

			return true
		}

		return false
	})

	if len(valueParts) == 2 {
		dt[strings.TrimSpace(valueParts[0])] = strings.TrimSpace(valueParts[1])

		return nil
	}

	return fmt.Errorf(`formato de etiqueta mal utilizado "%s"`, value)
}

// TODO: no utilizado
func (directoryTagsCmdParam) String() string {
	return ""
}

// Cada uno de los patrones utilizados para identificar archivos con información
// sobre los proyectos.
type filePatternsCmdParam []string

// Permite definir una lista de patrones separados por ":".
func (fm *filePatternsCmdParam) Set(value string) error {
	patterns := strings.Split(value, ":")

	for k := range patterns {
		patterns[k] = strings.TrimSpace(patterns[k])
	}

	*fm = append(*fm, patterns...)

	return nil
}

// TODO: no utilizado
func (filePatternsCmdParam) String() string {
	return ""
}

type directoryList []string

func (dl *directoryList) Set(value string) error {
	*dl = append(*dl, value)

	return nil
}

// TODO:
func (directoryList) String() string {
	return ""
}

// Estructura utilizada para configurar el funcionamiento de la aplicación.
type cmdParams struct {
	workingDirectories []string
	tags               directoryTagsCmdParam
	patterns           filePatternsCmdParam
	printResume        bool
	outputCSVFilename  string
	outputJSONFilename string
	ignoreFolders      directoryList
}

func (p *cmdParams) parse() error {
	flag := flag.NewFlagSet("cmdParams", flag.ContinueOnError)

	flag.SetOutput(io.Discard)
	flag.Var(&p.tags, "t", `Define una etiqueta de directorio con el formato: "etiqueta=dir/subdir"`)
	flag.Var(&p.patterns, "f", `Define uno o más nombres de archivos a buscar (separados por ":").`)
	flag.BoolVar(&p.printResume, "r", false, "Solo muestra un resumen de los parámetros recibidos.")
	flag.StringVar(&p.outputCSVFilename, "csv", "", "Guarda los proyectos encontrados en un archivo CSV.")
	flag.StringVar(&p.outputJSONFilename, "json", "", "Guarda los proyectos en un archivo JSON.")
	flag.Var(&p.ignoreFolders, "i", `Permite ignorar directorios específicos.`)

	errorInParams := flag.Parse(os.Args[1:])

	if errorInParams != nil {
		return errorInParams
	}

	for _, dir := range flag.Args() {
		path, err := filepath.Abs(dir)

		if err != nil {
			return err
		}

		p.workingDirectories = append(p.workingDirectories, path)
	}

	if len(p.workingDirectories) > 0 {
		p.workingDirectories = getSanitizedPathList(p.workingDirectories)
	}

	if len(p.workingDirectories) == 0 && errorInParams == nil {
		flag.SetOutput(os.Stdout)
		flag.PrintDefaults()
	}

	if len(p.patterns) == 0 {
		p.patterns = p.getDefaultFilePatterns()
	}

	return nil
}

func (cmdParams) getDefaultFilePatterns() []string {
	result := filePatternsCmdParam{}

	result.Set(default_projectFilePatterns)

	return result
}

//  //  //

func newCmdParams() *cmdParams {
	return &cmdParams{
		tags: make(directoryTagsCmdParam),
	}
}
