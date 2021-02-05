package app

import (
	"encoding/csv"
	"os"
)

func readCSV(filename string, delimiter rune) ([][]string, []string, error) {

	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return [][]string{}, []string{}, err
	}

	r := csv.NewReader(f)
	r.Comma = delimiter

	// first line is headers
	headers, err := r.Read()
	if err != nil {
		return [][]string{}, []string{}, err
	}

	lines, err := r.ReadAll()

	if err != nil {
		return [][]string{}, []string{}, err
	}

	return lines, headers, nil
}

func testCSV(filename string) error {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return err
	}
	return nil
}
