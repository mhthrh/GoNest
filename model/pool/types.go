package pool


//go:generate stringer -type=Types

type Types int

const (
	none Types = iota
	postgres
	oracle
	msSql
	redis
	mySql
	kafka
	ibmMq
	rabbit
)
