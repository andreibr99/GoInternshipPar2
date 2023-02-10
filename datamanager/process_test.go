package datamanager

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	type TestCase struct {
		name  string
		input [][]string
		want  [][]string
	}

	var tests = []TestCase{
		{
			name: "when input does not contain duplicates, the output is unchanged",
			input: [][]string{
				{"Milan", "Zemlak", "greyfrog78@gmail.com", "3262 Ivy Brook", "December 13, 2018", "$8,494.38"},
				{"Kailyn", "Mann", "Kailyn.Mann@skye.biz", "6447 Gladyce Courts", "April 4, 2016", "$9,738.54"},
				{"Lorenz", "Gutmann", "salmonsquirrel87@gmail.com", "1068 Hansen Light", "April 29, 2018", "$7,722.35"},
			},
			want: [][]string{
				{"Milan", "Zemlak", "greyfrog78@gmail.com", "3262 Ivy Brook", "December 13, 2018", "$8,494.38"},
				{"Kailyn", "Mann", "Kailyn.Mann@skye.biz", "6447 Gladyce Courts", "April 4, 2016", "$9,738.54"},
				{"Lorenz", "Gutmann", "salmonsquirrel87@gmail.com", "1068 Hansen Light", "April 29, 2018", "$7,722.35"},
			},
		},
		{
			name: "when input contains duplicates, the output is without duplicates",
			input: [][]string{
				{"Milan", "Zemlak", "greyfrog78@gmail.com", "3262 Ivy Brook", "December 13, 2018", "$8,494.38"},
				{"Kailyn", "Mann", "Kailyn.Mann@skye.biz", "6447 Gladyce Courts", "April 4, 2016", "$9,738.54"},
				{"Lorenz", "Gutmann", "salmonsquirrel87@gmail.com", "1068 Hansen Light", "April 29, 2018", "$7,722.35"},
				{"Kailyn", "Mann", "Kailyn.Mann@skye.biz", "6447 Gladyce Courts", "April 4, 2016", "$9,738.54"},
				{"Lorenz", "Gutmann", "salmonsquirrel87@gmail.com", "1068 Hansen Light", "April 29, 2018", "$7,722.35"},
			},
			want: [][]string{
				{"Milan", "Zemlak", "greyfrog78@gmail.com", "3262 Ivy Brook", "December 13, 2018", "$8,494.38"},
				{"Kailyn", "Mann", "Kailyn.Mann@skye.biz", "6447 Gladyce Courts", "April 4, 2016", "$9,738.54"},
				{"Lorenz", "Gutmann", "salmonsquirrel87@gmail.com", "1068 Hansen Light", "April 29, 2018", "$7,722.35"},
			},
		},
		{
			name: "when the input contains just 1 record, the output contains the same 1 record",
			input: [][]string{
				{"Milan", "Zemlak", "greyfrog78@gmail.com", "3262 Ivy Brook", "December 13, 2018", "$8,494.38"},
			},
			want: [][]string{
				{"Milan", "Zemlak", "greyfrog78@gmail.com", "3262 Ivy Brook", "December 13, 2018", "$8,494.38"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveDuplicates(tt.input)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Function result: %v, expected result: %v", got, tt.want)
			}
		})
	}
}

func TestGroupByFirstLetter(t *testing.T) {
	type TestCase struct {
		name  string
		input [][]string
		want  [][][]string
	}

	var tests = []TestCase{
		{
			name: "input with multiple records, output with the records grouped",
			input: [][]string{
				{"Anne", "Dare", "Anne.Dare@shea.name", "30305 Homenick Center", "February 8, 2016", "$6,457.46"},
				{"Enrico", "Jacobs", "Enrico.Jacobs@alberta.biz", "53033 Schaefer Summit", "January 30, 2020", "$3,079.66"},
				{"Milan", "Zemlak", "greyfrog78@gmail.com", "3262 Ivy Brook", "December 13, 2018", "$8,494.38"},
				{"Estella", "Treutel", "whiterabbit42@gmail.com", "1060 Rippin Shoal", "December 19, 2018", "$1,196.73"},
			},
			want: [][][]string{
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
			},
		},
		{
			name: "input contains 1 record",
			input: [][]string{
				{"Anne", "Dare", "Anne.Dare@shea.name", "30305 Homenick Center", "February 8, 2016", "$6,457.46"},
			},
			want: [][][]string{
				{
					{"Anne", "Dare", "Anne.Dare@shea.name", "30305 Homenick Center", "February 8, 2016", "$6,457.46"},
				},
			},
		},
		{
			name: "input with records that all have the same first letter",
			input: [][]string{
				{"Estella", "Treutel", "whiterabbit42@gmail.com", "1060 Rippin Shoal", "December 19, 2018", "$1,196.73"},
				{"Enrico", "Jacobs", "Enrico.Jacobs@alberta.biz", "53033 Schaefer Summit", "January 30, 2020", "$3,079.66"},
				{"Elon", "Musk", "ElonMusk@twitter.com", "27361 San Diego", "September 1, 2022", "$2,372.22"},
			},
			want: [][][]string{
				{
					{"Estella", "Treutel", "whiterabbit42@gmail.com", "1060 Rippin Shoal", "December 19, 2018", "$1,196.73"},
					{"Enrico", "Jacobs", "Enrico.Jacobs@alberta.biz", "53033 Schaefer Summit", "January 30, 2020", "$3,079.66"},
					{"Elon", "Musk", "ElonMusk@twitter.com", "27361 San Diego", "September 1, 2022", "$2,372.22"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupByFirstLetter(tt.input)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Function result: %v, expected result: %v", got, tt.want)
			}
		})
	}
}
