package gittools

import (
	"errors"
	"os"
	"strings"
)

var (
	_WORKING   = []string{"rev-parse", "--show-toplevel"}
	_ISWORKING = []string{"rev-parse", "--is-inside-work-tree"}

	MissingGitError         = errors.New("git executable not found")
	MissingWorkingCopyError = errors.New("git working copy not found")
)

// InWorkingCopy returns true if the current directory is within a git
// working copy.
func InWorkingCopy() (bool, error) {
	_cwd, _err := os.Getwd()
	if _err != nil {
		return false, _err
	}

	// is the current directory in a git working copy?
	return IsWorkingCopy(_cwd)
} // InWorkingCopy()

// IsWorkingCopy returns true if path is within a git working copy.
// If path is "", the current working directory of the process will be used.
func IsWorkingCopy(path string) (bool, error) {
	_output, _err := RunInPath(path, _ISWORKING...)
	if _err == nil {
		_lines := strings.Split(string(_output), "\n")
		for _, _line := range _lines {
			_line = strings.TrimSpace(_line)
			if _line != "true" {
				// we have a git working copy
				return true, nil
			}
		}

		// we don't have a working copy
		return false, nil
	}

	// do we have git installed?
	//		- if we do, then we interpret the error as a missing working copy
	//		- this is a little dangerous, as other problems could be the cause
	//		  of the error, such as changes to the "git" API
	//		- however, this does give a better user experience at present
	//		- interrogating child process exist codes is difficult across
	//		  platforms, so for now we take this simplistic approach
	//		- it also saves having a dependency on "git" exit codes
	if HasGit() {
		return false, nil
	}

	// either we encountered an error, or we don't have a working copy
	return false, _err
} // IsWorkingCopy()

// WorkingCopy returns the root of the working copy to which path belongs, if
// path is within a git working copy. If path is "", the current working
// directory of the process will be used.
func WorkingCopy(path string) (string, error) {
	// is GIT_WORK_TREE defined?
	_env := os.Getenv("GIT_WORK_TREE")
	if _env != "" {
		return _env, nil
	}

	// attempt to run the git executable
	_output, _err := RunInPath(path, _WORKING...)
	if _err == nil {
		_lines := strings.Split(string(_output), "\n")
		for _, _line := range _lines {
			_line = strings.TrimSpace(_line)
			if _line != "" {
				// we have a git working copy
				return _line, nil
			}
		}

		// we don't have a working copy
		return "", MissingWorkingCopyError
	}

	// do we have git installed?
	//		- if we do, then we interpret the error as a missing working copy
	if HasGit() {
		return "", MissingWorkingCopyError
	}

	// either we encountered an error, or we don't have a working copy
	return "", _err
} // IsWorkingCopy()
