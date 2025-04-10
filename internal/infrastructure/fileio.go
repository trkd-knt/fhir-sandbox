// infrastructure/fileio.go
package infrastructure

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strings"
)

// ReadFile reads a file and returns its content as []byte
func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// WriteFile writes []byte content to a file
func WriteFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

// ReadCSVFile reads a CSV file and returns headers and records
func ReadCSVFile(path string) ([]string, [][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1 // 可変列数でも許容
	headers, err := r.Read()
	if err != nil {
		return nil, nil, err
	}

	var records [][]string
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, err
		}
		records = append(records, rec)
	}
	return headers, records, nil
}

// WriteCSVFile writes headers and records to a CSV file
func WriteCSVFile(path string, headers []string, records [][]string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	if err := w.Write(headers); err != nil {
		return err
	}
	for _, rec := range records {
		if len(rec) != len(headers) {
			return errors.New("record length mismatch with headers")
		}
		if err := w.Write(rec); err != nil {
			return err
		}
	}
	w.Flush()
	return w.Error()
}

// SanitizeHeader makes a CSV header safe (optional helper)
func SanitizeHeader(field string) string {
	return strings.ReplaceAll(field, "-", "_")
}
