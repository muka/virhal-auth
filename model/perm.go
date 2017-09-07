package model

//Permission informations
type Permission string

const (
	//PermissionAdmin admin permission
	PermissionAdmin Permission = "admin"
	//PermissionPatient patient permission
	PermissionPatient Permission = "patient"
	//PermissionDoctor doctor permission
	PermissionDoctor Permission = "doctor"
	//PermissionGuardian guardian permission
	PermissionGuardian Permission = "guardian"
	//PermissionNextOfKeen next-of-keen permission
	PermissionNextOfKeen Permission = "next-of-keen"
)
