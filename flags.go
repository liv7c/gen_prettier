package main

import (
	"errors"
	"flag"
	"io"
)

// supportedFileExtensions is an array containing all the file
// extensions currently supported by the CLI
var supportedFileExtensions = map[string]bool{
	"rc":   true,
	"js":   true,
	"json": true,
	"yaml": true,
}

// parseFlags parses the args string slice into a config struct.
func parseFlags(w io.Writer, args []string) (config, error) {
	conf := config{}

	fs := flag.NewFlagSet("gen_prettier", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&conf.TargetDirectory, "d", ".", "Target directory")
	fs.StringVar(&conf.FileExtension, "ext", "rc", "File extension for your prettier file. Choose between rc, js, json or yaml")
	fs.IntVar(&conf.PrettierOptions.TabWidth, "tab-width", 2, "Tab width")
	fs.BoolVar(&conf.PrettierOptions.WithSemi, "semi-colon", true, "With or without semi colon")

	err := fs.Parse(args)
	if err != nil {
		return conf, err
	}

	return conf, nil
}

// validateConfig validates that the config struct generated with the CLI flags
// is valid.
func validateConfig(conf config) []error {
	var validationErrors []error

	if _, ok := supportedFileExtensions[conf.FileExtension]; !ok {
		validationErrors = append(validationErrors, errors.New("invalid file extension. supported extensions: rc, json, yaml or js"))
	}

	if conf.PrettierOptions.TabWidth > 12 {
		validationErrors = append(validationErrors, errors.New("tab width number must be lower or equal to 12"))
	}

	return validationErrors
}
