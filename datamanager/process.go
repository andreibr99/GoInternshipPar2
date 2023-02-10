package datamanager

import (
	"strings"
)

func RemoveDuplicates(records [][]string) [][]string {
	occurred := map[string]bool{}
	var result [][]string
	for _, row := range records {
		key := strings.Join(row, ",") // the key is going to be the entire row
		if !occurred[key] {
			occurred[key] = true

			result = append(result, row)
		}
	}
	return result
}

func GroupByFirstLetter(records [][]string) [][][]string {
	data := make([][][]string, 26)
	for i := 0; i < 26; i++ {
		for j, v := range records {
			firstCharacter := records[j][0][0:1]
			if firstCharacter == string(rune(i+65)) {
				data[i] = append(data[i], v)
			}
		}
	}
	// Take only the groups that are not empty
	var nonEmptyGroups [][][]string
	for _, group := range data {
		if len(group) > 0 {
			nonEmptyGroups = append(nonEmptyGroups, group)
		}
	}
	return nonEmptyGroups
}
