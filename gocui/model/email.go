package model

type Email struct {
	Id int
	// Email's path on the filesystem
	Path        string
	From        string
	To          string
	DeliveredTo string
	Subject     string
	Text        string
	HTML        string
	Date        string

	// Email flags
	IsSeen      bool
	IsImportant bool
	IsAnswered  bool
	IsSelected  bool
}
