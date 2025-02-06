package loader

//go:generate stringer -type=SslType

type SslType int

const (
	disable SslType = iota
	enable
)
