package pool

import "errors"

var (
	currentPoolCount = 0
)

type ConnectionPool struct {
	typ         CTypes
	inParameter interface{}
	count       int
	refreshTime int
}

func NewPool(typ CTypes, count, refreshTime int, params interface{}) (*ConnectionPool, error) {
	if count <= 0 {
		return nil, errors.New("count must be greater than zero")
	}
	return &ConnectionPool{
		typ:         typ,
		inParameter: params,
		count:       count,
		refreshTime: refreshTime,
	}, nil
}

func (cp *ConnectionPool) Get() (chan interface{}, error) {
	cnn := make(chan interface{}, cp.count)
	switch cp.typ {
	case postgres:
		for range cnn {

		}
		return nil, nil
	default:
		return nil, nil
	}

}
func (cp *ConnectionPool) Put(x ...interface{}) {

}
