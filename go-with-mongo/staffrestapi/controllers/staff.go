package controllers

import (
	"encoding/json"
	. "github.com/Vectormike/go-with-mongo/staffrestapi/config"
	. "github.com/Vectormike/go-with-mongo/staffrestapi/dao"
	. "github.com/Vectormike/go-with-mongo/staffrestapi/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

var config = Config{}
var dao = StaffDAO{}

// GET list of staff
func AllStaff(w http.ResponseWriter, r *http.Request) {
	staff, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	retResponse(w, http.StatusOK, staff)
}

// Get staff by id
func FindStaff(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	staff, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Staff ID")
		return
	}
	retResponse(w, http.StatusOK, staff)
}

// POST a new staff
func CreateStaff(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var staff Staff
	if err := json.NewDecoder(r.Body).Decode(&staff); err != nil {
		respondWithError(w, http.StatusBadRequest, "Unprocessable Entity")
	}
	staff.ID = bson.NewObjectId()
	if err := dao.Insert(staff); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	retResponse(w, http.StatusCreated, staff)
}

// PUT update an existing staff
func UpdateStaff(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var staff Staff
	if err := json.NewDecoder(r.Body).Decode(&staff); err != nil {
		respondWithError(w, http.StatusBadRequest, "Unprocessable Entity")
		return
	}
	if err := dao.Update(staff); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	retResponse(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing staff
func DeleteStaff(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var staff Staff
	if err := json.NewDecoder(r.Body).Decode(&staff); err != nil {
		respondWithError(w, http.StatusBadRequest, "Unprocessable Entity")
		return
	}
	if err := dao.Delete(staff); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	retResponse(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	retResponse(w, code, map[string]string{"error": msg})
}

func retResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
