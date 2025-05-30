package main

import (
	"encoding/json"
	"io"
	"os"
)

//  //  //

func writeJSONFile(collection projectCollection, filename string) error {
	var (
		file   *os.File
		output io.Writer
		err    error
	)

	if filename != "-" {
		file, err = os.Create(filename)

		if err != nil {
			return err
		}

		defer file.Close()

		output = file
	} else {
		output = os.Stdout
	}

	jsonFile := json.NewEncoder(output)

	jsonFile.SetIndent("", "\t")
	return jsonFile.Encode(collection.projects)
}
