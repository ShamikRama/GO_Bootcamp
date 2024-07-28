package ex00

import (
	pkg "FindPath02/pkg"
	"errors"
	"os"
)

var ErrWrongFlagsCombination = errors.New("flag -ext works only with flag -f")
var ErrNoSuchDirectory = errors.New("no such directory")
var ErrNoDirPassed = errors.New("no directory passed")

func Noflags(fl pkg.Flags) bool {
	return !fl.Dironly && !fl.Simonly && !fl.Fileonly
}

func NoDirectory(arg pkg.Args) error {
	if _, err := os.Stat(arg.DirPath); os.IsNotExist(err) {
		return ErrNoSuchDirectory
	}
	return nil
}

func WrongCombiantion(fl pkg.Flags) error {
	if fl.FileExtonly && !fl.Fileonly {
		return ErrWrongFlagsCombination
	}
	return nil
}

func NoDirPass(arg pkg.Args) error {
	if arg.DirPath == "" {
		return ErrNoDirPassed
	}
	return nil
}

func NotNilExt(fl pkg.Flags, arg pkg.Args) bool {
	if arg.Ext != "" {
		fl.FileExtonly = true
	}
	return fl.FileExtonly
}
