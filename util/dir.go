package util

import (
	"path/filepath"
	"runtime"
)

// RootDir returns relative path of current project.
func RootDir() string {
	if _, b, _, ok := runtime.Caller(0); ok {
		return filepath.Join(filepath.Dir(b), "..")
	}

	return ""
}
