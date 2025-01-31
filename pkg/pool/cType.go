package pool

//go:generate stringer -type=CTypes

type CTypes int

const (
	none CTypes = iota
	postgres
	oracle
	msSql
	redis
	mySql
	kafka
	ibmMq
	rabbit
)
