package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestParseFlags(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                   string
		args                   []string
		err                    error
		expectedConf           config
		expectedOutputContains string
	}{
		{
			name: "Generate the default config struct when no flags passed",
			args: []string{},
			expectedConf: config{
				TargetDirectory: ".",
				FileExtension:   "rc",
				PrettierOptions: prettierOptions{
					withSemi: true,
					tabWidth: 2,
				},
			},
		},
		{
			name: "generates correct config struct when flags get passed",
			args: []string{"-d", "/test/dir", "-ext", "rc", "-semi-colon=false", "-tab-width", "4"},
			err:  nil,
			expectedConf: config{
				TargetDirectory: "/test/dir",
				FileExtension:   "rc",
				PrettierOptions: prettierOptions{
					withSemi: false,
					tabWidth: 4,
				},
			},
		},
		{
			name: "shows correct help message when h flag gets passed",
			args: []string{"-h"},
			err:  errors.New("flag: help requested"),
			expectedConf: config{
				TargetDirectory: ".",
				FileExtension:   "rc",
				PrettierOptions: prettierOptions{
					withSemi: true,
					tabWidth: 2,
				},
			},
			expectedOutputContains: "Usage of gen_prettier:",
		},
	}

	buff := new(bytes.Buffer)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, err := parseFlags(buff, tc.args)
			if tc.err == nil && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}

			if tc.err != nil {
				if err == nil || err.Error() != tc.err.Error() {
					t.Errorf("expected error %v, got %v", tc.err, err)
				}
			}

			if c != tc.expectedConf {
				t.Errorf("expected %#v, got: %#v", c, tc.expectedConf)
			}

			if len(tc.expectedOutputContains) != 0 {
				gotOutput := buff.String()
				if !strings.Contains(gotOutput, tc.expectedOutputContains) {
					t.Errorf("expected output: %q, got: %q", tc.expectedOutputContains, gotOutput)
				}
			}

			buff.Reset()
		})
	}
}

func TestValidateConfig(t *testing.T) {
	testcases := []struct {
		conf config
		errs []error
	}{
		{
			conf: config{
				FileExtension:   "yaml",
				TargetDirectory: "/dir/path",
				PrettierOptions: prettierOptions{
					tabWidth: 4,
				},
			},
			errs: []error{},
		},
		{
			conf: config{
				PrettierOptions: prettierOptions{
					tabWidth: 14,
				},
				FileExtension: "ts",
			},
			errs: []error{
				errors.New("invalid file extension. supported extensions: rc, json, yaml or js"),
				errors.New("tab width number must be lower or equal to 12"),
			},
		},
	}

	for _, tc := range testcases {
		got := validateConfig(tc.conf)

		if len(tc.errs) == 0 && len(got) > 0 {
			t.Errorf("expected no errors, got %v", got)
		}

		if len(tc.errs) > 0 {
			for i, e := range tc.errs {
				if got[i] == nil || got[i].Error() != e.Error() {
					t.Errorf("expected: %v, got %v", e, got[i])
				}
			}
		}
	}
}
