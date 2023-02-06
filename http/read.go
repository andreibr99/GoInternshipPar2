package http

import (
	"encoding/json"
	"fmt"
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
	resp, err := http.Get(location)
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var data Data
	if err = json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
	}

	result := make([][]string, len(data.Results))
	for i, v := range data.Results {
		result[i] = append(result[i], v.First, v.Last, v.Email, v.Address, v.Created, v.Balance)
	}
	return result, err
}

func GetData(location string, noOfRecords int) ([][]string, error) {
	var finalData [][]string
	gotRecords := 0
	for noOfRecords > gotRecords {
		data, err := ReadData(location)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		gotRecords += len(data)
		finalData = append(finalData, data...)
	}
	finalData = finalData[:noOfRecords]

	return finalData, nil
}
