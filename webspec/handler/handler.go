package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-contrib/uuid"
	"github.com/metakeule/speclib"
	"io/ioutil"
	"net/http"
	"strconv"
)

func saveInfo(rw http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		writeError(rw, err)
		return
	}

	i := speclib.Info{}
	err = json.Unmarshal(b, &i)
	if err != nil {
		writeError(rw, err)
		return
	}
	spec.INFO = i
	Save()
	rw.Write([]byte("ok"))
}

func validate(rw http.ResponseWriter, req *http.Request) {
	err := spec.Validate(true)
	if err != nil {
		fmt.Fprintf(rw, "Error: %s\n", err.Error())
		return
	}
	rw.Write([]byte("ok"))
}

func changeParagraphSection(rw http.ResponseWriter, req *http.Request) {
	sec := req.URL.Query().Get("section")

	if sec == "" {
		rw.WriteHeader(400)
		rw.Write([]byte("section missing"))
		return
	}
	targetsection := req.URL.Query().Get("targetsection")

	if sec == "" {
		rw.WriteHeader(400)
		rw.Write([]byte("targetsection missing"))
		return
	}
	position, e := strconv.Atoi(req.URL.Query().Get("position"))
	if e != nil {
		writeError(rw, e)
		return
	}

	if position > len(spec.Sections[sec])-1 {
		rw.WriteHeader(400)
		rw.Write([]byte("not found"))
		return
	}

	p := spec.Sections[sec][position]

	p.Update()
	spec.AddParagraph(speclib.SectionObj[targetsection], &p)
	spec.RemoveParagraph(speclib.SectionObj[sec], position)
	spec.Update()
	Save()
	rw.Write([]byte("ok"))
}

func moveParagraph(rw http.ResponseWriter, req *http.Request) {
	sec := req.URL.Query().Get("section")

	if sec == "" {
		rw.WriteHeader(400)
		rw.Write([]byte("section missing"))
		return
	}
	from, e := strconv.Atoi(req.URL.Query().Get("from"))
	if e != nil {
		writeError(rw, e)
		return
	}

	to, e := strconv.Atoi(req.URL.Query().Get("to"))
	if e != nil {
		writeError(rw, e)
		return
	}

	err := spec.MoveParagraph(speclib.SectionObj[sec], from, to)
	if err != nil {
		writeError(rw, err)
		return
	}
	spec.Update()
	Save()
	rw.Write([]byte("ok"))
}

func saveSection(rw http.ResponseWriter, req *http.Request) {
	sec := req.URL.Query().Get("section")

	if sec == "" {
		rw.WriteHeader(400)
		rw.Write([]byte("section missing"))
		return
	}
	pos, e := strconv.Atoi(req.URL.Query().Get("position"))
	if e != nil {
		writeError(rw, e)
		return
	}

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		writeError(rw, err)
		return
	}

	p := speclib.Paragraph{}
	err = json.Unmarshal(b, &p)
	if err != nil {
		writeError(rw, err)
		return
	}
	p.Spec = spec

	p.Update()

	if pos == -1 {
		if p.Comments == nil {
			p.Comments = map[string]string{}
		}
		if sec == "OVERVIEW" {
			//spec.Paragraph = &p
			spec.OVERVIEW = &p
		} else {
			p.UUID = uuid.NewV4().String()
			spec.AddParagraph(speclib.SectionObj[sec], &p)
		}
		Save()
		rw.Write([]byte("ok"))
		return
	}
	spec.Sections[sec][pos] = p
	Save()
	rw.Write([]byte("ok"))
}

func deleteSection(rw http.ResponseWriter, req *http.Request) {
	pos, e := strconv.Atoi(req.URL.Query().Get("position"))
	if e != nil {
		writeError(rw, e)
		return
	}
	sec := req.URL.Query().Get("section")

	if sec == "" {
		rw.WriteHeader(400)
		rw.Write([]byte("section missing"))
		return
	}

	spec.RemoveParagraph(speclib.SectionObj[sec], pos)
	Save()
	rw.Write([]byte("ok"))
}

func specHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write(j(spec))
}
