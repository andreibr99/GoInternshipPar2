package datamanager

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

var filesLocation = "../jsonfiles"
var noOfRecords = 75

var errReader = errors.New("fake reader")
var errWriter = errors.New("fake writer")

func fakeReader(string, int) ([][]string, error) {
	return nil, errReader
}

func fakeWriter([][][]string, string) error {
	return errWriter
}

func TestController_Run(t *testing.T) {
	// Create a fake server for test data
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"results": [
			{
				"first": "Jewel",
      			"last": "Bernhard",
      			"email": "Jewel.Bernhard@nichole.net",
      			"address": "438 Berniece Shoals",
				"created": "May 8, 2017",
      			"balance": "$841.45"
			},
			{
				"first": "Leo",
      			"last": "Koelpin",
      			"email": "blackrabbit35@gmail.com",
      			"address": "1809 Demetrius Mills",
      			"created": "July 31, 2014",
      			"balance": "$4,994.16"
			},
			{
				"first": "Gideon",
      			"last": "Rau",
      			"email": "plumturtle25@gmail.com",
      			"address": "18733 Cleo Extensions",
      			"created": "April 18, 2016",
      			"balance": "$2,777.56"
			}
		]}`))
		if err != nil {
			return
		}
	}))
	defer server.Close()

	location := server.URL

	type testCase struct {
		name  string
		input error
		want  error
	}

	var tests = []testCase{
		{
			name:  "fake reader",
			input: NewController(fakeReader, RemoveDuplicates, GroupByFirstLetter, WriteToFiles).Run(location, noOfRecords, filesLocation),
			want:  errReader,
		},
		{
			name:  "fake writer",
			input: NewController(GetData, RemoveDuplicates, GroupByFirstLetter, fakeWriter).Run(location, noOfRecords, filesLocation),
			want:  errWriter,
		},
		{
			name:  "real data, no errors expected",
			input: NewController(GetData, RemoveDuplicates, GroupByFirstLetter, WriteToFiles).Run(location, noOfRecords, filesLocation),
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input
			if got != tt.want {
				t.Errorf("Function result: %v, expected result: %v", got, tt.want)
			}
		})
	}
}
