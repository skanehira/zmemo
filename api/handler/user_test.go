package handler

import (
	"testing"
	"zmemo/api/test"
)

func TestCreateUser(t *testing.T) {
	e := test.InitServer()
	j := []byte(`{"userName":"test1","password":"test1"}`)

	code, res := test.POST(e, "/users", j)

	t.Logf("code=%d, body=%s", code, res)
}

func TestGetUser(t *testing.T) {
	e := test.InitServer()
	j := []byte(`{"userName":"test1","password":"test1"}`)

	code, res := test.POST(e, "/users", j)

	code, res = test.GET(e, "/users")

	t.Logf("code=%d, body=%s", code, res)
}
