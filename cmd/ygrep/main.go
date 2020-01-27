package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/dj95/ygrep/internal/commandline"
	"github.com/dj95/ygrep/internal/filesearch"
	"github.com/dj95/ygrep/internal/yml"
)

func init() {
	// initialize the commandline flags
	commandline.InitializeFlags()

	// if the expression is not set...
	if viper.GetString("expression") == "" {
		// ...print the help page...
		pflag.PrintDefaults()

		// ...and exit with an error code
		os.Exit(1)
	}
}

func main() {

	// search for yaml files
	files, err := filesearch.FindYAML(
		viper.GetString("path"),
		viper.GetBool("recursive"),
	)

	if err != nil {
		panic(err)
	}

	expression := viper.GetString("expression")

	if viper.GetBool("insensitive") {
		expression = "(?i)" + expression
	}

	// compile the regex in order to speed up processing the file
	exp, err := regexp.Compile(expression)

	if err != nil {
		panic(err)
	}

	// iterate through all given files
	for _, file := range files {
		// read the file into a document
		doc, err := yml.ReadFile(file)

		if err != nil {
			fmt.Printf("err: %v\n", err)
			continue
		}

		// print all values matched by the given regular expression
		result := yml.FindValuesByKeyRegex(
			doc,
			exp,
		)

		// check for an empty result...
		if len(result) == 0 {
			// ...and stop processing
			continue
		}

		// otherwise print the file name...
		fmt.Printf(":: %s\n", file)

		// ...and the results
		for _, r := range result {
			fmt.Printf("%s", r)
		}

		fmt.Printf("\n")
	}
}
