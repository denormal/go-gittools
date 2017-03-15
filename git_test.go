package gittools_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/denormal/go-gittools"
)

func TestGit(t *testing.T) {
	_git, _err := gittools.Git()
	if _err != nil {
		t.Fatalf("git executable not found: %s", _err.Error())
	} else if _git == "" {
		t.Fatalf("git executable found, but path empty: %q", _git)
	}

	// ensure the git path exists and is executable
	_info, _err := os.Stat(_git)
	if _err != nil {
		t.Fatalf("git executable could not be inspected: %q", _err.Error())
	} else if _info.IsDir() {
		t.Fatalf("git path is directory not file: %s", _git)
	} else if _info.Mode()|0100 == 0 {
		t.Fatalf("git path is not executable by caller: %s", _git)
	}

	// amend the current PATH to be empty
	_path := os.Getenv("PATH")
	defer os.Setenv("PATH", _path)
	_err = os.Setenv("PATH", "")
	if _err != nil {
		t.Fatalf("unable to reset PATH: %s", _err.Error())
	}

	// run the test again to ensure Git() fails
	_git, _err = gittools.Git()
	if _err == nil {
		t.Error("expected error looking for git; none found")
	} else {
		_error, _ok := _err.(*exec.Error)
		if !_ok || _error.Err != exec.ErrNotFound {
			t.Errorf("expected ErrNotFound; got %q", _err.Error())
		}
	}

	// ensure the path to git is empty
	if _git != "" {
		t.Errorf("expected empty path from Git(); found %q", _git)
	}
} // TestGit()

func TestHasGit(t *testing.T) {
	// do we have git installed?
	if !gittools.HasGit() {
		t.Skip("git not installed")
	}

	// amend the current PATH to be emtpy
	_path := os.Getenv("PATH")
	defer os.Setenv("PATH", _path)
	_err := os.Setenv("PATH", "")
	if _err != nil {
		t.Fatalf("unable to reset PATH: %s", _err.Error())
	}

	// HasGit should now fail
	if gittools.HasGit() {
		t.Errorf("unexpected success: HasGit() with PATH %q", os.Getenv("PATH"))
	}
} // TestHasGit()
