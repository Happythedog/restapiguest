package main

import "gopkg.in/mgo.v2/bson"

// Represents a guest, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Guest struct {
	ID                        bson.ObjectId `bson:"_id" json:"id"`
	GuestId                   string        `bson:"guestid" json:"guestid"`
	Surname                   string        `bson:"surname" json:"surname"`
	GivenName                 string        `bson:"givenName" json:"givenName"`
	DateOfBirth               string        `bson:"dateofbirth" json:"dateofbirth"`
	Nationality               string        `bson:"nationality" json:"nationality"`
	ArrivedFrom               string        `bson:"arrived_from" json:"arrived_from"`
	BlackListStatus           string        `bson:"blackliststatus" json:"blackliststatus"`
	IdentificationType        string        `bson:"identificationType" json:"identificationType"`
	DateOfArrivalInHotel      string        `bson:"dateOfArrivalInHotel" json:"dateOfArrivalInHotel"`
	VerificationStatus        string        `bson:"verificationStatus" json:"verificationStatus"`
	Remarks                   string        `bson:"remarks" json:"remarks"`
	HotelName                 string        `bson:"hotelName,omitempty" json:"hotelName,omitempty"`
	HotelAddress              string        `bson:"hotelAddress,omitempty" json:"hotelAddress,omitempty"`
	Zone                      string        `bson:"zone,omitempty" json:"zone,omitempty"`
	City                      string        `bson:"city,omitempty" json:"city,omitempty"`
	State                     string        `bson:"string,omitempty" json:"string,omitempty"`
	HotelPhoneNo              string        `bson:"hotelPhoneNo,omitempty" json:"hotelPhoneNo,omitempty"`
	FormCAppNo                string        `bson:"formCAppNo,omitempty" json:"formCAppNo,omitempty"`
	VisaNumber                string        `bson:"visaNumber,omitempty" json:"visaNumber,omitempty"`
	ValidTill                 string        `bson:"validTill,omitempty" json:"validTill,omitempty"`
	VisaPlaceOfIssue          string        `bson:"visaPlaceOfIssue,omitempty" json:"visaPlaceOfIssue,omitempty"`
	VisaDateOfIssue           string        `bson:"visaDateOfIssue,omitempty" json:"visaDateOfIssue,omitempty"`
	VisaType                  string        `bson:"visaType,omitempty" json:"visaType,omitempty"`
	PassportNumber            string        `bson:"passportNumber,omitempty" json:"passportNumber,omitempty"`
	PassportPlaceOfIssue      string        `bson:"passportPlaceOfIssue,omitempty" json:"passportPlaceOfIssue,omitempty"`
	PassportDateOfIssue       string        `bson:"passportDateOfIssue,omitempty" json:"passportDateOfIssue,omitempty"`
	PassportExpiryDate        string        `bson:"passportExpiryDate,omitempty" json:"passportExpiryDate,omitempty"`
	AddressCountryOfResidence string        `bson:"addressCountryOfResidence,omitempty" json:"addressCountryOfResidence,omitempty"`
	AddressCurrent            string        `bson:"addressCurrent,omitempty" json:"addressCurrent,omitempty"`
	DateOfArrivalInIndia      string        `bson:"dateOfArrivalInIndia,omitempty" json:"dateOfArrivalInIndia,omitempty"`
	EmployedInIndia           string        `bson:"employedInIndia,omitempty" json:"employedInIndia,omitempty"`
	PurposeOfVisit            string        `bson:"purposeOfVisit,omitempty" json:"purposeOfVisit,omitempty"`
	NextDestination           string        `bson:"nextDestination,omitempty" json:"nextDestination,omitempty"`
	IndianContactNumber       string        `bson:"indianContactNumber,omitempty" json:"indianContactNumber,omitempty"`
	IndianMobileNumber        string        `bson:"indianMobileNumber,omitempty" json:"indianMobileNumber,omitempty"`
	PermanentMobileNumber     string        `bson:"permanentMobileNumber,omitempty" json:"permanentMobileNumber,omitempty"`
}
