package web

import (
	"fmt"
)

type Guide struct {
	Page
	Note string `xml:"note"`
}

func (r *Guide) Format() string {
	return fmt.Sprintf(
		"*%s*\n "+
			"%s"+
			"%s\n ",
		r.Title,
		r.Note,
		r.Body)
}