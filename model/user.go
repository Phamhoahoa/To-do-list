package model
import (
	"github.com/dgrijalva/jwt-go"
)
type  User struct{
	Id int
	Name string
	Pass string
	Token string
}

type InfoRegister struct{
	Name string
	Pass1 string
	Pass2 string
}
type Token struct{
	UserId int
	jwt.StandardClaims
}