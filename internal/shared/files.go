package shared

import (
	"path/filepath"
	"runtime"
)

// This is a way to get the root of our project
// Handy for processes that needs to use a file or other resources
var (
	_, b, _, _ = runtime.Caller(0)
	Root       = filepath.Join(filepath.Dir(b), "../..")
)
