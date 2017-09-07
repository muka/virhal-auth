package model

import (
	"time"

	"github.com/satori/go.uuid"
)

// ServiceEnabled list enabled services
type ServiceEnabled struct {
	Atlante   bool `json:"atlante"`
	Rtc       bool `json:"rtc"`
	Games     bool `json:"games"`
	Phr       bool `json:"phr"`
	Location  bool `json:"location"`
	Reminders bool `json:"reminders"`
	Cube3D    bool `json:"3Dcube"`
}

// Lang language code
type Lang string

const (
	//LangIt it language
	LangIt Lang = "it"
	//LangEn en language
	LangEn Lang = "en"
	//LangFr fr language
	LangFr Lang = "fr"
	//LangDe de language
	LangDe Lang = "de"
	//LangEl el language
	LangEl Lang = "el"
	//LangMk mk language
	LangMk Lang = "mk"
	//LangRo ro language
	LangRo Lang = "ro"
	//LangSl sl language
	LangSl Lang = "sl"
)

//Sso token
type Sso struct {
	Atlante   string `json:"atlante"`
	Biophr    string `json:"biophr"`
	Chino     string `json:"chino"`
	Raptor    string `json:"raptor"`
	FitForAll string `json:"fitforall"`
	Trilogis  string `json:"trilogis"`
	Webrtc    string `json:"webrtc"`
	Cube3D    string `json:"3dcube"`
}

// GeneralPreferences general preferences storage
type GeneralPreferences struct {
	FontSize    string `json:"fontSize"`
	ColorScheme string `json:"colorScheme"`
	Environment string `json:"environment"`
}

// User profile informations
type User struct {
	ID             string         `json:"id"`
	UserID         string         `json:"userId"`
	Username       string         `json:"username"`
	FirstName      string         `json:"firstName"`
	LastName       string         `json:"lastName"`
	UserType       string         `json:"userType"`
	DateOfBirth    time.Time      `json:"dateOfBirth"`
	Email          string         `json:"email"`
	Service        ServiceEnabled `json:"service"`
	Lang           Lang           `json:"lang"`
	SessionToken   string         `json:"sessionToken"`
	ContactPhone   string         `json:"contactPhone"`
	NextOfKinName  string         `json:"nextOfKinName"`
	GeneralRemarks string         `json:"generalRemarks"`
	MedicalRemarks string         `json:"medicalRemarks"`
	ServiceID      int            `json:"pilotId"`
	AssignedUsers  []User         `json:"assigned_users"`
	AssignedDoctor []User         `json:"assigned_doctor"`
}

//NewUser init an user
func NewUser() User {
	return User{
		AssignedDoctor: make([]User, 0),
		AssignedUsers:  make([]User, 0),
		ID:             uuid.NewV4().String(),
		Lang:           LangEn,
		Service: ServiceEnabled{
			Atlante:   true,
			Rtc:       true,
			Games:     true,
			Phr:       true,
			Location:  true,
			Reminders: true,
			Cube3D:    true,
		},
	}
}