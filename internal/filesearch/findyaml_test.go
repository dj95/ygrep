package filesearch

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/**
 * Mocks for tests
 */
type MockFileInfo struct {
	os.FileInfo

	isDir bool
	name  string
}

func (m *MockFileInfo) Name() string {
	return m.name
}

func (m *MockFileInfo) Size() int64 {
	return 1
}

func (m *MockFileInfo) Mode() os.FileMode {
	return 0644
}

func (m *MockFileInfo) ModTime() time.Time {
	return time.Now()
}

func (m *MockFileInfo) IsDir() bool {
	return m.isDir
}

func (m *MockFileInfo) SysDir() interface{} {
	return nil
}

/**
 * Testing functions
 */

func TestFindYAML(t *testing.T) {
	tests := []struct {
		description    string
		path           string
		recursive      bool
		expectedResult []string
		expectedError  bool
	}{
		{
			description: "search recursive without subfolders",
			path:        "../../test/nestedfolder",
			recursive:   true,
			expectedResult: []string{
				"../../test/nestedfolder/bar.yml",
			},
			expectedError: false,
		},
		{
			description:    "path does not exists",
			path:           "../nonexisting/folder",
			recursive:      true,
			expectedResult: []string{},
			expectedError:  true,
		},
		{
			description: "search recursive with subfolders",
			path:        "../../test/",
			recursive:   true,
			expectedResult: []string{
				"../../test/foo.yml",
				"../../test/nestedfolder/bar.yml",
			},
			expectedError: false,
		},
		{
			description:    "search non-recursive with non-existing path",
			path:           "../nonexisting/folder",
			recursive:      false,
			expectedResult: []string{},
			expectedError:  true,
		},
		{
			description: "search non-recursive with subfolders",
			path:        "../../test",
			recursive:   false,
			expectedResult: []string{
				"../../test/foo.yml",
			},
			expectedError: false,
		},
	}

	for _, test := range tests {
		result, err := FindYAML(test.path, test.recursive)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if err != nil {
			continue
		}

		for i, r := range result {
			assert.Equalf(
				t,
				test.expectedResult[i],
				r,
				test.description,
			)
		}
	}
}

func TestFindRecursive(t *testing.T) {
	tests := []struct {
		description    string
		path           string
		expectedResult []string
		expectedError  bool
	}{
		{
			description: "search recursive without subfolders",
			path:        "../../test/nestedfolder",
			expectedResult: []string{
				"../../test/nestedfolder/bar.yml",
			},
			expectedError: false,
		},
		{
			description:    "path does not exists",
			path:           "../nonexisting/folder",
			expectedResult: []string{},
			expectedError:  true,
		},
		{
			description: "search recursive with subfolders",
			path:        "../../test/",
			expectedResult: []string{
				"../../test/foo.yml",
				"../../test/nestedfolder/bar.yml",
			},
			expectedError: false,
		},
	}

	for _, test := range tests {
		result, err := findRecursive(test.path)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if err != nil {
			continue
		}

		for i, r := range result {
			assert.Equalf(
				t,
				test.expectedResult[i],
				r,
				test.description,
			)
		}
	}
}

func TestIsYAMLFile(t *testing.T) {
	tests := []struct {
		description    string
		fileInfo       os.FileInfo
		expectedResult bool
	}{
		{
			description: "is directory",
			fileInfo: &MockFileInfo{
				isDir: true,
				name:  "example_dir",
			},
			expectedResult: false,
		},
		{
			description: "directory with yaml suffix",
			fileInfo: &MockFileInfo{
				isDir: true,
				name:  "example_dir.yml",
			},
			expectedResult: false,
		},
		{
			description: "file with correct long suffix",
			fileInfo: &MockFileInfo{
				isDir: false,
				name:  "example_file.yaml",
			},
			expectedResult: true,
		},
		{
			description: "file with correct short suffix",
			fileInfo: &MockFileInfo{
				isDir: false,
				name:  "example_file.yml",
			},
			expectedResult: true,
		},
		{
			description: "file with wrong suffix",
			fileInfo: &MockFileInfo{
				isDir: false,
				name:  "example_file.txt",
			},
			expectedResult: false,
		},
	}

	for _, test := range tests {
		result := isYAMLFile(test.fileInfo)

		assert.Equalf(t, test.expectedResult, result, test.description)
	}
}
