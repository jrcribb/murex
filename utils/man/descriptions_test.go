//go:build !js && !windows
// +build !js,!windows

package man

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestParseLineFlags(t *testing.T) {
	tests := []string{
		`-a`,
		`--foobar`,
		`-a, --foobar`,
		`-f FILE`,
		`-f FILE, --file=FILE`,
		`-e PATTERNS, --regexp=PATTERNS`,
		`-E, --extended-regexp`,
		`--exclude-from=FILE`,
		`--backup[=CONTROL]`,
		`-R, -r, --recursive`,
		`--list-cmds=group[,group...]`,
		`--exec-path[=<path>]`,
		`--config-env=<name>=<envvar>`,
		`--[no]-help`,
		`--help Output a usage message and exit.`,
	}

	length := len(tests)
	for i := 0; i < length; i++ {
		tests = append(tests, tests[i]+" An autogenerated description")
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		pl := parseLineFlags([]byte(test))
		switch {
		case pl.Position == len(test):
			// success
		case pl.Description != "":
			// success
		default:
			t.Errorf("Could not match %s in test %d: len(test)==%d, result==%d", test, i, len(test), pl.Position)
			if pl.Position < len(test) {
				t.Logf("  scanned so far: '%s'", test[:pl.Position])
			}
		}
	}
}

/*func TestParseDescriptionsLines(t *testing.T) {
	count.Tests(t, 1)

	files, err := manPages.ReadDir(".")
	if err != nil {
		t.Error(err.Error())
	}

	for _, entry := range files {
		file, err := manPages.Open(entry.Name())
		if err != nil {
			t.Errorf("%s: %s", entry.Name(), err.Error())
		}

		gz, err := gzip.NewReader(file)
		if err != nil {
			t.Errorf("%s: %s", entry.Name(), err.Error())
		}

		descriptions := make(map[string]string)
		parseDescriptionsLines(streams.NewReader(gz), &descriptions)
		if len(descriptions) == 0 {
			t.Errorf("%d descriptions returned for '%s'", len(descriptions), entry.Name())
		}

		gz.Close()
		file.Close()
	}
}
*/
