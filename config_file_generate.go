package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates/prettierrc.go.tmpl
var prettierrcFileTmpl string

//go:embed templates/prettierrc.js.go.tmpl
var prettierJSFileTmpl string

//go:embed templates/prettierrc.yaml.go.tmpl
var prettierYAMLFileTmpl string

//go:embed templates/prettierrc.json.go.tmpl
var prettierJSONFileTmpl string

// createConfigFile creates the prettier config file
func createConfigFile(conf config) error {
	tmplMap := map[string]map[string]string{
		"js": {
			".prettierrc.js": prettierJSFileTmpl,
		},
		"rc": {
			".prettierrc": prettierrcFileTmpl,
		},
		"yaml": {
			".prettierrc.yaml": prettierYAMLFileTmpl,
		},
		"json": {
			".prettierrc.json": prettierJSONFileTmpl,
		},
	}

	if _, ok := tmplMap[conf.FileExtension]; !ok {
		return fmt.Errorf("could not find template for given file extension: %s", conf.FileExtension)
	}

	return renderToFile(conf, tmplMap[conf.FileExtension])
}

// renderToFile takes a config struct and tmplMap and creates the file
// with its different dynamic values coming from the config struct.
func renderToFile(conf config, tmplMap map[string]string) error {
	tmpl := template.New("prettier")

	for filename, t := range tmplMap {
		f, err := os.Create(filepath.Join(conf.TargetDirectory, filename))
		if err != nil {
			return err
		}

		tmpl, err = tmpl.Parse(t)
		if err != nil {
			return err
		}

		err = tmpl.Execute(f, conf)
		if err != nil {
			return err
		}

		f.Close()
	}

	return nil
}
