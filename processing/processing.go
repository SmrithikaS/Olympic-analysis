package processing

import (
	"archive/zip"
	"encoding/csv"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Unzip(src string, dest string) ([]string, error) {
	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}
		defer rc.Close()

		path := filepath.Join(dest, f.Name)
		filenames = append(filenames, path)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return filenames, err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, rc)
			if err != nil {
				return filenames, err
			}
		}
	}
	return filenames, nil
}

func ReadCSV(filepath string) ([][]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func CleanData(records [][]string) [][]string {
	data := [][]string{}
	for _, record := range records {
		cleanData := []string{}
		for _, field := range record {
			cleanedField := strings.TrimSpace(field)

			if cleanedField == "" {
				cleanedField = "-1"
			}
			cleanedField = strings.ReplaceAll(cleanedField, `"`, `'`)

			cleanedField = strings.ToLower(cleanedField)

			cleanData = append(cleanData, cleanedField)
		}

		data = append(data, cleanData)
	}
	return data
}
