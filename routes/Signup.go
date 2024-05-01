package routes

import (
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	/*	var t models.User
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			http.Error(w, "Error reciving data "+err.Error(), 400)
			return
		}
		if len(t.Email) == 0 {
			http.Error(w, "Email is required", 400)
			return
		}
		if len(t.Password) < 6 {
			http.Error(w, "six digits minium lenght", 400)
			return
		}

		_,e,_ := bd.CheckUserExist(t.Email)

		if e == true{
			http.Error(w, "this user extist", 400)
			return
		}

		_,status,err := bd.Insert(t)
		if err != nil {
			http.Error(w, "error inserting user " +err.Error(), 400)
			return
		}

		if status == false {
			http.Error(w, "cant insert user" +err.Error(), 400)
			return
		}*/

	w.WriteHeader(http.StatusCreated)
}
