// Package commandline Specify the commandline handling.
package commandline

import (
	"fmt"

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
	pflag.BoolP("help", "h", false, "Print the help")

	// parse the pflags
	pflag.Parse()

	// bind the pflags
	viper.BindPFlags(pflag.CommandLine)
}

// PrintHelp Print the help and the default tag
func PrintHelp() {
	fmt.Printf("Usage: ygrep [OPTION]... PATTERN [PATH]\n")
	fmt.Printf("Search PATTERN in each yaml file of the PATH\n")
	fmt.Printf("Example: ygrep -ri foo ./test\n")
	fmt.Printf("PATTERN should contain a regular expression that matches the\n")
	fmt.Printf("key(s) to search for.\n\n")
	fmt.Printf("Options:\n")
	pflag.PrintDefaults()
	fmt.Printf("\n")
	fmt.Printf("Report bugs at: https://github.com/dj95/ygrep/issues\n")
	fmt.Printf("Homepage: https://github.com/dj95/ygrep\n")
}
