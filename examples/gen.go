package examples

//go:generate go run github.com/j-mnr/ctenum create enum color.yaml
//go:generate go run github.com/j-mnr/ctenum create enum day.yaml
//go:generate go run github.com/j-mnr/ctenum create enum -output weekdays-json/day_ctenum.go day.json
//go:generate go run github.com/j-mnr/ctenum create enum -package dir -output dir/direction_ctenum.go direction.yaml
//go:generate go run github.com/j-mnr/ctenum create enum -package env -output lifecycle/lifecycle_ctenum.go lifecycle.yaml
//go:generate go run github.com/j-mnr/ctenum create enum -package http -output http/httpmethod_ctenum.go httpmethod.yaml
