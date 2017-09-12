package model

// Role informations
type Role struct {
	Name        string       `json:"name" binding:"required"`
	Permissions []Permission `json:"permissions,omitempty"`
}
