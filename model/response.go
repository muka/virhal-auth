package model

//ResponseLogin a response login
type ResponseLogin struct {
	StatusCode         string             `json:"statusCode"`
	UserInfo           PublicUser         `json:"userInfo"`
	GeneralPreferences GeneralPreferences `json:"generalPreferences"`
	Sso                Sso                `json:"sso"`
}

//NewResponseLogin init a ResponseLogin
func NewResponseLogin(user *User, sso Sso) ResponseLogin {

	statusCode := "ok"
	if user == nil {
		statusCode = "error"
	}
	return ResponseLogin{
		GeneralPreferences: GeneralPreferences{
			ColorScheme: "highContrast",
			Environment: "shared",
			FontSize:    "200",
		},
		Sso:        sso,
		StatusCode: statusCode,
		UserInfo:   user.ToPublicUser(),
	}
}
