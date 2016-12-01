package gittools_test

import (
	"regexp"
	"testing"

	"github.com/denormal/go-gittools"
)

var _VERSION *regexp.Regexp

func TestVersion(t *testing.T) {
	// if we don't have git installed, then skip this test
	if !gittools.HasGit() {
		t.Skip("git not installed")
	}

	// otherwise, attempt to retrieve the version
	_version, _err := gittools.Version()
	if _err != nil {
		t.Fatalf("unexpected Version() error: %s", _err.Error())
	} else if _version == "" {
		t.Fatalf("unexpected empty git version: %q", _version)
	} else if !_VERSION.Match([]byte(_version)) {
		t.Fatalf(
			"unexpected version; expected dotted-decimal, got %q",
			_version,
		)
	}
} // TestVersion()

func init() {
	_VERSION = regexp.MustCompile("^\\d+(\\.\\d+)*$")
} // init()
