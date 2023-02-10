package datamanager

import (
	"errors"
	"testing"
)

var location = "https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=prettyjson&sole"
var filesLocation = "C:\\Users\\andre\\GolandProjects\\GoInternshipPart2\\jsonfiles"
var noOfRecords = 75

var errReader = errors.New("fake reader")
var errWriter = errors.New("fake writer")

func fakeReader(string, int) ([][]string, error) {
	return nil, errReader
}

func fakeWriter([][][]string, string) error {
	return errWriter
}

func TestController(t *testing.T) {
	type testCase struct {
		name  string
		input error
		want  error
	}

	var tests = []testCase{
		{
			name:  "fake reader",
			input: Controller(location, noOfRecords, filesLocation, fakeReader, RemoveDuplicates, GroupByFirstLetter, WriteToFiles),
			want:  errReader,
		},
		{
			name:  "fake writer",
			input: Controller(location, noOfRecords, filesLocation, GetData, RemoveDuplicates, GroupByFirstLetter, fakeWriter),
			want:  errWriter,
		},
		{
			name:  "real data, no errors expected",
			input: Controller(location, noOfRecords, filesLocation, GetData, RemoveDuplicates, GroupByFirstLetter, WriteToFiles),
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
