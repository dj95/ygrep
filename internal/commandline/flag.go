// Package commandline Specify the commandline handling.
package commandline

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// InitializeFlags Initializes the commandline flags.
func InitializeFlags() {
	// create the commandline flags
	pflag.StringP("path", "p", ".", "choose the path to work in")
	pflag.StringP("expression", "e", "", "the regular expression to use")
	pflag.BoolP("recursive", "r", false, "search for yaml files recursively")
	pflag.BoolP("insensitive", "i", false, "use the expression case insensitive")

	// parse the pflags
	pflag.Parse()

	// bind the pflags
	viper.BindPFlags(pflag.CommandLine)
}
