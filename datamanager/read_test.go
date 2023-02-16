package datamanager

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDataValidLocation(t *testing.T) {
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
		input func() int
		want  int
	}

	var tests = []testCase{
		{
			name: "calling for 1 record",
			input: func() int {
				data, err := GetData(location, 1)
				if err != nil {
					t.Errorf("Expected no error but got %v", err)
				}
				return len(data)
			},
			want: 1,
		},
		{
			name: "calling for the same number of records that are in the body",
			input: func() int {
				data, err := GetData(location, 3)
				if err != nil {
					t.Errorf("Expected no error but got %v", err)
				}
				return len(data)
			},
			want: 3,
		},
		{
			name: "calling more records than the number of records from the body",
			input: func() int {
				data, err := GetData(location, 5)
				if err != nil {
					t.Errorf("Expected no error but got %v", err)
				}
				return len(data)
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input()
			if got != tt.want {
				t.Errorf("Function result: %v, expected result: %v", got, tt.want)
			}
		})
	}
}

func TestGetDataInvalidLocation(t *testing.T) {
	location := "invalid location"
	_, err := GetData(location, 170)

	if err == nil {
		t.Error("Expected error but got nil")
	}
}

func TestGetDataServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	location := server.URL
	_, err := GetData(location, 1)

	if err == nil {
		t.Error("Expected error but got nil")
	}
}

func TestGetDataInvalidResponseBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"invalid": "response"}`))
		if err != nil {
			return
		}
	}))
	defer server.Close()

	location := server.URL
	_, err := GetData(location, 1)

	if err == nil {
		t.Error("Expected error but got nil")
	}
}
