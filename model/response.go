package model

//ResponseLogin a response login
type ResponseLogin struct {
	PublicUser
}

//NewResponseLogin init a ResponseLogin
func NewResponseLogin(user *User) ResponseLogin {
	return ResponseLogin{
		PublicUser: user.ToPublicUser(),
	}
}
