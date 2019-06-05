package main

import (
	"github.com/MwlLj/go-math/topsis"

	"encoding/json"
	"io/ioutil"
)

func readFile(fileName string) (report *topsis.Report, err error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &report)
	return
}
