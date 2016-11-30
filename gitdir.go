package gittools

import (
	"os"
	"strings"
)

var _DIR = []string{"rev-parse", "--git-dir"}

// GitDir returns the git directory for the working copy in which path is
// located, or an error if the path cannot be resolved, or is not located
// within a working copy. If path is "", the current working directory of
// the process will be used.
func GitDir(path string) (string, error) {
	// do we have the GIT_DIR environment variable set?
	_env := os.Getenv("GIT_DIR")
	if _env != "" {
		return _env, nil
	}

	// attempt to resolve the .git directory within the given path hierarchy
	_output, _err := RunInPath(path, _DIR...)
	if _err == nil {
		_lines := strings.Split(string(_output), "\n")
		for _, _line := range _lines {
			_line = strings.TrimSpace(_line)
			if _line != "" {
				return _line, nil
			}
		}
		return "", MissingWorkingCopyError
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
		return "", MissingWorkingCopyError
	}

	// we could not determine if we are in a git working copy
	return "", _err
} // GitDir()
