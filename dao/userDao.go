package dao

import(
	"TODOLIST/model"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"context"
	"github.com/dgrijalva/jwt-go"
	"os"
)


func CheckPasswordHash(pass, passhash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passhash), []byte(pass))
	return err == nil
}
func Login(name, pass string) (bool,int){
	Db:=OpenDbConnection()
	var user model.User
	ctx := context.Background()

	stmt:=fmt.Sprintf("select * from user where name='%s'",name)

	_=Db.QueryRowContext(ctx, stmt).Scan(&user.Id,&user.Name,&user.Pass)
	check:=CheckPasswordHash(pass,user.Pass)
	return check,user.Id
}

func Register(name, pass string ) string{
	Db:=OpenDbConnection()

	stmt:=fmt.Sprintf("insert into user(name, pass) values('%s','%s')",name,pass)
	_,err:=Db.Exec(stmt)
	fmt.Println(err)
	var user model.User
	var resp model.Response
	tk := &model.Token{UserId: user.Id}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"),tk)
	tokenString, _:=token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString
	user.Pass = ""
	resp.Success = "Successful!"
	response := resp.Success

	Db.Close()
	return response
}




