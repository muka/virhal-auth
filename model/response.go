package model

//ResponseLogin a response login
type ResponseLogin struct {
	StatusCode         string             `json:"statusCode"`
	UserInfo           User               `json:"userInfo"`
	GeneralPreferences GeneralPreferences `json:"generalPreferences"`
	Sso                Sso                `json:"sso"`
}

//NewResponseLogin init a ResponseLogin
func NewResponseLogin(user *User) ResponseLogin {

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
		Sso:        Sso{},
		StatusCode: statusCode,
		UserInfo:   *user,
	}
}
