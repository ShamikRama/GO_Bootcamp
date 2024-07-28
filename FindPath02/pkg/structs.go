package pkg

import (
	"flag"
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
	flag.StringVar(&arg.Ext, "ext", "", "Specification of file")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		arg.Ext = args[0]
	}
	return fl, arg, err
}
