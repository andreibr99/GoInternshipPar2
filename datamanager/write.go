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

// WriteToFiles receives a 3D slice of strings with the first field being the index for
// each group of records and a location where the files will be created.
// It writes one json file for each group in the slice and returns an error if there is one.
func WriteToFiles(groupedData [][][]string, filesLocation string) error {
	for _, records := range groupedData {
		letter := records[0][0][0:1]
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
		//Store JSON data
		fileData := make(map[string]interface{})
		fileData["index"] = letter
		fileData["records"] = newRecords
		fileData["total_records"] = len(newRecords)

		//Convert the data to JSON
		jsonData, err := json.MarshalIndent(fileData, "", "  ")
		if err != nil {
			return err
		}

		//Write the JSON data to a file
		filename := filepath.Join(filesLocation, letter+".json")
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
