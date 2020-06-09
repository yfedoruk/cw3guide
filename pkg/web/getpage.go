package web

import (
	"encoding/xml"
	"github.com/yfedoruck/cw3guide/pkg/env"
	"github.com/yfedoruck/cw3guide/pkg/fail"
	"os"
	"path/filepath"
)

type Cw3Page interface {
	Format() string
}

func GetPage(name string) string {
	var page Cw3Page
	page = &Page{}
	switch name {
	case "guide", "help", "start", "man":
		name = "guide"
		page = &Guide{}
	default:
		if IsNotExist(name) {
			name = "notfound"
		}
	}

	parse(name, page)
	return page.Format()
}

func parse(filename string, r interface{}) {
	file, err := os.Open(env.BasePath() + filepath.FromSlash("/data/"+filename+".xml"))
	fail.Check(err)

	defer func() {
		var err = file.Close()
		fail.Check(err)
	}()

	fi, err := file.Stat()
	fail.Check(err)

	var data = make([]byte, fi.Size())
	_, err = file.Read(data)
	fail.Check(err)

	err = xml.Unmarshal(data, r)
	fail.Check(err)
}

func IsNotExist(filename string) bool {
	if _, err := os.Stat(env.BasePath() + filepath.FromSlash("/data/"+filename+".xml")); os.IsNotExist(err) {
		return true
	}
	return false
}

func ImagePath(filename string) string {
	return env.BasePath() + filepath.FromSlash("/data/img/"+filename+".png")
}

func IsPhoto(command string) bool {
	switch command {
	case "herbsimg", "recipesimg":
		return true
	}
	return false
}