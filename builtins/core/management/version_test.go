package management_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestVersion(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `version`,
			Stdout: `Murex: [0-9]+\.[0-9]+\.[0-9]+( (.*?))?\nBuilt: .*?\nLicense: .*?\nCopyright: 2018-[0-9]{4} Laurence Morgan\n`,
		},
		{
			Block:  `version: --copyright`,
			Stdout: `2018-[0-9]{4} Laurence Morgan\n`,
		},
		{
			Block:  `version: --license`,
			Stdout: `^GPL v2\n$`,
		},
		{
			Block:  `version: --no-app-name`,
			Stdout: `[0-9]+\.[0-9]+\.[0-9]+( (.*?))?\n`,
		},
		{
			Block:  `version: --short`,
			Stdout: `[0-9]+\.[0-9]+`,
		},
		{
			Block:   `version: --sdfsadf`,
			Stderr:  `not a valid parameter`,
			ExitNum: 1,
		},
		{
			Block:   `version: sdfsadf`,
			Stderr:  `not a valid parameter`,
			ExitNum: 1,
		},
	}

	test.RunMurexTestsRx(tests, t)
}
