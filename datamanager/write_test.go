package datamanager

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteToFiles(t *testing.T) {
	// Test writing multiple files
	err := os.Mkdir("testdata", 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
	}

	groupedData := [][][]string{
		{
			{"Anne", "Dare", "Anne.Dare@shea.name", "30305 Homenick Center", "February 8, 2016", "$6,457.46"},
		},
		{
			{"Enrico", "Jacobs", "Enrico.Jacobs@alberta.biz", "53033 Schaefer Summit", "January 30, 2020", "$3,079.66"},
			{"Estella", "Treutel", "whiterabbit42@gmail.com", "1060 Rippin Shoal", "December 19, 2018", "$1,196.73"},
		},
		{
			{"Milan", "Zemlak", "greyfrog78@gmail.com", "3262 Ivy Brook", "December 13, 2018", "$8,494.38"},
		},
	}
	err = WriteToFiles(groupedData, "testdata")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Verify files were created
	for _, letter := range []string{"A", "E", "M"} {
		filename := filepath.Join("testdata", letter+".json")
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			t.Errorf("Expected file %s to exist, but it does not", filename)
		}
	}

	defer func() {
		err := os.RemoveAll("testdata")
		if err != nil {
			t.Errorf("Failed to remove test files: %v", err)
		}
	}()
}
