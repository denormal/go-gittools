package gittools_test

import (
	"os"
	"strconv"
	"testing"
	"time"

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

	// ensure the environment isn't set
	_env := os.Getenv("GIT_DIR")
	_err = os.Setenv("GIT_DIR", "")
	if _err != nil {
		t.Fatalf("unable to set GIT_DIR environment variable: %s", _err.Error())
	}
	defer func() {
		if _env == "" {
			os.Unsetenv("GIT_DIR")
		} else {
			os.Setenv("GIT_DIR", _env)
		}
	}()

	// using the current working directory or "" should give the same results
	for _, _path := range []string{_cwd, ""} {
		// are we in a working copy?
		_is, _err := gittools.IsWorkingCopy(_path)
		if _err != nil {
			if _err != gittools.MissingWorkingCopyError {
				t.Fatalf("unexpected error: %s", _err.Error())
			}
		}

		if _is {
			// we're in a working copy so ensure we get a non-empty response
			_dir, _err := gittools.GitDir(_path)
			if _err != nil {
				t.Errorf("%q: unexpected error: %s", _path, _err.Error())
			}
			if _dir == "" {
				t.Errorf(
					"%q: expected to be in working copy; none found",
					_path,
				)
			}
		} else {
			// we're not in a working copy so ensure we get an empty response
			_dir, _err := gittools.GitDir(_path)
			if _err != nil {
				if _err != gittools.MissingWorkingCopyError {
					t.Errorf("%q: unexpected error: %s", _path, _err.Error())
				}
			}
			if _dir != "" {
				t.Errorf(
					"%q: expected not to be in working copy; found %q",
					_path, _dir,
				)
			}
		}
	}

	// if we manually set the GIT_DIR environment variable, do we get the
	// expected result from GitDir()?
	_value := strconv.FormatInt(time.Now().UnixNano(), 16)
	_err = os.Setenv("GIT_DIR", _value)
	if _err != nil {
		t.Fatalf("unable to set GIT_DIR environment variable: %s", _err.Error())
	}
	_dir, _err := gittools.GitDir("")
	if _err != nil {
		t.Fatalf("unexpected error: %s", _err.Error())
	} else if _dir != _value {
		t.Fatalf(
			"git directory mismatch; expected %q, got %q",
			_value, _dir,
		)
	}
} // TestGitDir()
