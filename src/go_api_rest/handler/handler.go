package handler

import (
	"encoding/json"
	"go_api_rest/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// func GetUsers(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	// w.Header().Set("Content-Type", "text/xml")

// 	db.Connect()
// 	users, _ := models.ListUsers()
// 	db.Close()

//		// output, _ := xml.Marshal(users)
//		// output, _ := yaml.Marshal(users)
//		output, _ := json.Marshal(users)
//		fmt.Fprintln(w, string(output))
//	}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if users, err := models.ListUsers(); err != nil {
		models.SendNoFound(w)
	} else {
		models.SendData(w, users)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if user, err := getUserByRequest(r); err != nil {
		models.SendNoFound(w)
	} else {
		models.SendData(w, user)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	//Obtener Registro
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(w)
	} else {
		user.Save()
		models.SendData(w, user)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	//Obtener Registro
	var userId int64

	if user, err := getUserByRequest(r); err != nil {
		models.SendNoFound(w)
	} else {
		userId = user.Id
	}

	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(w)
	} else {
		user.Id = userId
		user.Save()
		models.SendData(w, user)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if user, err := getUserByRequest(r); err != nil {
		models.SendNoFound(w)
	} else {
		user.Delete()
		models.SendData(w, user)
	}
}

func getUserByRequest(r *http.Request) (models.User, error) {
	//Obtener ID
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	if user, err := models.GetUser(userId); err != nil {
		return *user, err
	} else {
		return *user, nil
	}

}
