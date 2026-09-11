package ecore

import (
	"fmt"
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

// OneGlob() returs only one filepath if matches requirement after glob (*.some)
func OneGlob(input string) (string, error) {
	matches, err := filepath.Glob(input)
	if err != nil {
		panic(err)
	}

	if len(matches) != 1 {
		return "", fmt.Errorf("expected 1 match not %d in %q", len(matches), input)
	}

	return matches[0], nil
}
