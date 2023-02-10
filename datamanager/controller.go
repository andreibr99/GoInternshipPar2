package datamanager

import "fmt"

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
