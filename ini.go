package main

import (
	"path"

	"gopkg.in/ini.v1"
)

//  //  //

func loadProjectFromIniFile(filename string) (*project, error) {
	iniFile, err := ini.Load(filename)

	if err != nil {
		return nil, err
	}

	section := iniFile.Section("")
	name := section.Key("name").String()

	if name == "" {
		// No es un archivo de configuración de proyectos
		return nil, nil
	}

	result := project{
		name: name,
		path: path.Dir(filename),
		tags: section.Key("tag").Strings(","),
	}

	return &result, nil
}
