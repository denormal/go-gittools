package gittools_test

import (
	"io/ioutil"
	"os"
	"strconv"
	"testing"
	"time"

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
	if _err != nil {
		t.Errorf("%q: unexpected error: %s", _dir, _err.Error())
	} else if _is {
		t.Errorf(
			"%q: expected not to be in working copy; working copy found",
			_dir,
		)
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
	// if we have GIT_WORK_TREE defined, make sure we clear the environment
	// for the following tests
	//		- we will test its effect later on
	_env := os.Getenv("GIT_WORK_TREE")
	if _env != "" {
		os.Unsetenv("GIT_WORK_TREE")
		defer os.Setenv("GIT_WORK_TREE", _env)
	}

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

	// if we set the environment variable GT_WORK_TREE, do we get the value
	// back as from WorkingCopy()?
	_value := strconv.FormatInt(time.Now().UnixNano(), 16)
	_err = os.Setenv("GIT_WORK_TREE", _value)
	if _err != nil {
		t.Fatalf(
			"unable to set GIT_WORK_TREE environment variable: %s",
			_err.Error(),
		)
	}
	_wc, _err := gittools.WorkingCopy("")
	if _err != nil {
		t.Fatalf("unexpected error: %s", _err.Error())
	} else if _wc != _value {
		t.Fatalf("working copy mismatch; expected %q, got %q", _value, _wc)
	}
} // TestWorkingCopy()
