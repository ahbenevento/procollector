package main

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
)

//  //  //

func writeCSVFile(collection projectCollection, filename string) error {
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

	csvWriter := csv.NewWriter(output)

	defer csvWriter.Flush()

	tags := collection.getAllTags()
	tagsCount := len(tags)
	record := make([]string, tagsCount+2)

	// Encabezado
	record[0] = "NAME"
	record[1] = "PATH"

	for i, tag := range tags {
		record[i+2] = strings.ToUpper(tag)
	}

	if err = csvWriter.Write(record); err != nil {
		return err
	}

	for _, project := range collection.projects {
		record[0] = project.Name
		record[1] = project.Path

		for i, tag := range tags {
			if project.hasTag(tag) {
				record[i+2] = tag
			} else {
				record[i+2] = ""
			}
		}

		if err = csvWriter.Write(record); err != nil {
			return err
		}
	}

	return nil
}
