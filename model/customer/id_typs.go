package customer

//go:generate stringer -type=IdTypes

type IdTypes int

const (
	none  IdTypes = iota
	Card
	Passport
	GovernmentPaper
)
