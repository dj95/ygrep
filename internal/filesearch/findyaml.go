// Package filesearch Search for different files.
package filesearch

import (
	"os"
	"path/filepath"
	"strings"
)

// FindYAML Find all files with .yml or .yaml suffix in a specified directory.
func FindYAML(path string, recursive bool) ([]string, error) {
	// use another function for reading recursively
	if recursive {
		return findRecursive(path)
	}

	// read the current directory listing
	files, err := os.ReadDir(path)

	if err != nil {
		return []string{}, err
	}

	filenames := []string{}

	// iterate through the files
	for _, f := range files {
        fileInfo, err := f.Info()

        if err != nil {
            continue
        }

		if !isYAMLFile(fileInfo) {
			continue
		}

		filenames = append(
			filenames,
			path+"/"+f.Name(),
		)
	}

	return filenames, nil
}

func findRecursive(path string) ([]string, error) {
	var files []string

	// check whether the path exists or not
	if _, err := os.Stat(path); err != nil {
		return []string{}, err
	}

	// walk theough the files
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		// if the file is not a yaml file...
		if !isYAMLFile(info) {
			// ...return directly
			return nil
		}

		// otherwise save the filename
		files = append(files, path)

		return nil
	})

	return files, err
}

func isYAMLFile(fileInfo os.FileInfo) bool {
	// sort out directories
	if fileInfo.IsDir() {
		return false
	}

	// if the suffix is .yml...
	if strings.HasSuffix(fileInfo.Name(), ".yml") {
		return true
	}

	// ...or .yaml, return true
	if strings.HasSuffix(fileInfo.Name(), ".yaml") {
		return true
	}

	// the file is not a yaml file
	return false
}
