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
					SingleQuote: true,
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
					SingleQuote: true,
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
					SingleQuote: true,
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
					SingleQuote: true,
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

func TestCreateConfigFile_DoesNotOverrideExistingPrettierConfig(t *testing.T) {
	tempDir := t.TempDir()
	expectedDir := "./testdata"

	testcase := struct {
		conf                config
		generatedFile       string
		expectedFileContent string
	}{
		conf: config{
			FileExtension:   "rc",
			TargetDirectory: tempDir,
			PrettierOptions: prettierOptions{
				ArrowParens: "always",
				WithSemi:    false,
				SingleQuote: true,
				TabWidth:    4,
			},
		},
		generatedFile:       filepath.Join(tempDir, ".prettierrc"),
		expectedFileContent: filepath.Join(expectedDir, ".prettierrc"),
	}

	err := createConfigFile(testcase.conf)
	if err != nil {
		t.Fatalf("expected no error, got non-nil error: %v", err)
	}

	// try to recreate file
	err = createConfigFile(testcase.conf)
	if err == nil {
		t.Error("expected err, got nil instead")
	}

	// check that prettier config file content is correct
	content, err := os.ReadFile(testcase.generatedFile)
	if err != nil {
		log.Fatal(err)
	}

	expectedContent, err := os.ReadFile(testcase.expectedFileContent)
	if err != nil {
		log.Fatal(err)
	}

	if !bytes.Equal(content, expectedContent) {
		t.Errorf("expected: %q, got: %q", string(expectedContent), string(content))
	}

	if !bytes.Equal(content, expectedContent) {
		t.Errorf("expected %q, got %q instead", expectedContent, content)
	}
}
