package model

type OutgoingEmail struct {
	From        string
	To          string
	Subject     string
	Text        string
	DSN         string
	Attachments []string
}
