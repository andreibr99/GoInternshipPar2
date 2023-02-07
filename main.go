package main

import (
	"GoInternshipPart2/datamanager"
	"fmt"
)

func main() {
	location := "https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=prettyjson&sole"
	filesLocation := "C:\\Users\\andre\\GolandProjects\\GoInternshipPart2\\jsonfiles"

	data, err := datamanager.GetData(location, 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	data = append(data, data[1])
	for i, v := range data {
		fmt.Printf("%v. -> %v\n", i+1, v)
	}

	/*data2 := [][]string{
		{"mam", "ar"}, {"mam", "ar"}, {"mamm", "arr", "arr"}, {"m u", "4"}, {"m u", "4"}, {"aaa"},
	}*/
	filteredData := datamanager.RemoveDuplicates(data)

	fmt.Println("----------------------------------------------------------")

	for i, v := range filteredData {
		fmt.Printf("%v. -> %v\n", i+1, v)
	}
	fmt.Printf("Records in data: %v\n", len(data))
	fmt.Printf("Records without dups: %v\n", len(filteredData))

	/*groupedData := datamanager.GroupByFirstLetter2(filteredData)

	for i, v := range groupedData {
		fmt.Printf("%v:\n", string(i))
		fmt.Println(len(v))
		for j, k := range v {
			fmt.Printf("%v. -> %v\n", j+1, k)
		}
	}

	err = datamanager.WriteToFiles(groupedData, filesLocation)
	if err != nil {
		fmt.Println(err)
		return
	}*/

	groupedData := datamanager.GroupByFirstLetter(filteredData)
	for i, v := range groupedData {
		fmt.Printf("%v:\n", string(rune(i+65)))
		fmt.Println(len(v))
		for j, k := range v {
			fmt.Printf("%v. -> %v\n", j+1, k)
		}
	}
	err = datamanager.WriteToFiles(groupedData, filesLocation)

}
