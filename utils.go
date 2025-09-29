package main

import (
	"slices"
	"strings"
)

//  //  //

// Retorna una lista ordenada, sin duplicados de las carpetas padres de otras
// indicadas desde la línea de comando.
func getSanitizedPathList(paths []string) []string {
	if len(paths) == 1 {
		return paths
	}

	slices.SortFunc(paths, func(path1, path2 string) int {
		if len(path1) < len(path2) {
			return -1
		}

		return 1
	})

	result := []string{}

	for _, path := range paths {
		if path == "/" {
			continue
		}

		for _, parent := range result {
			if parent == path {
				goto skip
			}

			if strings.HasPrefix(path, parent+"/") {
				goto skip
			}
		}

		result = append(result, path)

	skip:
	}

	return result
}
