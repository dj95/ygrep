package commandline

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestInitializeFlags(t *testing.T) {
	tests := []struct {
		description          string
		flagName             string
		expectedDefaultValue interface{}
	}{
		{
			description:          "path",
			flagName:             "path",
			expectedDefaultValue: ".",
		},
		{
			description:          "path",
			flagName:             "expression",
			expectedDefaultValue: "",
		},
		{
			description:          "path",
			flagName:             "insensitive",
			expectedDefaultValue: false,
		},
		{
			description:          "path",
			flagName:             "recursive",
			expectedDefaultValue: false,
		},
	}

	InitializeFlags()

	for _, test := range tests {
		value := viper.Get(test.flagName)

		assert.Equalf(t, test.expectedDefaultValue, value, test.description)
	}
}
