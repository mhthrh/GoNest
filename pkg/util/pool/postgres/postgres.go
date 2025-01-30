package postgres

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/mhthrh/common-lib/config/model"
	customModelError "github.com/mhthrh/common-lib/errors"
	customeError "github.com/mhthrh/common-lib/errors/pool"
	"github.com/mhthrh/common-lib/pkg/util/pool"
	"log"
	"time"
)

const (
	psql                   = "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
	checkCnnStatusDuration = time.Millisecond * 10
	totalWaitForReleaseAll = time.Second * 10
)

var (
	dic    = map[string]*pool.Connection{}
	stop   = make(chan struct{}, 1)
	newCnn = make(chan pool.Connection, 1)
)

type Config struct {
	db          model.DB
	count       int
	refreshTime time.Duration
}

func New(d model.DB, count int, refresh time.Duration) pool.IConnection {
	return Config{
		db:          d,
		count:       count,
		refreshTime: refresh,
	}
}

func (p Config) Get() (*pool.Connection, *customModelError.XError) {
	for {
		for _, d := range dic {
			if !d.InUse {
				d.InUse = true
				return d, nil
			}
		}
	}
}

func (p Config) Put(key uuid.UUID) *customModelError.XError {
	if dic[key.String()] == nil {
		return customeError.DbCnnNotExist(nil)
	}
	dic[key.String()].InUse = false
	return nil
}

// Initialize must run with goroutine
func (p Config) Initialize() {
	go func() {
		for {
			<-time.After(p.refreshTime)
			for s, connection := range dic {
				if connection == nil {
					delete(dic, s)
				}
				if connection.Err != nil {
					delete(dic, s)
				}
				if connection.Cnn.(*sql.DB).Ping() != nil {
					connection = nil
					delete(dic, s)
				}
			}
		}
	}()

	for {
		select {
		case <-stop:
			return
		default:
			if len(dic) < p.count {
				key := uuid.New()
				cnn, err := sql.Open(p.db.Driver, fmt.Sprintf(psql, p.db.Host, p.db.Port, p.db.UserName, p.db.Password, p.db.DbName, p.db.SSLMode))
				if err != nil {
					log.Println(customeError.DbConnectionFailed(customModelError.RunTimeError(err)))
				}
				dic[key.String()] = &pool.Connection{
					CId:   key,
					Typ:   1,
					Cnn:   cnn,
					InUse: false,
					Err:   nil,
				}
			}
		}
	}
}

func (p Config) Release(uuid uuid.UUID) *customModelError.XError {
	pol := dic[uuid.String()]
	if pol == nil {
		return customeError.DbCnnNotExist(nil)
	}
	if pol.InUse {
		return customeError.DbCnnNotExist(nil)
	}
	_ = pol.Cnn.(*sql.DB).Close()
	delete(dic, uuid.String())
	return nil
}

func (p Config) ReleaseAll() *customModelError.XError {
	stop <- struct{}{}
	for s, pol := range dic {
		var total time.Duration
		if pol.InUse {
			for {
				<-time.After(checkCnnStatusDuration)
				total += checkCnnStatusDuration
				if total.Seconds() > totalWaitForReleaseAll.Seconds() {
					return customeError.ReleaseAllError(nil)
				}
				break
			}
		}
		total = time.Duration(0)
		_ = pol.Cnn.(*sql.DB).Close()
		delete(dic, s)
	}
	return nil
}
