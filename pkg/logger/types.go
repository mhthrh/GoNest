package logger



//go:generate stringer -type=Types

type Types int

const (
	none Types = iota
	info
	warn
	typeError
	fatal
	typePanic
)

