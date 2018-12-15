package test

import (
	"king_and_angle/modeland_angle/model"
	"king_and_angle/resource_angle/resource"
	"testing"
)

func TestVerify(t *testing.T){
	var(
		result string
		err error
	)
	result, err = resource.Verify("5602216028", "5602216028")
	if err != nil || result == ""{
		t.Fail()
	}

	result, err = resource.Verify("", "")
	if err.Error() != "verify error" || result != ""{
		t.Fail()
	}
}


func TestGetUserById(t *testing.T){
	var(
		user model.User
	)

	user = resource.GetUserById(50)
	if user.ID != 50{
		t.Fail()
	}

	user = resource.GetUserById(1111111)
	if user.ID != 0{
		t.Fail()
	}

	user = resource.GetUserById(20)
	if user.ID != 0{
		t.Fail()
	}
}

func TestGetByUserName(t *testing.T){
	var(
		user model.User
	)

	user = resource.GetUserByUsername("5602216028")
	if user.Username != "5602216028"{
		t.Fail()
	}

	user = resource.GetUserByUsername("5602216033")
	if user.Username != ""{
		t.Fail()
	}

	user = resource.GetUserByUsername("")
	if user.Username != ""{
		t.Fail()
	}
}
