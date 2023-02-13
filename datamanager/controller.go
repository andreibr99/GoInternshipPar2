package datamanager

import "fmt"

// Controller orchestrates the process of reading, filtering, grouping and
// writing data.
// It takes 7 parameters:
//
//	location: the location of the data to be read.
//	noOfRecords: the number of records to be read.
//	filesLocation: the location to store the grouped data.
//	reader: a function that reads data from the location and returns a 2D slice of strings and an error if there is one.
//	filter: a function that filters the records and returns a filtered 2D slice of strings.
//	grouper: a function that groups the filtered data and returns a 3D slice of strings.
//	writer: a function that writes the grouped data into separate files and returns an error if there is one.
//
// The function returns an error if any.
func Controller(
	location string,
	noOfRecords int,
	filesLocation string,
	reader func(location string, noOfRecords int) ([][]string, error),
	filter func(records [][]string) [][]string,
	grouper func(records [][]string) [][][]string,
	writer func(groupedData [][][]string, filesLocation string) error,
) error {
	data, err := reader(location, noOfRecords)
	if err != nil {
		fmt.Printf("unable to read the data, reason: %v\n", err)
		return err
	}
	fmt.Printf("successfully read %v records\n", len(data))

	filteredData := filter(data)
	fmt.Println("deleted duplicates")

	groupedData := grouper(filteredData)
	fmt.Println("grouped data by first name initial")

	err = writer(groupedData, filesLocation)
	if err != nil {
		fmt.Printf("unable to write into files, reason: %v\n", err)
		return err
	}
	fmt.Println("successfully added the data into separate files")

	return nil
}
