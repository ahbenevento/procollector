package main

import (
	"path"

	"gopkg.in/ini.v1"
)

//  //  //

func loadProjectFromIniFile(filename string, includeDisabledProject bool) (*project, error) {
	iniFile, err := ini.Load(filename)

	if err != nil {
		return nil, err
	}

	section := iniFile.Section("")
	name := section.Key("name").String()
	disabled, _ := section.Key("disabled").Bool()

	if (disabled && !includeDisabledProject) || name == "" {
		// No es un archivo de configuración de proyectos
		return nil, nil
	}

	result := newProject(name, path.Dir(filename))
	result.Tags = section.Key("tag").Strings(",")

	if disabled {
		result.Enabled = false
	}

	return result, nil
}
