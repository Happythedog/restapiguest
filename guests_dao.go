package main

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type GuestsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "guests"
)

// Establish a connection to database
func (m *GuestsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of guests
func (m *GuestsDAO) FindAll() ([]Guest, error) {
	var Guests []Guest
	err := db.C(COLLECTION).Find(bson.M{}).All(&Guests)
	return Guests, err
}

// Find a Guest by its id
func (m *GuestsDAO) FindById(id string) (Guest, error) {
	var Guest Guest
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&Guest)
	return Guest, err
}

// Insert a Guest into database
func (m *GuestsDAO) Insert(Guest Guest) error {
	err := db.C(COLLECTION).Insert(&Guest)
	return err
}

// Delete an existing Guest
func (m *GuestsDAO) Delete(Guest Guest) error {
	err := db.C(COLLECTION).Remove(&Guest)
	return err
}

// Update an existing Guest
func (m *GuestsDAO) Update(Guest Guest) error {
	err := db.C(COLLECTION).UpdateId(Guest.ID, &Guest)
	return err
}
