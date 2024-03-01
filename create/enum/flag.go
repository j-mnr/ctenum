package enum

//go:generate go run github.com/j-mnr/ctenum create enum flag.go
//ctenum:type=flag output=flag_ctenum.go
const (
	// _type is the Go Type that will be used for the enum name.
	_type = "type"
	// output is the file relative to where ctenum was called.
	output = "output"
	// _package is the name of the package to put the enum into.
	_package = "package"
	// makeFile is a boolean that defaults to making a file named
	// <type>_ctenumer.go in the current relative directory.
	makeFile = "make-file"
)
