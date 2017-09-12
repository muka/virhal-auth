package model

//Application store a service deployment
type Application struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
