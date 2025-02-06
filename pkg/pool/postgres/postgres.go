package postgres

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/mhthrh/common-lib/config/model"
	customModelError "github.com/mhthrh/common-lib/model/error"
	"github.com/mhthrh/common-lib/pkg/pool"
	"github.com/samber/lo"
	"strings"
	"sync"
	"time"
)

const (
	psql           = "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
	waitForFreeCnn = time.Millisecond * 100
)

var (
	connections map[string]pool.Connection
	ins         pool.IConnection
	once        sync.Once
)

type Config struct {
	db model.DB
	m  *sync.Mutex
}

func init() {
	connections = make(map[string]pool.Connection)
}

func New(db model.DB) (pool.IConnection, *customModelError.XError) {
	if strings.Trim(db.Host, " ") == "" {
		return nil, pool.InputParamsMismatch(nil)
	}
	once.Do(func() {
		ins = Config{
			db: db,
			m:  &sync.Mutex{},
		}
	})
	return ins, nil
}

func (c Config) Maker(request <-chan pool.Request, response chan<- pool.Response) {
	stop := false
	defer func() {
		if stop {
			response <- pool.Response{
				Total: 0,
				InUse: 0,
				Error: pool.StopSignal(nil),
			}
		}
	}()

	for {
		select {
		case r := <-request:

			if r.Stop {
				stop = true
				return
			}
			if r.Type != pool.Types(1) {
				response <- pool.Response{
					Total: 0,
					InUse: 0,
					Error: pool.ConnectionTypeNotAcceptable(nil),
				}
				continue
			}

			switch {
			case r.Count == 0:
				response <- pool.Response{
					Total: uint(len(connections)),
					InUse: uint(len(lo.PickBy(connections, func(key string, value pool.Connection) bool {
						return value.InUse == true
					}))),
					Error: nil,
				}
				continue
			case r.Count > 0:
				if len(connections) > int(r.Count) {
					response <- pool.Response{
						Total: r.Count,
						InUse: uint(len(lo.PickBy(connections, func(key string, value pool.Connection) bool {
							return value.InUse == true
						}))),
						Error: pool.MaximumConnection(nil),
					}
					break
				}
				for range int(r.Count) - len(connections) {
					m, err := newConnection(c.db)
					if err != nil {
						response <- pool.Response{
							Total: uint(len(connections)),
							InUse: uint(0),
							Error: err,
						}
						break
					}
					connections = merge(connections, m)
				}
			}
			response <- pool.Response{
				Total: uint(len(connections)),
				InUse: uint(len(lo.PickBy(connections, func(key string, value pool.Connection) bool {
					return value.InUse == true
				}))),
				Error: nil,
			}
			continue

		}
	}
}

func (c Config) Manager(cmd <-chan pool.ManageRequest, conn chan<- *pool.Connection) {
	for {
		select {
		case command := <-cmd:
			switch command.Command {
			case pool.Commands(0):
				conn <- &pool.Connection{
					Id:    uuid.UUID{},
					Cnn:   nil,
					InUse: false,
					Err:   pool.CommandNotExist(nil),
				}
			case pool.Commands(1):
				c.m.Lock()
				stop := time.After(waitForFreeCnn)
				for {
					select {
					case <-stop:
						c.m.Unlock()
						conn <- &pool.Connection{
							Id:    uuid.UUID{},
							Cnn:   nil,
							InUse: false,
							Err:   pool.FreeConnectionNotExist(nil),
						}
					default:
						for _, cn := range connections {
							if !cn.InUse {
								cn.InUse = true
								conn <- &cn
								c.m.Unlock()
								goto indicator
							}
						}
					}
				}
			case pool.Commands(2):
				cn, ok := connections[command.ID.String()]
				if !ok {
					conn <- &pool.Connection{
						Id:    uuid.UUID{},
						Cnn:   nil,
						InUse: false,
						Err:   pool.DbCnnNotExist(nil),
					}
					continue
				}
				cn.InUse = false
				conn <- &pool.Connection{
					Id:    uuid.UUID{},
					Cnn:   nil,
					InUse: false,
					Err:   customModelError.Success(),
				}

			}
		indicator:
		}
	}

}

func (c Config) Refresh(s chan struct{}, e chan *customModelError.XError) {
	for {
		select {
		case <-s:
			for id, conn := range connections {
				if conn.Cnn == nil {
					delete(connections, id)
					con, err := newConnection(c.db)
					if err != nil {
						e <- err
						continue
					}
					connections = merge(connections, con)
				}
				if conn.Cnn.(*sql.DB).Ping() != nil {
					_ = conn.Cnn.(*sql.DB).Close()
					conn.Cnn = nil
					delete(connections, id)
					con, err := newConnection(c.db)
					if err != nil {
						e <- err
						continue
					}
					connections = merge(connections, con)
				}
			}
			e <- customModelError.Success()
		}
	}
}

func (c Config) Release(request chan pool.ReleaseRequest, e chan *customModelError.XError) {
	for {
		select {
		case r := <-request:
			if r.Stop {
				e <- customModelError.Success()
				return
			}
			connection, ok := connections[r.ID.String()]
			if !ok {
				e <- pool.DbCnnNotExist(nil)
				continue
			}
			c.m.Lock()
			_ = connection.Cnn.(*sql.DB).Close()
			connection.Cnn = nil
			delete(connections, r.ID.String())
			c.m.Unlock()
			e <- customModelError.Success()
		}
	}
}

func (c Config) ReleaseAll(byForce bool) *customModelError.XError {
	c.m.Lock()
	defer c.m.Unlock()
	for _, connection := range connections {
		if !byForce && connection.InUse {
			return pool.ConnectionInUse(nil)
		}
		_ = connection.Cnn.(*sql.DB).Close()
		connection.Cnn = nil
	}
	return nil
}

func newConnection(d model.DB) (m map[string]pool.Connection, e *customModelError.XError) {
	m = make(map[string]pool.Connection)
	cnn, err := sql.Open(d.Driver, fmt.Sprintf(psql, d.Host, d.Port, d.UserName, d.Password, d.DbName, d.SSLMode))
	if err != nil {
		e = pool.DbConnectionFailed(customModelError.RunTimeError(err))
		return nil, e
	}
	key := uuid.New()
	m[key.String()] = pool.Connection{
		Id:    key,
		Cnn:   cnn,
		InUse: false,
	}
	return m, nil
}
func merge(dictionaries ...map[string]pool.Connection) map[string]pool.Connection {
	res := make(map[string]pool.Connection)
	for _, di := range dictionaries {
		for k, v := range di {
			res[k] = v
		}
	}
	return res
}
