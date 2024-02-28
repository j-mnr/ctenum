package examples

//go:generate ctenum create enum color.yaml
//go:generate ctenum create enum day.yaml
//go:generate ctenum create enum -output weekdays-json/day_ctenum.go day.json
//go:generate ctenum create enum -package dir -output dir/direction_ctenum.go direction.yaml
//go:generate ctenum create enum -package env -output lifecycle/lifecycle_ctenum.go lifecycle.yaml
//go:generate ctenum create enum -package http -output http/httpmethod_ctenum.go httpmethod.yaml
