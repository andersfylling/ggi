package ggi

import (
	"os"
	"strings"
)


// GetFiles returns a list of files from the pkg of your choice.
// To be sure, just use an absolute path which can be retrieved from runtime:
//  package main
//
//  import (
//    "fmt"
//    "path/filepath"
//    "runtime"
//
//    "github.com/andersfylling/ggi"
//   )
//
//   var (
//     _, b, _, _ = runtime.Caller(0)
//     basepath   = filepath.Dir(b)
//     genpath    = "/generate/testing" // diff path; from root pkg to this files pkg
//   )
//
//   func main() {
//	   path := basepath[:len(basepath)-len(genpath)]
//	   files, err := ggi.GetFiles(path)
//	   if err != nil {
//       panic(err)
//     }
//
//     // TODO: make sure these prints the .go files in your desired directory
//     for _, f := range files {
//       fmt.Println(f)
//     }
//   }
func GetFiles(path string) (files []string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		f := file.Name()
		isGoFile := strings.HasSuffix(f, ".go")
		isInSubDir := strings.Contains(f, "/")
		isGenFile := strings.HasSuffix(f, "_gen.go")
		if f == path || !isGoFile || isInSubDir || isGenFile {
			continue
		}

		files = append(files, path + "/" + f)
	}

	return files, nil
}