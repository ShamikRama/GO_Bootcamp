package ex00

import (
	"FindPath02/pkg"
	"fmt"
	"os"
	"path/filepath"
)

func IterateOverEntities(arg pkg.Args, fl pkg.Flags) error {
	return IterateDir(arg.DirPath, arg, fl)
}

func IterateDir(dirpath string, arg pkg.Args, fl pkg.Flags) error {
	entries, err := os.ReadDir(dirpath)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return err
		}
		path := filepath.Join(dirpath, entry.Name())

		switch {
		case fl.Simonly && IsSymlink(info):
			PrintPath(path, info)
		case fl.Dironly && info.IsDir():
			PrintPath(path, info)
			if err := IterateDir(path, arg, fl); err != nil {
				return err
			}
		case fl.FileExtonly && filepath.Ext(path) == ("."+arg.Ext):
			PrintPath(path, info)
		case fl.Fileonly && IsFile(info):
			PrintPath(path, info)
		case pkg.Noflags(fl):
			PrintPath(path, info)
		}
	}
	return nil
}

func PrintPath(path string, info os.FileInfo) {
	if info.Mode()&os.ModeSymlink == os.ModeSymlink {
		targetPath, err := os.Readlink(path)
		if err != nil {
			targetPath = "[broken]"
		}
		fmt.Println(path, "->", targetPath)
	} else {
		fmt.Println(path)
	}
}

func IsFile(info os.FileInfo) bool {
	return info.Mode().IsRegular()
}

func IsSymlink(info os.FileInfo) bool {
	return info.Mode()&os.ModeSymlink == os.ModeSymlink
}
