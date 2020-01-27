// Package yml Holds all operations to read and parse yml files.
package yml

import (
	"io/ioutil"
	"regexp"

	"gopkg.in/yaml.v2"
)

// ReadFile Read a file into a yaml document.
func ReadFile(filename string) (map[interface{}]interface{}, error) {
	// read the file
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	// initialize the map
	m := make(map[interface{}]interface{})

	// unmarshal the yaml file
	err = yaml.Unmarshal([]byte(data), &m)

	if err != nil {
		return nil, err
	}

	return m, nil
}

// FindValuesByKeyRegex Prints a value when the key is matched
// by the given regular expression.
func FindValuesByKeyRegex(
	data map[interface{}]interface{},
	expression *regexp.Regexp,
) [][]byte {
	// initialize the result array
	result := [][]byte{}

	// iterate through all keys with their related value
	for key, value := range data {
		// initialize working variables
		var ok bool
		var keyString string

		// try to convert the key to a string
		if keyString, ok = key.(string); !ok {
			continue
		}

		// if the expression does not match the key...
		if !expression.Match([]byte(keyString)) {
			// ...skip to the next one
			continue
		}

		// convert the object back to a yaml string
		// error handling does not seem to work... when an error is
		// produces, e.g. by marshalling an inlined struct to duplicate
		// keys, the program panics and does not return an error...
		valueYML, _ := yaml.Marshal(map[string]interface{}{
			keyString: value,
		})

		// save the result
		result = append(result, valueYML)

		// if the value is another yaml object...
		if v, ok := value.(map[interface{}]interface{}); ok {
			// ...process it as well...
			newResults := FindValuesByKeyRegex(v, expression)

			// ...and save the results
			result = append(result, newResults...)
		}
	}

	return result
}
