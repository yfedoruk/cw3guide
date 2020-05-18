package web

import (
	"encoding/xml"
	"fmt"
)

type Page struct {
	XMLName xml.Name `xml:"root"`
	Title   string   `xml:"title"`
	Body    string   `xml:"body"`
}

func (r *Page) Format() string {
	return fmt.Sprintf(
		"*%s*\n "+
			"%s\n "+
			"🏰 Гайд /guide",
		r.Title,
		r.Body)
}
