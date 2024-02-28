package enums

//go:generate ctenum create enum action.go
//ctenum:type=action package=enums make-file=true
const (
	create  = "create"
	version = "version"
	help    = "help"
)
