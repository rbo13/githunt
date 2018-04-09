package main

import (
	"testing"
)

func TestItShouldReturnUser(t *testing.T) {

	user, err := GetUserByUsername("whaangbuu")

	expectedLoginName := "whaangbuu"
	gotLoginName := user.Login

	if err != nil {
		t.FailNow()
	}

	if expectedLoginName != gotLoginName {
		t.Errorf("Expects: %s, but got %s", expectedLoginName, gotLoginName)
	}

	t.Log(user)
}
