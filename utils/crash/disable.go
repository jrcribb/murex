//go:build no_crash_handler
// +build no_crash_handler

package crash

/*
	We want tests to panic!

	So this weird source file disables the crash handler for tests
*/

func init() {
	disable_handler = true
}
