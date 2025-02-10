package postgres

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	customModelError "github.com/mhthrh/GoNest/model/error"
	"github.com/mhthrh/GoNest/model/loader"
	cPool "github.com/mhthrh/GoNest/model/pool"
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
	connections map[string]cPool.Connection
	ins         cPool.IConnection
	once        sync.Once
)

type Config struct {
	db loader.DB
	m  *sync.Mutex
}

func init() {
	connections = make(map[string]cPool.Connection)
}

func New(db loader.DB) (cPool.IConnection, *customModelError.XError) {
	if strings.Trim(db.Host, " ") == "" {
		return nil, cPool.InputParamsMismatch(nil)
	}
	once.Do(func() {
		ins = Config{
			db: db,
			m:  &sync.Mutex{},
		}
	})
	return ins, nil
}

func (c Config) Maker(request <-chan cPool.Request, response chan<- cPool.Response) {
	stop := false
	defer func() {
		if stop {
			response <- cPool.Response{
				Total: 0,
				InUse: 0,
				Error: cPool.StopSignal(nil),
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
			if r.Type != cPool.Types(1) {
				response <- cPool.Response{
					Total: 0,
					InUse: 0,
					Error: cPool.ConnectionTypeNotAcceptable(nil),
				}
				continue
			}

			switch {
			case r.Count == 0:
				response <- cPool.Response{
					Total: uint(len(connections)),
					InUse: uint(len(lo.PickBy(connections, func(key string, value cPool.Connection) bool {
						return value.InUse == true
					}))),
					Error: nil,
				}
				continue
			case r.Count > 0:
				if len(connections) > int(r.Count) {
					response <- cPool.Response{
						Total: r.Count,
						InUse: uint(len(lo.PickBy(connections, func(key string, value cPool.Connection) bool {
							return value.InUse == true
						}))),
						Error: cPool.MaximumConnection(nil),
					}
					break
				}
				for range int(r.Count) - len(connections) {
					m, err := newConnection(c.db)
					if err != nil {
						response <- cPool.Response{
							Total: uint(len(connections)),
							InUse: uint(0),
							Error: err,
						}
						break
					}
					connections = merge(connections, m)
				}
			}
			response <- cPool.Response{
				Total: uint(len(connections)),
				InUse: uint(len(lo.PickBy(connections, func(key string, value cPool.Connection) bool {
					return value.InUse == true
				}))),
				Error: nil,
			}
			continue

		}
	}
}

func (c Config) Manager(cmd <-chan cPool.ManageRequest, conn chan<- *cPool.Connection) {
	for {
		select {
		case command := <-cmd:
			switch command.Command {
			case cPool.Commands(0):
				conn <- &cPool.Connection{
					Id:    uuid.UUID{},
					Cnn:   nil,
					InUse: false,
					Err:   cPool.CommandNotExist(nil),
				}
			case cPool.Commands(1):
				c.m.Lock()
				stop := time.After(waitForFreeCnn)
				for {
					select {
					case <-stop:
						c.m.Unlock()
						conn <- &cPool.Connection{
							Id:    uuid.UUID{},
							Cnn:   nil,
							InUse: false,
							Err:   cPool.FreeConnectionNotExist(nil),
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
			case cPool.Commands(2):
				cn, ok := connections[command.ID.String()]
				if !ok {
					conn <- &cPool.Connection{
						Id:    uuid.UUID{},
						Cnn:   nil,
						InUse: false,
						Err:   cPool.DbCnnNotExist(nil),
					}
					continue
				}
				cn.InUse = false
				conn <- &cPool.Connection{
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

func (c Config) Release(request chan cPool.ReleaseRequest, e chan *customModelError.XError) {
	for {
		select {
		case r := <-request:
			if r.Stop {
				e <- customModelError.Success()
				return
			}
			connection, ok := connections[r.ID.String()]
			if !ok {
				e <- cPool.DbCnnNotExist(nil)
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
			return cPool.ConnectionInUse(nil)
		}
		_ = connection.Cnn.(*sql.DB).Close()
		connection.Cnn = nil
	}
	return nil
}

func newConnection(d loader.DB) (m map[string]cPool.Connection, e *customModelError.XError) {
	m = make(map[string]cPool.Connection)
	cnn, err := sql.Open(d.Driver, fmt.Sprintf(psql, d.Host, d.Port, d.UserName, d.Password, d.DbName, d.SSLMode))
	if err != nil {
		e = cPool.DbConnectionFailed(customModelError.RunTimeError(err))
		return nil, e
	}
	key := uuid.New()
	m[key.String()] = cPool.Connection{
		Id:    key,
		Cnn:   cnn,
		InUse: false,
	}
	return m, nil
}
func merge(dictionaries ...map[string]cPool.Connection) map[string]cPool.Connection {
	res := make(map[string]cPool.Connection)
	for _, di := range dictionaries {
		for k, v := range di {
			res[k] = v
		}
	}
	return res
}
