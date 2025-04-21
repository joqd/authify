package domain

type Group struct {
	ID          int
	Name        string
	Permissions []Permission
}
