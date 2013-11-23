package handler

import (
	"github.com/metakeule/speclib"
	"net/http"
	"os"
	"path/filepath"
)

var spec = &speclib.Spec{Sections: map[string][]speclib.Paragraph{}}

func New(json string) *http.ServeMux {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		panic("environment variable GOPATH")
	}
	dir := filepath.Join(gopath, "src", "github.com/metakeule/speccer/webspec")
	fileserver := http.FileServer(http.Dir(dir))
	err := spec.LoadJson(json)
	if err != nil {
		panic(err.Error())
	}
	mux := http.NewServeMux()
	mux.Handle("/", layout)
	mux.HandleFunc("/spec.json", specHandler)
	mux.HandleFunc("/saveSection", saveSection)
	mux.HandleFunc("/moveParagraph", moveParagraph)
	mux.HandleFunc("/changeParagraphSection", changeParagraphSection)
	mux.HandleFunc("/saveInfo", saveInfo)
	mux.HandleFunc("/deleteSection", deleteSection)
	mux.HandleFunc("/validate", validate)
	mux.Handle("/public/", fileserver)
	return mux
}
