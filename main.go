package main

import (
	"GoInternshipPart2/datamanager"
	"fmt"
)

func main() {
	location := "https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=prettyjson&sole"
	filesLocation := "jsonfiles"
	noOfRecords := 100

	err := datamanager.Controller(location, noOfRecords, filesLocation,
		datamanager.GetData, datamanager.RemoveDuplicates, datamanager.GroupByFirstLetter, datamanager.WriteToFiles)

	if err != nil {
		fmt.Printf("unable to run program, reason: %v", err)
		return
	}
	fmt.Println("All good")

}
