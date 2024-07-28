package pkg

import (
	"errors"
	"flag"
	"os"
)

type Flags struct {
	Dironly     bool
	Simonly     bool
	Fileonly    bool
	FileExtonly bool
}

type Args struct {
	Ext     string
	DirPath string
}

func FindFlagAndArg() (fl Flags, arg Args, err error) {
	flag.BoolVar(&fl.Dironly, "d", false, "Only directory")
	flag.BoolVar(&fl.Simonly, "sl", false, "Only symlinks")
	flag.BoolVar(&fl.Fileonly, "f", false, "Only filenames")
	flag.StringVar(&arg.Ext, "ext", "", "Specification of file extension")
	flag.Parse()

	if arg.Ext != "" && !fl.Fileonly {
		return fl, arg, ErrWrongFlagsCombination
	}

	if _, err := os.Stat(arg.DirPath); os.IsNotExist(err) {
		return fl, arg, ErrNoSuchDirectory
	}

	return fl, arg, nil
}

var ErrWrongFlagsCombination = errors.New("flag -ext works only with flag -f")
var ErrNoSuchDirectory = errors.New("no such directory")
var ErrNoDirPassed = errors.New("no directory passed")

func Noflags(fl Flags) bool {
	return !fl.Dironly && !fl.Simonly && !fl.Fileonly
}

func NoDirectory(arg Args) error {
	if _, err := os.Stat(arg.DirPath); os.IsNotExist(err) {
		return ErrNoSuchDirectory
	}
	return nil
}

func WrongCombiantion(fl Flags) error {
	if fl.FileExtonly && !fl.Fileonly {
		return ErrWrongFlagsCombination
	}
	return nil
}

func NoDirPass(arg Args) error {
	if arg.DirPath == "" {
		return ErrNoDirPassed
	}
	return nil
}

func NotNilExt(fl Flags, arg Args) bool {
	return arg.Ext != ""
}
