package pool



//go:generate stringer -type=Commands

type Commands int

const (
	nothing Commands = iota
	get
	set
	getIfNotAnyFree

)
