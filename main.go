package main

import (
	"GoInternshipPart2/http"
	"fmt"
)

func main() {
	location := "https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=prettyjson&sole"

	data, err := http.GetData(location, 740)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, v := range data {
		fmt.Printf("%v. -> %v\n", i+1, v)
	}

}
