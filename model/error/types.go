package error

//go:generate stringer -type=Types
// Types please add new types just at the end
type Types int

const (
	runTime Types = iota
	System
	general
	notImplemented
	notSupported
	reserved1
	reserved2
	reserved3
	reserved4
	database
	customer
	city
	country
	pool
	loader
	address
	config
)
