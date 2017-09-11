package model

import (
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
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
	ObjectID bson.ObjectId `json:"-" bson:"_id"`
	ID       string        `json:"id"`
	Enabled  bool          `json:"enabled"`
	Username string        `json:"username" binding:"required,min=3,max=64"`
	Password string        `json:"password,omitempty" binding:"required,min=3"`
	Email    string        `json:"email" binding:"required,min=4,email"`

	UserID         string         `json:"userId,omitempty" binding:"required,uuid4"`
	UserType       string         `json:"userType,omitempty" binding:"required"`
	Roles          []Role         `json:"roles" binding:"required"`
	FirstName      string         `json:"firstName,omitempty"`
	LastName       string         `json:"lastName,omitempty"`
	DateOfBirth    string         `json:"dateOfBirth,omitempty"`
	Service        ServiceEnabled `json:"service"`
	Lang           Lang           `json:"lang,omitempty"`
	SessionToken   string         `json:"sessionToken,omitempty"`
	ContactPhone   string         `json:"contactPhone,omitempty"`
	NextOfKinName  string         `json:"nextOfKinName,omitempty"`
	GeneralRemarks string         `json:"generalRemarks,omitempty"`
	MedicalRemarks string         `json:"medicalRemarks,omitempty"`
	ServiceID      int            `json:"pilotId,omitempty"`
	AssignedUsers  []User         `json:"assigned_users"`
	AssignedDoctor []User         `json:"assigned_doctor"`
}

//ToPublicUser return a specialized struct for serialization
func (u *User) ToPublicUser() PublicUser {
	p := PublicUser{
		User: u,
	}
	return p
}

type omit *struct{}

//PublicUser expose defined fields on JSON marshalling
type PublicUser struct {
	*User
	ObjectID omit `json:"-"`
	Password omit `json:"password,omitempty"`
}

//NewUser init an user
func NewUser() User {
	return User{
		AssignedDoctor: make([]User, 0),
		AssignedUsers:  make([]User, 0),
		Roles:          make([]Role, 0),
		ID:             uuid.NewV4().String(),
		ObjectID:       bson.NewObjectId(),
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
