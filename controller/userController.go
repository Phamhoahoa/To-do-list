package controller

import(
	"net/http"
	"TODOLIST/model"
	"TODOLIST/dao"
	"encoding/json"
	// "strconv"
	"strings"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	
)

func HashPassword(pass string) (string, error){
	passHash,err:=bcrypt.GenerateFromPassword([]byte(pass),14)
	return string(passHash),err
}

func Register(w http.ResponseWriter, r *http.Request){
	var info model.InfoRegister
	var res model.Response
	decoder:=json.NewDecoder(r.Body)
	err:=decoder.Decode(&info)
	if err != nil{
		fmt.Println("Error decode!")
		return
	}
	r.Body.Close()

	fmt.Println(info.Name,info.Pass1,info.Pass2)

	if strings.Compare(info.Pass1,info.Pass2)==0{
		if len(info.Pass1)>8{
			passhash,_:=HashPassword(info.Pass1)
			
			err:=dao.Register(info.Name,passhash)
			if err!=nil{
				fmt.Println(err)
				res.Err="Register fail!"
				json.NewEncoder(w).Encode(res)
			}else{
				res.Success="Register success!"
				json.NewEncoder(w).Encode(res)
			}
		}else{
			res.Err="Password must long than 8 characters!"
			json.NewEncoder(w).Encode(res)
		}
		
	}else{
		res.Err="Does not match password!"
		json.NewEncoder(w).Encode(res)
	}
	
}

func Login(w http.ResponseWriter, r *http.Request){
	var user model.User
	var res model.Response
	decoder:=json.NewDecoder(r.Body)
	err:=decoder.Decode(&user)
	if err != nil{
		fmt.Println("Error decode!")
		return
	}
	r.Body.Close()
	fmt.Println(user.Name,user.Pass)
	check,id:=dao.Login(user.Name,user.Pass)
	fmt.Println(check)
	user.Id=id
	if check==true{
		json.NewEncoder(w).Encode(user)
	}else{
		res.Err="Username or password wrong!"
		json.NewEncoder(w).Encode(res)
	}



}