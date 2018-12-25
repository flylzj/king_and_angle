package test

import (
	"github.com/dgrijalva/jwt-go"
	"model"
	"resource"
	"testing"
	"time"
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


func TestCreateToken(t *testing.T){
	j := resource.NewJWT()
	token, _ := j.CreateToken(resource.CustomClaims{ID:58, StandardClaims:jwt.StandardClaims{ExpiresAt:time.Now().Add(time.Hour * 1).Unix()}})
	t.Log(token)

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
