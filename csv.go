package main

import (
	"encoding/csv"
	"os"
)

//  //  //

func writeCSVFile(collection projectCollection, filename string) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	csvWriter := csv.NewWriter(file)

	defer csvWriter.Flush()

	tags := collection.getAllTags()
	tagsCount := len(tags)
	record := make([]string, tagsCount+2)

	for _, project := range collection.projects {
		record[0] = project.name
		record[1] = project.path

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
