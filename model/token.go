package model

import "time"

//Token for authentication
type Token struct {
	Expires time.Time
	Value   string
	Secret  string
}
