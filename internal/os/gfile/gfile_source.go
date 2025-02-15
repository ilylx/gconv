package gfile

import (
	"github.com/ilylx/gconv/internal/gregex"
	"os"
	"runtime"
	"strings"
)

var (
	// goRootForFilter is used for stack filtering purpose.
	goRootForFilter = runtime.GOROOT()
)

func init() {
	if goRootForFilter != "" {
		goRootForFilter = strings.Replace(goRootForFilter, "\\", "/", -1)
	}
}

// MainPkgPath returns absolute file path of package main,
// which contains the entrance function main.
//
// It's only available in develop environment.
//
// Note1: Only valid for source development environments,
// IE only valid for systems that generate this executable.
//
// Note2: When the method is called for the first time, if it is in an asynchronous goroutine,
// the method may not get the main package path.
func MainPkgPath() string {
	// Only for source development environments.
	if goRootForFilter == "" {
		return ""
	}
	path := mainPkgPath.Val()
	if path != "" {
		return path
	}
	lastFile := ""
	for i := 1; i < 10000; i++ {
		if _, file, _, ok := runtime.Caller(i); ok {
			if goRootForFilter != "" && len(file) >= len(goRootForFilter) && file[0:len(goRootForFilter)] == goRootForFilter {
				continue
			}
			if gregex.IsMatchString(`/github.com/[^/]+/gf/`, file) &&
				!gregex.IsMatchString(`/github.com/[^/]+/gf/\.example/`, file) {
				continue
			}
			if Ext(file) != ".go" {
				continue
			}
			lastFile = file
			if gregex.IsMatchString(`package\s+main`, GetContents(file)) {
				mainPkgPath.Set(Dir(file))
				return Dir(file)
			}
		} else {
			break
		}
	}
	if lastFile != "" {
		for path = Dir(lastFile); len(path) > 1 && Exists(path) && path[len(path)-1] != os.PathSeparator; {
			files, _ := ScanDir(path, "*.go")
			for _, v := range files {
				if gregex.IsMatchString(`package\s+main`, GetContents(v)) {
					mainPkgPath.Set(path)
					return path
				}
			}
			path = Dir(path)
		}
	}
	return ""
}
