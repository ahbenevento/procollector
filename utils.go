package main

import (
	"slices"
	"strings"
)

//  //  //

func addUniqueString(list *[]string, filenames ...string) {
	for _, filename := range filenames {
		if !slices.Contains(*list, filename) {
			*list = append(*list, filename)
		}
	}
}

func getSanitizedPathList(paths []string) []string {
	if len(paths) == 1 {
		return paths
	}

	result := []string{}

	for _, path := range paths {
		if path == "/" {
			continue
		}

		add := true

		for k, parent := range result {
			if parent == path {
				continue
			}

			if strings.HasPrefix(path, parent) {
				add = false

				break
			} else if strings.HasPrefix(parent, path) {
				result = slices.Delete(result, k, k+1)

				break
			}
		}

		if add {
			addUniqueString(&result, path)
		}
	}

	return result
}
