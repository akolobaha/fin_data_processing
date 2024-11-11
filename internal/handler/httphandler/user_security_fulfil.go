package httphandler

import (
	"encoding/json"
	"fin_data_processing/db"
	"fin_data_processing/internal/entities"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func AddSecurityFulfil(w http.ResponseWriter, r *http.Request) {
	gDB := db.GetGormDbConnection()
	var newUserSecFulfil entities.UserSecurityFulfil
	json.NewDecoder(r.Body)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newUserSecFulfil); err != nil {
		http.Error(w, err.Error(), 400)
	}

	result := gDB.Create(&newUserSecFulfil)
	if err := result.Error; err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	renderJSON(w, newUserSecFulfil)
}

func SecurityFulfilsList(w http.ResponseWriter, r *http.Request) {
	gDB := db.GetGormDbConnection()
	var userSecFulfils []entities.UserSecurityFulfil

	// Применяем пагинацию к запросу
	paginatedDB := Paginate(r)(gDB)

	result := paginatedDB.Order("id").Find(&userSecFulfils)

	if err := result.Error; err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	renderJSON(w, userSecFulfils)
}

func SecurityFulfilOne(w http.ResponseWriter, r *http.Request) {
	gDB := db.GetGormDbConnection()
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 0)
	if err != nil {
		http.Error(w, "invalid syntax", 400)
		return
	}

	var userSecFulfil entities.UserSecurityFulfil
	result := gDB.First(&userSecFulfil, id)
	if err := result.Error; err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	renderJSON(w, userSecFulfil)
}

func SecurityFulfilDelete(w http.ResponseWriter, r *http.Request) {
	gDB := db.GetGormDbConnection()
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 0)
	if err != nil {
		http.Error(w, "invalid syntax", 400)
		return
	}

	var userSecFulfil entities.UserSecurityFulfil
	result := gDB.Delete(userSecFulfil, id)
	if err := result.Error; err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	renderJSON(w, []string{"ok"})
}
