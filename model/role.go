package model

// Role informations
type Role struct {
	Name        string `binding:"required"`
	Permissions []Permission
}

//RolePatient patient role
var RolePatient = Role{"patient", []Permission{}}

//RoleDoctor doctor role
var RoleDoctor = Role{"doctor", []Permission{}}

//RoleGuardian guardian role
var RoleGuardian = Role{"guardian", []Permission{}}

//RoleNextOfKeen next-of-keen role
var RoleNextOfKeen = Role{"next-of-keen", []Permission{}}
