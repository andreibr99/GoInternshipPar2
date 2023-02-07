package datamanager

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Record struct {
	First   string `json:"first"`
	Last    string `json:"last"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Created string `json:"created"`
	Balance string `json:"balance"`
}

func WriteToFiles(groupedData [][][]string, dataLocation string) error {
	for i, records := range groupedData {
		letter := i + 65
		newRecords := make([]Record, len(records))

		//Convert the records to data structure
		for j, r := range records {
			newRecords[j] = Record{
				First:   r[0],
				Last:    r[1],
				Email:   r[2],
				Address: r[3],
				Created: r[4],
				Balance: r[5],
			}
		}

		//Skip letters with no records
		if len(newRecords) == 0 {
			continue
		}

		//Store JSON data
		fileData := make(map[string]interface{})
		fileData["index"] = string(rune(letter))
		fileData["records"] = newRecords
		fileData["total_records"] = len(newRecords)

		//Convert the data to JSON
		jsonData, err := json.MarshalIndent(fileData, "", "  ")
		if err != nil {
			return err
		}

		//Write the JSON data to a file
		filename := filepath.Join(dataLocation, string(rune(letter))+".json")
		file, err := os.Create(filename)
		if err != nil {
			return err
		}

		_, err = file.Write(jsonData)
		if err != nil {
			return err
		}

		err = file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
