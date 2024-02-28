package create

import (
	"flag"
	"io"
	"os"

	"github.com/j-mnr/ctenum/create/enum"
)

func Cmd(version string, args []string, out, errW io.Writer) {
	ffs := flag.NewFlagSet("create flags", flag.ExitOnError)
	ffs.SetOutput(errW)
	ffs.Usage = func() {
		ffs.Output().Write([]byte("Usage of create:\n\tctenum create [flags] entity\n"))
		ffs.PrintDefaults()
		os.Exit(1)
	}
	if err := ffs.Parse(args); err != nil {
		ffs.Output().Write([]byte(err.Error()))
		ffs.Usage()
	}
	if ffs.NArg() == 0 {
		ffs.Usage()
	}
	switch ffs.Args()[0] {
	case "enum":
		enum.Build(version, ffs.Args()[1:], out, errW)
	default:
	}
}
