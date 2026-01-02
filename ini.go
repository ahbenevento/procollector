package main

import (
	"path"
	"slices"
	"strings"

	"gopkg.in/ini.v1"
)

//  //  //

func getProjectNameFromIniSection(section *ini.Section) string {
	return section.Key("name").String()
}

func getProjectRootFromIniSection(section *ini.Section) string {
	if projectPath := section.Key("rootPath").String(); projectPath != "" {
		return projectPath
	}

	return section.Key("root").String()
}

func checkDisabledProjectFromIniSection(section *ini.Section) bool {
	if disabled, _ := section.Key("disabled").Bool(); disabled {
		return true
	}

	return false
}

func getTagsFromIniSection(section *ini.Section) ([]string, bool) {
	tags := section.Key("tag").Strings(",")
	add := false

	if len(tags) > 0 {
		if tags[0] = strings.TrimSpace(tags[0]); tags[0] != "" && tags[0][0] == '+' {
			add = true
			tags[0] = strings.TrimSpace(tags[0][1:])
		}
	}

	return tags, add
}

func updateProjectByIniSection(project *project, section *ini.Section) (updated bool) {
	if disabled := checkDisabledProjectFromIniSection(section); disabled {
		project.Enabled = false
		updated = true
	}

	if name := getProjectNameFromIniSection(section); name != "" {
		project.Name = name
		updated = true
	}

	if tags, add := getTagsFromIniSection(section); len(tags) > 0 {
		if add {
			project.Tags = append(project.Tags, tags...)
		} else {
			project.Tags = tags
		}

		updated = true
	}

	if projectPath := getProjectRootFromIniSection(section); projectPath != "" {
		project.Path = projectPath
		updated = true
	}

	return
}

func loadProjectFromIniFile(filename string, includeDisabledProject bool) (*project, error) {
	iniFile, err := ini.Load(filename)

	if err != nil {
		return nil, err
	}

	section := iniFile.Section("")
	name := getProjectNameFromIniSection(section)
	disabled := checkDisabledProjectFromIniSection(section)

	if (disabled && !includeDisabledProject) || name == "" {
		// No es un archivo de configuración de proyectos
		return nil, nil
	}

	result := newProject(name, path.Dir(filename))
	result.Tags, _ = getTagsFromIniSection(section)

	if disabled {
		result.Enabled = false
	}

	if projectPath := getProjectRootFromIniSection(section); projectPath != "" {
		result.Path = projectPath
	}

	sections := iniFile.Sections()

	if len(sections) == 1 {
		return result, nil
	}

	// Buscar configuraciones especiales según la carpeta donde se encuentre el
	// proyecto
	for _, section := range sections {
		if !strings.ContainsAny(section.Name(), "/\\") || !strings.HasPrefix(result.Path, section.Name()) {
			continue
		}

		updateProjectByIniSection(result, section)
		break
	}

	if !result.Enabled && !includeDisabledProject {
		return nil, nil
	}

	// Buscar una sección que coincida con el usuario@host
	user := getUserAndHost()
	is := slices.IndexFunc(sections, func(s *ini.Section) bool {
		return s.Name() == user
	})

	if is != -1 {
		if updateProjectByIniSection(result, sections[is]) && !result.Enabled && !includeDisabledProject {
			return nil, nil
		}
	} else if hostname := getHostname(); hostname != "" {
		// Buscar una sección que coincida solo con el host
		is = slices.IndexFunc(sections, func(s *ini.Section) bool {
			return s.Name() == hostname
		})

		if is != -1 {
			if updateProjectByIniSection(result, sections[is]) && !result.Enabled && !includeDisabledProject {
				return nil, nil
			}
		}
	}

	return result, nil
}
