package gittools

import (
	"strings"
)

var _DIR = []string{"rev-parse", "--git-dir"}

// GitDir returns the git directory for the working copy in which path is
// located, or an error if the path cannot be resolved, or is not located
// within a working copy. If path is "", the current working directory of
// the process will be used.
func GitDir(path string) (string, error) {
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

	// we could not determine if we are in a git working copy
	return "", _err
} // GitDir()
