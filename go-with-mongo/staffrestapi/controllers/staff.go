package controllers

import (
	"encoding/json"
	. "github.com/Vectormike/staffrestapi/config"
	. "github.com/Vectormike/staffrestapi/dao"
	. "github.com/Vectormike/staffrestapi/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"os"
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
	respondWithJson(w, http.StatusOK, staff)
}

// Get staff by id
func FindStaff(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	staff, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Staff ID")
		return
	}
	respondWithJson(w, http.StatusOK, staff)
}

// POST a new staff
func CreateStaff(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var staff Staff
	if err := json.NewDecoder(r.Body).Decode(&staff); err != nil {
		respondWithError(w, httpStatusBadRequest, "Unprocessable Entity")
	}
	staff.ID = bson.NewObjectId()
	if err := dao.Insert(staff); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusCreated, staff)
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
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
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
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
