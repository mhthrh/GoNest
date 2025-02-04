package pool



//go:generate stringer -type=CCommands

type CCommands int

const (
	nothing CCommands = iota
	get
	set
	getIfNotAnyFree

)
