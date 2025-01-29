package pool

import (
	"database/sql"
	"fmt"
	"github.com/mhthrh/common-lib/config/model"
	"time"
)

const (
	psql = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
)

type PostgresSQL struct {
	DB          model.DB
	Count       int
	RefreshTime time.Duration
}
//New MUST run with goroutine
func (p PostgresSQL) New(s chan struct{}) chan Pool {
	c := make(chan Pool, p.Count)
		defer close(c)
		for {
			select {
			case <-s:
			//close all release all and return
			case <-time.After(p.RefreshTime):
			default:
				cnn, err := sql.Open(p.DB.Driver, fmt.Sprintf(psql, p.DB.Host, p.DB.Port, p.DB.UserName, p.DB.Password, p.DB.DbName))
				if err != nil {

				}
			}
	}()
}

func (p PostgresSQL) Release(pools chan Pool) {
	//TODO implement me
	panic("implement me")
}
