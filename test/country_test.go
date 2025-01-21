package test

import (
	"common-lib/model/address"
	"testing"
)

func TestCountries(t *testing.T) {
	l, e := address.Countries()
	if e != nil {
		t.Error(e)
	}
	if len(l) == 0 {
		t.Error("list is empty")
	}
	if l[0].Code != "AF" {
		t.Error("list is wrong")
	}
}
