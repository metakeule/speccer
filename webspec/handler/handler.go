package handler

import (
	"encoding/json"
	"fmt"

	"gopkg.in/go-on/go.uuid.v1"
	. "gopkg.in/go-on/lib.v2/html"
	. "gopkg.in/go-on/lib.v2/types"
	// . "gopkg.in/go-on/lib.v2/html/h"
	// . "gopkg.in/go-on/lib.v2/html/tag"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/metakeule/speclib"
)

func (h *handler) markdown(rw http.ResponseWriter, req *http.Request) {
	h.SpecLock.RLock()
	defer h.SpecLock.RUnlock()
	usage := req.URL.Query().Get("usage")
	md := h.Spec.Markdown(filterForUsage(usage))
	rw.Header().Set("Content-Type", "text/x-markdown; charset=UTF-8")
	rw.Write([]byte(translate(h.Spec.INFO.Language, md)))
}

func (h *handler) saveInfo(rw http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		writeError(rw, err)
		return
	}

	h.SpecLock.Lock()
	defer h.SpecLock.Unlock()

	i := speclib.Info{}
	err = json.Unmarshal(b, &i)
	if err != nil {
		writeError(rw, err)
		return
	}
	h.Spec.INFO = i
	h.Save()
	rw.Write([]byte("ok"))
}

func (h *handler) validate(rw http.ResponseWriter, req *http.Request) {
	h.SpecLock.RLock()
	defer h.SpecLock.RUnlock()
	err := h.Spec.Validate(true)
	if err != nil {
		fmt.Fprintf(rw, "Error: %s\n", err.Error())
		return
	}
	rw.Write([]byte("OK"))
}

func (h *handler) changeParagraphSection(rw http.ResponseWriter, req *http.Request) {
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

	if position > len(h.Spec.Sections[sec])-1 {
		rw.WriteHeader(400)
		rw.Write([]byte("not found"))
		return
	}
	h.SpecLock.Lock()
	defer h.SpecLock.Unlock()
	p := h.Spec.Sections[sec][position]

	p.Update()
	h.Spec.AddParagraph(targetsection, &p)
	h.Spec.RemoveParagraph(sec, position)
	h.Spec.Update()
	h.Save()
	rw.Write([]byte("ok"))
}

func (h *handler) moveParagraph(rw http.ResponseWriter, req *http.Request) {
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
	h.SpecLock.Lock()
	defer h.SpecLock.Unlock()
	err := h.Spec.MoveParagraph(sec, from, to)
	if err != nil {
		writeError(rw, err)
		return
	}
	h.Spec.Update()
	h.Save()
	rw.Write([]byte("ok"))
}

func (h *handler) saveSection(rw http.ResponseWriter, req *http.Request) {
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
	h.SpecLock.Lock()
	defer h.SpecLock.Unlock()
	p := speclib.Paragraph{}
	err = json.Unmarshal(b, &p)
	if err != nil {
		writeError(rw, err)
		return
	}
	p.Spec = h.Spec

	p.Update()

	if pos == -1 {
		if p.Comments == nil {
			p.Comments = map[string]string{}
		}
		if sec == "OVERVIEW" {
			//h.Spec.Paragraph = &p
			h.Spec.OVERVIEW = &p
		} else {
			p.UUID = uuid.NewV4().String()
			h.Spec.AddParagraph(sec, &p)
		}
		h.Save()
		rw.Write([]byte("ok"))
		return
	}
	h.Spec.Sections[sec][pos] = p
	h.Save()
	rw.Write([]byte("ok"))
}

func (h *handler) deleteSection(rw http.ResponseWriter, req *http.Request) {
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
	h.SpecLock.Lock()
	defer h.SpecLock.Unlock()
	h.Spec.RemoveParagraph(sec, pos)
	h.Save()
	rw.Write([]byte("ok"))
}

func (h *handler) specHandler(rw http.ResponseWriter, req *http.Request) {
	h.SpecLock.RLock()
	defer h.SpecLock.RUnlock()
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.Write(j(h.Spec))
}

func (h *handler) index(rw http.ResponseWriter, req *http.Request) {
	theme := req.URL.Query().Get("theme")
	if theme == "" {
		theme = "cerulean"
	}

	usage := req.URL.Query().Get("usage")

	if usage == "" {
		usage = "documentation"
	}
	hideNavbar := req.URL.Query().Get("hide-navbar")

	h.SpecLock.RLock()
	defer h.SpecLock.RUnlock()
	body := BODY(
		themeSelect(theme), h.usageSelect(usage),
		xmp(
			Attrs_("theme", theme),
			Style{"display", "none"},
			translate(h.Spec.INFO.Language, h.Spec.Markdown(filterForUsage(usage))),
		),
		JsSrc("public/js/strapdown.min.js"),
	)
	if hideNavbar != "" {
		body.Add(Class("hide-navbar"))
	}
	HTML5(
		HTML(
			HEAD(
				TITLE(h.Spec.OVERVIEW.Title),
				JsSrc("public/js/jquery-1.10.2.min.js"),
				JsSrc("public/js/themes/bootstrap.min.js"),
				JsSrc("public/js/theme.js"),
			),
			body,
		),
	).WriteTo(rw)
}

func (h *handler) serverTranslations(rw http.ResponseWriter, req *http.Request) {
	h.SpecLock.RLock()
	defer h.SpecLock.RUnlock()
	tr := translations[h.Spec.INFO.Language]
	rw.Write(j(tr))
}
