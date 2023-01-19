package dao

import (
	. "github.com/Vectormike/staffrestapi/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type StaffDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "staff"
)

// Establish a connection to database
func (m *StaffDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Get all staff
func (m *StaffDAO) FindAll() ([]Staff, error) {
	var staff []Staff
	err := db.C(COLLECTION).Find(bson.M{}).All(&staff)
	return staff, err
}

// Get staff by id
func (m *StaffDAO) FindById(id string) (Staff, error) {
	var staff Staff
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&staff)
	return staff, err
}

// Update staff by id
func (m *StaffDAO) Update(staff Staff) error {
	err := db.C(COLLECTION).UpdateId(staff.ID, &staff)
	return err
}

// Delete staff by id
func (m *StaffDAO) Delete(staff Staff) error {
	err := db.C(COLLECTION).Remove(&staff)
	return err
}

// Save Staff
func (m *StaffDAO) Insert(staff Staff) error {
	err := db.C(COLLECTION).Insert(&staff)
	return err
}
