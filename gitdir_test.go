package gittools_test

import (
	"os"
	"testing"

	"github.com/denormal/go-gittools"
)

func TestGitDir(t *testing.T) {
	// we should be in a working copy
	_cwd, _err := os.Getwd()
	if _err != nil {
		t.Fatalf(
			"unable to determine current working directory: %s",
			_err.Error(),
		)
	}

	// using the current working directory or "" should give the same results
	for _, _path := range []string{_cwd, ""} {
		// only perform this test if this is not a working copy
		_is, _err := gittools.IsWorkingCopy(_path)
		if _err != nil {
			if _err != gittools.MissingWorkingCopyError {
				t.Fatalf("unexpected error: %s", _err.Error())
			}
		}

		if !_is {
			t.Skipf("%q: not a working copy", _path)
		} else {
			_dir, _err := gittools.GitDir(_path)
			if _err != nil {
				t.Errorf("%q: unexpected error: %s", _path, _err.Error())
			} else if _dir == "" {
				t.Errorf(
					"%q: expected to be in working copy; none found",
					_path,
				)
			}
		}
	}
} // TestGitDir()
