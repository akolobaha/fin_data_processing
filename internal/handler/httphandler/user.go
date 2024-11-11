package httphandler

import (
	"encoding/json"
	"fin_data_processing/db"
	"fin_data_processing/internal/entities"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	gDB := db.GetGormDbConnection()
	var newUser entities.User
	json.NewDecoder(r.Body)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newUser); err != nil {
		http.Error(w, err.Error(), 400)
	}

	result := gDB.Create(&newUser)
	if err := result.Error; err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	renderJSON(w, newUser)
}

func UsersList(w http.ResponseWriter, r *http.Request) {
	gDB := db.GetGormDbConnection()
	var users []entities.User

	// Применяем пагинацию к запросу
	paginatedDB := Paginate(r)(gDB)

	result := paginatedDB.Order("id").Find(&users)

	if err := result.Error; err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	renderJSON(w, users)
}

func UserOne(w http.ResponseWriter, r *http.Request) {

	gDB := db.GetGormDbConnection()
	vars := mux.Vars(r)
	userId, err := strconv.ParseUint(vars["userId"], 10, 0)

	if err != nil {
		http.Error(w, "invalid syntax", 400)
		return
	}
	var user entities.User

	result := gDB.First(&user, userId)

	if err := result.Error; err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	renderJSON(w, user)
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	gDB := db.GetGormDbConnection()
	vars := mux.Vars(r)
	userId, err := strconv.ParseUint(vars["id"], 10, 0)
	if err != nil {
		http.Error(w, "invalid syntax", 400)
		return
	}

	var user entities.User
	result := gDB.Delete(user, userId)
	if err := result.Error; err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
