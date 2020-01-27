package yml

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	tests := []struct {
		description    string
		filename       string
		expectedResult map[interface{}]interface{}
		expectedError  bool
	}{
		{
			description:    "non-existing file",
			filename:       "./nonexisting.file",
			expectedResult: nil,
			expectedError:  true,
		},
		{
			description:    "non yaml file",
			filename:       "../../test/noo.txt",
			expectedResult: nil,
			expectedError:  true,
		},
		{
			description: "yaml file",
			filename:    "../../test/foo.yml",
			expectedResult: map[interface{}]interface{}{
				"foo":  "bar",
				"some": "randomkey",
				"this": map[interface{}]interface{}{
					"is":     "a",
					"nested": "object",
				},
			},
			expectedError: false,
		},
	}

	for _, test := range tests {
		result, err := ReadFile(test.filename)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if err != nil {
			continue
		}

		assert.Truef(
			t,
			reflect.DeepEqual(test.expectedResult, result),
			test.description,
		)
	}
}

func TestFindValuesByKeyRegex(t *testing.T) {
	tests := []struct {
		description    string
		data           map[interface{}]interface{}
		expression     *regexp.Regexp
		expectedResult [][]byte
	}{
		{
			description:    "empty data map",
			data:           map[interface{}]interface{}{},
			expression:     nil,
			expectedResult: [][]byte{},
		},
		{
			description: "one matching key",
			data: map[interface{}]interface{}{
				"foo": "bar",
			},
			expression: regexp.MustCompile("f"),
			expectedResult: [][]byte{
				[]byte("foo: bar\n"),
			},
		},
		{
			description: "key is no string",
			data: map[interface{}]interface{}{
				0: "bar",
			},
			expression:     regexp.MustCompile("f"),
			expectedResult: [][]byte{},
		},
		{
			description: "expression does not match",
			data: map[interface{}]interface{}{
				"foo": "bar",
			},
			expression:     regexp.MustCompile("a"),
			expectedResult: [][]byte{},
		},
		{
			description: "nested object",
			data: map[interface{}]interface{}{
				"foo": map[interface{}]interface{}{
					"bar": "works",
				},
			},
			expression: regexp.MustCompile("f"),
			expectedResult: [][]byte{
				[]byte("foo:\n  bar: works\n"),
			},
		},
	}

	for _, test := range tests {
		result := FindValuesByKeyRegex(
			test.data,
			test.expression,
		)

		assert.Truef(
			t,
			reflect.DeepEqual(test.expectedResult, result),
			test.description,
		)
	}
}
