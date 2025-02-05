package customer

//go:generate stringer -type=Types

type Types int

const (
	none Types = iota
	Card
	Passport
	GovernmentPaper
)
