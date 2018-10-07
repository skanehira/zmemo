package handler

import (
	"encoding/json"
	"net/http"
	"testing"
	"zmemo/api/model"
	"zmemo/api/test"
)

// init server
var e = test.InitServer()

func generateNewUser() map[string]interface{} {
	return map[string]interface{}{"name": "test1", "password": "test1"}
}

func TestCreateUser(t *testing.T) {
	newUser := generateNewUser()

	j, err := json.Marshal(newUser)

	if err != nil {
		t.Errorf("cannot marsharl data cause: %s", err)
	}

	code, res := test.POST(e, "/users", j)

	user := new(model.User)

	if err := json.Unmarshal(res, user); err != nil {
		t.Errorf("cannot unmarshal response cause: %s", err)
		return
	}

	if code != http.StatusOK {
		t.Errorf("response code is unexcept: got=%d, expect=%d", code, http.StatusOK)
	}
	if user.ID == "" {
		t.Errorf("userid is empty in response")
		return
	}

	if user.Name != newUser["name"] {
		t.Errorf("user name is unexcept in response: got=%s, expect=%s", user.Name, newUser["name"])
		return
	}

	if user.Password != newUser["password"] {
		t.Errorf("user password is unexcept in response: got=%s, expect=%s", user.Password, newUser["password"])
		return
	}
}

func TestGetUser(t *testing.T) {
	newUser := generateNewUser()

	j, err := json.Marshal(newUser)

	if err != nil {
		t.Errorf("cannot marsharl data cause: %s", err)
	}

	test.POST(e, "/users", j)

	code, _ := test.GET(e, "/users")

	if code != http.StatusOK {
		t.Errorf("response code is unexcept: got=%d, expect=%d", code, http.StatusOK)
	}

}

func TestUpdateUser(t *testing.T) {
	test.InitServer()

}
