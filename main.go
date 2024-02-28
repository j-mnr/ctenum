package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/j-mnr/ctenum/create"
	"github.com/j-mnr/ctenum/enums"
)

var Version = "dev"

func main() {
	if err := run(context.Background(), os.Args[1:], os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context, args []string, out, errW io.Writer) error {
	ffs := flag.NewFlagSet("ctenum flags", flag.ExitOnError)
	ffs.SetOutput(errW)
	ffs.Usage = func() {
		ffs.Output().Write([]byte("Usage of ctenum:\n\tctenum [flags] action file\n"))
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
	action, err := enums.ToActionEnum(ffs.Args()[0])
	if err != nil {
		fmt.Fprintln(errW, err)
		ffs.Usage()
	}
	switch action {
	case enums.ActionEnumCreate:
		create.Cmd(Version, ffs.Args()[1:], out, errW)
	default:
		panic("Impossible to make it in default case for switch action!")
	}

	return nil
}
