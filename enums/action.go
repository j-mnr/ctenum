package enums

//go:generate go run github.com/j-mnr/ctenum create enum action.go
//ctenum:type=action package=enums make-file=true
const (
	create  = "create"
	version = "version"
	help    = "help"
)
