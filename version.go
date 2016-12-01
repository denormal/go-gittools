package gittools

import (
	"regexp"
)

// declare the regular expression for matching version strings
var _VERSION *regexp.Regexp

// Version returns the version string for the installed git executable,
// or an error if this cannot be determined.
func Version() (string, error) {
	// attempt to extract the version of the current git executable
	//		- we don't need to be in a particular directory for this
	//		  so default to the current directory
	_bytes, _err := Run("--version")
	if _err != nil {
		return "", _err
	}

	// attempt to parse the version string
	//		- we're looking for a dotted sequence of numbers
	return _VERSION.FindString(string(_bytes)), nil
} // Version()

func init() {
	// compile the regular expression pattern
	//		- we're looking for a.c.b numbers
	//		- this may need to be expanded to include modifiers (e.g. "-dev")
	_VERSION = regexp.MustCompile("\\d+(\\.\\d+)*")
} // init()
