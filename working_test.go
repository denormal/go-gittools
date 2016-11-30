package gittools_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/denormal/go-gittools"
)

func TestIsWorkingCopy(t *testing.T) {
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
		_is, _err := gittools.IsWorkingCopy(_path)
		if _err != nil {
			if _err != gittools.MissingWorkingCopyError {
				t.Errorf("%q: unexpected error: %s", _path, _err.Error())
			}
		} else if !_is {
			t.Errorf("%q: expected to be in working copy; none found", _path)
		}
	}

	// create a temporary directory
	//		- this should not be a working copy
	_dir, _err := ioutil.TempDir("", "")
	if _err != nil {
		t.Fatalf("unable to create temporary directory: %s", _err.Error())
	}
	defer os.RemoveAll(_dir)

	// ensure this is not a working directory
	_is, _err := gittools.IsWorkingCopy(_dir)
	if _err == nil {
		t.Errorf(
			"%q: expected error: %s",
			_dir, gittools.MissingWorkingCopyError.Error(),
		)
	} else if _is {
		t.Errorf(
			"%q: expected not to be in working copy; working copy found",
			_dir,
		)
	} else if _err != gittools.MissingWorkingCopyError {
		t.Errorf("%q: unexpected error: %s", _dir, _err.Error())
	}
} // TestIsWorkingCopy()

func TestInWorkingCopy(t *testing.T) {
	// we should be in a working copy
	//		- if we're not, ensure it's correctly reported
	_in, _err := gittools.InWorkingCopy()
	if _err != nil {
		if _err != gittools.MissingWorkingCopyError {
			t.Errorf("unexpected error: %s", _err.Error())
		}
	} else if !_in {
		t.Errorf("expected to be in working copy; none found")
	}
} // TestInWorkingCopy()

func TestWorkingCopy(t *testing.T) {
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
				t.Errorf("unexpected error: %s", _err.Error())
			}
		}

		if !_is {
			t.Skipf("%q: not a working copy", _path)
		} else {
			_wc, _err := gittools.WorkingCopy(_path)
			if _err != nil {
				t.Errorf("%q: unexpected error: %s", _path, _err.Error())
			} else if _wc == "" {
				t.Errorf(
					"%q: expected to be in working copy; none found",
					_path,
				)
			}
		}
	}
} // TestWorkingCopy()
