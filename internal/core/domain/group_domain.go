package domain

type Group struct {
	ID          uint
	Name        string
	Permissions []Permission
}
