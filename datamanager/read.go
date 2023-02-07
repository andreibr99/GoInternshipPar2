package datamanager

import (
	"encoding/json"
	"io"
	"net/http"
)

type Data struct {
	Results []struct {
		First   string `json:"first"`
		Last    string `json:"last"`
		Email   string `json:"email"`
		Address string `json:"address"`
		Created string `json:"created"`
		Balance string `json:"balance"`
	} `json:"results"`
}

func ReadData(location string) ([][]string, error) {
	//Read the JSON data from location
	resp, err := http.Get(location)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//Convert the JSON to data structure
	var data Data
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	//Store the obtained data into a 2D slice
	result := make([][]string, len(data.Results))
	for i, v := range data.Results {
		result[i] = append(result[i], v.First, v.Last, v.Email, v.Address, v.Created, v.Balance)
	}
	return result, nil
}

func GetData(location string, noOfRecords int) ([][]string, error) {
	var finalData [][]string
	//Call the ReadData() as many times as needed to get the number of records specified
	for noOfRecords > len(finalData) {
		data, err := ReadData(location)
		if err != nil {
			return nil, err
		}
		finalData = append(finalData, data...)
	}
	finalData = finalData[:noOfRecords]

	return finalData, nil
}
