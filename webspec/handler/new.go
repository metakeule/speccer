package handler

import (
	"fmt"
	"github.com/metakeule/speclib"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

type Saver interface {
	Save(*speclib.Spec) error
}

type Loader interface {
	Load(*speclib.Spec) error
}

type handler struct {
	*http.ServeMux
	SpecLock *sync.RWMutex
	Spec     *speclib.Spec
	Saver    Saver
	Loader   Loader
}

func (h *handler) mountAll() {
	h.ServeMux.HandleFunc("/spec.html", h.index)
	h.ServeMux.HandleFunc("/spec.md", h.markdown)
	h.ServeMux.Handle("/admin", h.adminlayout())
	h.ServeMux.HandleFunc("/translations.json", h.serverTranslations)
	h.ServeMux.HandleFunc("/spec.json", h.specHandler)
	h.ServeMux.HandleFunc("/saveSection", h.saveSection)
	h.ServeMux.HandleFunc("/moveParagraph", h.moveParagraph)
	h.ServeMux.HandleFunc("/changeParagraphSection", h.changeParagraphSection)
	h.ServeMux.HandleFunc("/saveInfo", h.saveInfo)
	h.ServeMux.HandleFunc("/deleteSection", h.deleteSection)
	h.ServeMux.HandleFunc("/validate", h.validate)
	h.ServeMux.Handle("/public/", fileserver)
}

type FileSaver string

func (fs FileSaver) Save(s *speclib.Spec) error {
	return ioutil.WriteFile(string(fs), []byte(s.Json()), 0644)
}

func (h *handler) Save() {
	h.Spec.Update()
	err := h.Saver.Save(h.Spec)
	if err != nil {
		fmt.Printf("Can't save spec with uuid %s: %s\n", h.Spec.OVERVIEW.UUID, err.Error())
		return
	}
	err = h.Loader.Load(h.Spec)
	if err != nil {
		fmt.Printf("Can't reload spec with uuid %s: %s\n", h.Spec.OVERVIEW.UUID, err.Error())
	}
}

type StringLoader string

func (sl StringLoader) Load(s *speclib.Spec) error {
	return s.LoadJson(string(sl))
}

type FileLoader string

func (fl FileLoader) Load(s *speclib.Spec) error {
	bytes, err := ioutil.ReadFile(string(fl))
	if err != nil {
		return err
	}
	return StringLoader(string(bytes)).Load(s)
}

//func (h *handler) LoadJson(json string) error { return h.Spec.LoadJson(json) }

func (h *handler) Load() {
	err := h.Loader.Load(h.Spec)
	if err != nil {
		panic(err.Error())
	}
}

var fileserver http.Handler

func init() {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		panic("environment variable GOPATH missing")
	}
	dir := filepath.Join(gopath, "src", "github.com/metakeule/speccer/webspec")
	fileserver = http.FileServer(http.Dir(dir))
}

func New(l Loader, s Saver) *handler {
	h := &handler{}
	h.ServeMux = http.NewServeMux()
	h.SpecLock = &sync.RWMutex{}
	h.Spec = &speclib.Spec{Sections: map[string][]speclib.Paragraph{}}
	h.Saver = s
	h.Loader = l
	h.Load()
	h.mountAll()
	return h
}
