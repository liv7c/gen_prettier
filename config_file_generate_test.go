package main

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func Test_createConfigFile(t *testing.T) {
	tempDir := t.TempDir()
	expectedDir := "./testdata"

	testcases := []struct {
		name                string
		conf                config
		generatedFile       string
		expectedFileContent string
	}{
		{
			name: "rc config file",
			conf: config{
				FileExtension:   "rc",
				TargetDirectory: tempDir,
				PrettierOptions: prettierOptions{
					ArrowParens: "always",
					WithSemi:    false,
					TabWidth:    4,
				},
			},
			generatedFile:       filepath.Join(tempDir, ".prettierrc"),
			expectedFileContent: filepath.Join(expectedDir, ".prettierrc"),
		},
		{
			name: "js config file",
			conf: config{
				FileExtension:   "js",
				TargetDirectory: tempDir,
				PrettierOptions: prettierOptions{
					ArrowParens: "always",
					WithSemi:    false,
					TabWidth:    4,
				},
			},
			generatedFile:       filepath.Join(tempDir, ".prettierrc"),
			expectedFileContent: filepath.Join(expectedDir, ".prettierrc"),
		},
		{
			name: "json config file",
			conf: config{
				FileExtension:   "json",
				TargetDirectory: tempDir,
				PrettierOptions: prettierOptions{
					ArrowParens: "always",
					WithSemi:    false,
					TabWidth:    4,
				},
			},
			generatedFile:       filepath.Join(tempDir, ".prettierrc"),
			expectedFileContent: filepath.Join(expectedDir, ".prettierrc"),
		},
		{
			name: "yaml config file",
			conf: config{
				FileExtension:   "yaml",
				TargetDirectory: tempDir,
				PrettierOptions: prettierOptions{
					ArrowParens: "always",
					WithSemi:    true,
					TabWidth:    2,
				},
			},
			generatedFile:       filepath.Join(tempDir, ".prettierrc"),
			expectedFileContent: filepath.Join(expectedDir, ".prettierrc"),
		},
	}

	for _, tc := range testcases {

		err := createConfigFile(tc.conf)
		if err != nil {
			t.Fatalf("expected no error, got non-nil error: %v", err)
		}

		content, err := os.ReadFile(tc.generatedFile)
		if err != nil {
			log.Fatal(err)
		}

		expectedContent, err := os.ReadFile(tc.expectedFileContent)
		if err != nil {
			log.Fatal(err)
		}

		if !bytes.Equal(content, expectedContent) {
			t.Errorf("expected: %q, got: %q", string(expectedContent), string(content))
		}
	}

}
