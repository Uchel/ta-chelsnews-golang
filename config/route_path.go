package config

import (
	"path/filepath"
	"runtime"
)

//Get current file  full path  from  runtime

var (
	_, b, _, _ = runtime.Caller(0)

	ProjectRoutePath = filepath.Join(filepath.Dir(b), "../")
)
