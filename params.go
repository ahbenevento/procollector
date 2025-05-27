package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

//  //  //

// Cada una de las etiquetas que pueden definirse mediante los formatos:
// "tag=dir/subdir"
// "tag:dir/subdir"
type direcotryTags map[string]string

// Permite definir una etiqueta y su correspondiente carpeta utilizando los
// separadores "=" o ":".
func (dt direcotryTags) Set(value string) error {
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
func (dt direcotryTags) String() string {
	return ""
}

// Cada uno de los patrones utilizados para identificar archivos con información
// sobre los proyectos.
type filePatterns []string

func (fm *filePatterns) Set(value string) error {
	patterns := strings.Split(value, ":")

	for k := range patterns {
		patterns[k] = strings.TrimSpace(patterns[k])
	}

	*fm = append(*fm, patterns...)

	return nil
}

// TODO: no utilizado
func (fm filePatterns) String() string {
	return ""
}

// Impide que el paquete "flag" muestre los errores al parsear la linea de
// comandos.
type nullWriter struct{}

func (nullWriter) Write([]byte) (int, error) {
	return 0, nil
}

// Estructura utilizada para configurar el funcionamiento de la aplicación.
type params struct {
	workingDirectories []string
	tags               direcotryTags
	patterns           filePatterns
}

// Parse la linea de comandos.
func (p *params) parse() error {
	flag := flag.NewFlagSet("params", flag.ContinueOnError)
	flag.Usage = func() {}
	hideErrors := nullWriter{}

	flag.SetOutput(hideErrors)

	flag.Var(&p.tags, "t", `Define una etiqueta de directorio con el formato: "etiqueta=dir/subdir"`)
	flag.Var(&p.patterns, "f", `Define uno o más nombres de archivos a buscar (separados por ":"`)

	if err := flag.Parse(os.Args[1:]); err != nil {
		return err
	}

	if len(flag.Args()) == 0 {
		return errors.New("por favor ingrese una carpeta donde localizar sus proyectos")
	}

	p.workingDirectories = flag.Args()

	return nil
}

//  //  //

func newParams() *params {
	return &params{
		tags: make(direcotryTags),
	}
}
