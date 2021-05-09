package virtualterm_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/virtualterm"
)

func TestWriteBasic(t *testing.T) {
	count.Tests(t, 1)

	term := virtualterm.NewTerminal(5, 3)
	test := "Hello world!"
	exp := "Hello\n worl\nd!   \n"

	term.Write([]rune(test))
	act := term.Export()

	if exp != act {
		t.Error("Expected output does not match actual output")
		t.Logf("  Expected: '%s'", exp)
		t.Logf("  Actual:   '%s'", act)
		t.Logf("  exp bytes: %v", []byte(exp))
		t.Logf("  act bytes: %v", []byte(act))
	}
}

func TestWriteNewLineCrLf1(t *testing.T) {
	count.Tests(t, 1)

	term := virtualterm.NewTerminal(5, 3)
	test := "1\n2\n3\n4\n5"
	exp := "   4 \n    5\n     \n"

	term.Write([]rune(test))
	act := term.Export()

	if exp != act {
		t.Error("Expected output does not match actual output")
		t.Logf("  Expected: '%s'", exp)
		t.Logf("  Actual:   '%s'", act)
		t.Logf("  exp bytes: %v", []byte(exp))
		t.Logf("  act bytes: %v", []byte(act))
	}
}

func TestWriteNewLineCrLf2(t *testing.T) {
	count.Tests(t, 1)

	term := virtualterm.NewTerminal(5, 3)
	test := "1\r\n2\r\n3\r\n4\r\n5"
	exp := "3    \n4    \n5    \n"

	term.Write([]rune(test))
	act := term.Export()

	if exp != act {
		t.Error("Expected output does not match actual output")
		t.Logf("  Expected: '%s'", exp)
		t.Logf("  Actual:   '%s'", act)
		t.Logf("  exp bytes: %v", []byte(exp))
		t.Logf("  act bytes: %v", []byte(act))
	}
}

func TestWriteNewLineLf(t *testing.T) {
	count.Tests(t, 1)

	term := virtualterm.NewTerminal(5, 3)
	term.LfIncCr = true
	test := "1\n2\n3\n4\n5"
	exp := "3    \n4    \n5    \n"

	term.Write([]rune(test))
	act := term.Export()

	if exp != act {
		t.Error("Expected output does not match actual output")
		t.Logf("  Expected: '%s'", exp)
		t.Logf("  Actual:   '%s'", act)
		t.Logf("  exp bytes: %v", []byte(exp))
		t.Logf("  act bytes: %v", []byte(act))
	}
}

func TestWriteTwice(t *testing.T) {
	count.Tests(t, 1)

	term := virtualterm.NewTerminal(5, 3)
	term.LfIncCr = true
	w1 := "foo\n"
	w2 := "bar\n"

	term.Write([]rune(w1))
	term.Write([]rune(w2))

	exp := "foo  \nbar  \n     \n"
	act := term.Export()

	if exp != act {
		t.Error("Expected output does not match actual output")
		t.Logf("  Expected: '%s'", exp)
		t.Logf("  Actual:   '%s'", act)
		t.Logf("  exp bytes: %v", []byte(exp))
		t.Logf("  act bytes: %v", []byte(act))
	}
}

func TestWriteSgrBasicExportHtml(t *testing.T) {
	count.Tests(t, 1)

	term := virtualterm.NewTerminal(120, 1)
	term.LfIncCr = true
	test := fmt.Sprintf("Normal%sBold%sUnderscore%sReset", ansi.Bold, ansi.Underscore, ansi.Reset)
	exp := `<span class="">Normal</span><span class="sgr-bold">Bold</span><span class="sgr-bold sgr-underscore">Underscore</span><span class="">Reset</span><span class="">                                                                                               
</span>`

	term.Write([]rune(test))
	act := strings.TrimSpace(term.ExportHtml())

	if exp != act {
		t.Error("Expected output does not match actual output")
		t.Logf("  Expected: '%s'", exp)
		t.Logf("  Actual:   '%s'", act)
		//t.Logf("  exp bytes: %v", []byte(exp))
		//t.Logf("  act bytes: %v", []byte(act))
	}
}
