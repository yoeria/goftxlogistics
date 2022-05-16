package main

import (
	"io/fs"
	fp "path/filepath"
)

// Recursively find a file (or multiple) in tree and return path
func FindFile(root, filename string) (sPath string, err error) {
	fp.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.Name() == filename {
			sPath = path
		}
		return nil
	})

	return
}
