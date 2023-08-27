package main

import (
	"fmt"
	"os"
)

// prettierOptions is a struct that contains a series of prettier
// options that the user can customize from the CLI command
type prettierOptions struct {
	WithSemi bool
	TabWidth int
}

// config contains all the configuration settings
// that the CLI program will use (localPath, which prettier extension to use,
// and other options set by the user).
type config struct {
	// target directory where the prettier file will get created
	TargetDirectory string
	// file extension to use (rc, json, js or yaml file)
	FileExtension string
	// prettier options to customize
	PrettierOptions prettierOptions
}

func main() {
	conf, err := parseFlags(os.Stdout, os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	errors := validateConfig(conf)
	if len(errors) > 0 {
		for _, e := range errors {
			fmt.Fprintln(os.Stderr, e)
		}
		os.Exit(1)
	}

	err = createConfigFile(conf)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("prettier file successfully created!")
}
