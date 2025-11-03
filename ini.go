package main

import (
	"path"
	"strings"

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

	sections := iniFile.Sections()

	if len(sections) == 1 {
		return result, nil
	}

	for _, section := range sections {
		if !strings.HasPrefix(result.Path, section.Name()) {
			continue
		}

		name := section.Key("name").String()
		disabled, _ := section.Key("disabled").Bool()

		if (disabled && !includeDisabledProject) || name == "" {
			break
		}

		result.Name = name

		if tags := section.Key("tag").Strings(","); len(tags) > 0 {
			result.Tags = tags
		}

		if disabled {
			result.Enabled = false
		}

		break
	}

	return result, nil
}
