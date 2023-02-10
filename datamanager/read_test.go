package datamanager

import (
	"testing"
)

func TestGetDataValidLocation(t *testing.T) {
	location := "https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=prettyjson&sole"
	data, err := GetData(location, 170)
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
	//Check if the length of the output data is equal with the input value
	expectedNoOfRecords := 170
	if len(data) != expectedNoOfRecords {
		t.Errorf("Expected %v records but got %v", expectedNoOfRecords, len(data))
	}
}

func TestGetDataInvalidLocation(t *testing.T) {
	location := "invalid location"
	_, err := GetData(location, 170)

	if err == nil {
		t.Error("Expected error but got nil")
	}
}
