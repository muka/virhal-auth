package model

// RequestLogin contains login request
type RequestLogin struct {
	Username string `json:"username" binding:"required,min=3,max=64"`
	Password string `json:"password" binding:"required,min=3"`
}

// RequestRegister contains login request
type RequestRegister struct {
	Username string `json:"username" binding:"required,min=3,max=64"`
	Password string `json:"password" binding:"required,min=3"`
	Email    string `json:"email" binding:"required,min=4,email"`
}

// ToUser convert to a User model
func (r *RequestRegister) ToUser() User {
	u := NewUser()
	u.Username = r.Username
	u.Password = r.Password
	u.Email = r.Email
	return u
}
