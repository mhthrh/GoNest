package address

import (
	"bytes"
	csvFile "common-lib/pkg/util/file"
	"encoding/csv"
)

const (
	path = "/file/countries"
	name = "Countries.csv"
)

type Country struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

func Countries() ([]Country, error) {
	f := csvFile.File{
		Name: name,
		Path: path,
		Data: nil,
	}
	e := f.Read()
	if e != nil {
		return nil, e
	}
	reader := csv.NewReader(bytes.NewReader(f.Data))

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	countries := make([]Country, len(rows))
	for i, row := range rows {
		countries[i] = Country{
			ID:   row[0],
			Name: row[1],
			Code: row[2],
		}
	}
	return countries, nil
}
