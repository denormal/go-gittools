package gittools

import (
	"os"
	"os/exec"
	"path/filepath"
)

// Run attempts to execute the git command in the current process working
// directory with args arguments. If git is executed successfully, the slice
// of bytes output to STDOUT by git will be returned, otherwise an error will
// be returned.
func Run(args ...string) ([]byte, error) {
	return RunInPath("", args...)
} // Run()

// RunInPath attempts to execute the git command from the path directory (or
// the parent of path if path represents a file), with args arguments. If
// git is executed successfully, the slice of bytes output to STDOUT by git
// will be returned, otherwise an error will be returned. If path is the
// empty string, the current process working directory will be used.
func RunInPath(path string, args ...string) ([]byte, error) {
	// is git installed?
	_git, _ := Git()
	if _git == "" {
		return nil, MissingGitError
	}

	// if we have a path, attempt to change into it before executing
	// the git command
	if path != "" {
		var _dir string

		// do we have a file or a directory?
		_info, _err := os.Stat(path)
		if _err != nil {
			return nil, _err
		} else if _info.IsDir() {
			_dir = path
		} else {
			_dir, _ = filepath.Split(path)
		}

		_cwd, _err := os.Getwd()
		if _err != nil {
			return nil, _err
		}

		// attempt to change into the given path
		_err = os.Chdir(_dir)
		if _err != nil {
			return nil, _err
		}
		defer os.Chdir(_cwd)
	}

	// execute the git command
	return exec.Command(_git, args...).Output()
} // RunInPath()
