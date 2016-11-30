package gittools

import (
	"os/exec"
	"path/filepath"
)

// Git returns the absolute path to the locally installed git executable, if
// found. Otherwise, Git will return an error.
func Git() (string, error) {
	_path, _err := exec.LookPath("git")
	if _err == nil {
		_path, _err = filepath.Abs(_path)
		if _err == nil {
			return _path, _err
		}
	}

	return "", _err
} // Git()

// HasGit returns true if the host system has git installed and if the
// git executable is located within the current user's PATH.
func HasGit() bool {
	_path, _ := Git()
	if _path != "" {
		return true
	} else {
		return false
	}
} // HasGit()
