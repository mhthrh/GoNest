package customer

//go:generate stringer -type=Status

type Status int

const (
	Unknown  Status = iota // Unknown status
	Active                 // Active user
	Inactive               // Inactive user
	Banned                 // Banned user
)

