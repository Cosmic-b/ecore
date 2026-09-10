package ecore

import (
	"os"
	"path/filepath"
)

// scan path dir for format or dirs ".fmt" "dir" returns list[]
func ScanFor(path string, format string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var found []string

	for _, entry := range entries {
		if format == "dir" && entry.IsDir() {
			found = append(found, entry.Name())
			continue
		}

		if !entry.IsDir() && filepath.Ext(entry.Name()) == format {
			found = append(found, entry.Name())
		}
	}	
	return found, nil
}

// ListDiff a[] minus b[]
func ListDiff(a, b []string) []string {
	var diff []string

	exists := make(map[string]bool)

	for _, item := range b {
		exists[item] = true
	}

	for _, item := range a {
		if !exists[item] {
			diff = append(diff, item)
		}
	}

	return diff
}
