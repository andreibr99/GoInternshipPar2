package datamanager

import "fmt"

func NewController(
	reader func(location string, noOfRecords int) ([][]string, error),
	filter func(records [][]string) [][]string,
	grouper func(records [][]string) [][][]string,
	writer func(groupedData [][][]string, filesLocation string) error,
) Controller {
	return Controller{
		reader:  reader,
		filter:  filter,
		grouper: grouper,
		writer:  writer,
	}
}

type Controller struct {
	reader  func(location string, noOfRecords int) ([][]string, error)
	filter  func(records [][]string) [][]string
	grouper func(records [][]string) [][][]string
	writer  func(groupedData [][][]string, filesLocation string) error
}

func (c Controller) Run(location string, noOfRecords int, filesLocation string) error {
	data, err := c.reader(location, noOfRecords)
	if err != nil {
		fmt.Printf("unable to read the data, reason: %v\n", err)
		return err
	}
	fmt.Printf("successfully read %v records\n", len(data))

	filteredData := c.filter(data)
	fmt.Println("deleted duplicates")

	groupedData := c.grouper(filteredData)
	fmt.Println("grouped data by first name initial")

	err = c.writer(groupedData, filesLocation)
	if err != nil {
		fmt.Printf("unable to write into files, reason: %v\n", err)
		return err
	}
	fmt.Println("successfully added the data into separate files")

	return nil
}
