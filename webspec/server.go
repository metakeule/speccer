package main

import (
	"encoding/json"
	"fmt"
	"github.com/metakeule/goh4"
	. "github.com/metakeule/goh4/tag"
	. "github.com/metakeule/goh4/tag/short"
	"github.com/metakeule/ng"
	"github.com/metakeule/speclib"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func dropDownSection(section string) *goh4.Element {
	return LI(
		CLASS("dropdown"),
		AHref("#", ID(section+"-drop"),
			CLASS("dropdown-toggle"),
			ATTR("data-toggle", "dropdown"), section,
			B(CLASS("caret")),
		),
		UL(CLASS("dropdown-menu"), ATTR("role", "menu", "aria-labelledby", section+"-drop"),
			LI(
				A(
					ng.Href(ng.Template("#"+section+"-new")),
					ATTR("data-toggle", "tab"),
					"new"),
			),
			LI(
				ng.Repeat("s", section+"S"),
				A(
					ATTR("data-toggle", "tab"),
					ng.Href(ng.Template("#"+section)),
					"{{s.Title}}",
					ng.Click(ng.E("setContent('"+section+"')"))),
			),
		),
	)
}

func input(section string, fieldName string, i ...interface{}) *goh4.Element {
	i = append(i,
		ng.Model(section+"."+fieldName),
		CLASS("form-control"),
		ID(fieldName))
	//ATTR("value", "{{"+tempvar+"."+fieldName+"}}"))
	return DIV(
		CLASS("form-group"),
		LabelFor(fieldName, fieldName),
		InputText(fieldName, i...),
	)
}

func textarea(section string, fieldName string, i ...interface{}) *goh4.Element {
	i = append(i,
		ng.Model(section+"."+fieldName),
		CLASS("form-control"),
		ID(fieldName),
		ATTR("rows", "10", "name", fieldName))
	return DIV(CLASS("form-group"),
		LabelFor(fieldName, fieldName),
		TEXTAREA(i...),
	)
}

/*
"Responsible": "mra",
        "Title": "Rights Management vs responsible",
        "Text": "Rights Management is a Non Goal but on the other hand we have\npersons responsible for paragraphs. \n\nIf anyone may edit any paragraph, then the responsibility\nmust constantly change (and is worthless) or it is false.\n",
        "Comments": {"mra": "Yes the responsability should change often in order to request feedback.\nIf multiple persons work on a spec, the spec is considered teamwork\nand the people should work together.\n\nOnly one person is responsible for each paragraph at a single point in\ntime. This could be enforced by external tools as well as rights management\nif necessary. Although with version control it should not be necessary."}
        "LastUpdate": "2013-11-12",
        "State": "OBSOLET",
        "Deadline": "2013-11-11",
        "EstimatedHours": 1
*/
func contentSection(section string) *goh4.Element {
	return Doc(
		DIV(
			BR(),
			ATTR("role", "form"),
			CLASS("tab-pane"),
			ID(section+"-new"),
			FORM(
				ng.Submit(ng.E("createSection('"+section+"')")),
				//ng.Form(section+"-new"),
				InputHidden("section", ATTR("value", section)),
				DIV(CLASS("form-group"),
					LabelFor("title", "Title"),
					InputText("title", CLASS("form-control"), ID("title"), ATTR("value", "")),
				),
				DIV(CLASS("form-group"),
					LabelFor("title", "Text"),
					TEXTAREA(CLASS("form-control"), ATTR("rows", "10"), HTML("")),
				),
				BUTTON(ATTR("type", "submit"), CLASS("btn"), CLASS("btn-default"), "Save"),
			),
		),
		DIV(
			BR(),
			//ng.Repeat("s", section),
			ATTR("role", "form"),
			CLASS("tab-pane"),
			ID(section),
			FORM(
				ng.Submit(ng.E("updateSection('"+section+"')")),
				//InputHidden("section", ID("Section"), ATTR("value", section)),
				//	ng.Form(section+"{{s.Position}}"),
				input(section, "LastUpdate"),
				input(section, "State"),
				input(section, "Deadline"),
				input(section, "EstimatedHours"),
				input(section, "Title"),
				textarea(section, "Text"),
				/*
					DIV(CLASS("form-group"),
						LabelFor("State", "State"),
						InputText("State",
							ng.Model(section+".State"), CLASS("form-control"), ID("State"), ATTR("value", "{{s.State}}")),
					),
					DIV(CLASS("form-group"),
						LabelFor("Deadline", "Deadline"),
						InputText("Deadline",
							ng.Model(section+".Deadline"), CLASS("form-control"), ID("Deadline"), ATTR("value", "{{s.Deadline}}")),
					),
					DIV(CLASS("form-group"),
						LabelFor("EstimatedHours", "EstimatedHours"),
						InputText("EstimatedHours", ng.Model(section+".EstimatedHours"), CLASS("form-control"), ID("EstimatedHours"), ATTR("value", "{{s.EstimatedHours}}")),
					),
					DIV(CLASS("form-group"),
						LabelFor("title", "Title"),
						InputText("title", ng.Model(section), CLASS("form-control"), ID("title"), ATTR("value", "{{s.Title}}")),
					),
					DIV(CLASS("form-group"),
						LabelFor("title", "Text"),
						TEXTAREA(CLASS("form-control"), ng.Model(section), ATTR("rows", "10", "name", "comment"), HTML("{{s.Text}}")),
					),
				*/
				BUTTON(ATTR("type", "submit"), CLASS("btn"), CLASS("btn-default"), "Save"),
				A(goh4.Style{"float", "right"},
					ng.Href(ng.Template("#{{"+section+".Position}}-delete")),
					"delete", CLASS("btn"), CLASS("btn-danger")),

				DIV(
					H2("COMMENTS"),
					DIV(
						A(
							ng.Href(ng.Template("#{{"+section+".Position}}-comment-new")),
							"new", CLASS("btn"), CLASS("btn-default")),
					),
					DIV(
						ATTR("ng-repeat", "(author, c) in "+section+".Comments"),
						H3("{{author}}"),
						PRE("{{c}}"),
						A(
							ng.Href(ng.Template("#{{"+section+".Position}}-comment-{{author}}")),
							"edit", CLASS("btn"), CLASS("btn-default")),
					),
				),
			),
		),
	)
}

func updateSection(rw http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)

	//err := req.ParseForm()
	if err != nil {
		rw.Write([]byte(err.Error()))
		return
	}

	v := &Paragraph{}
	err = json.Unmarshal(b, v)
	if err != nil {
		rw.Write([]byte(err.Error()))
		return
	}

	pos, e := strconv.Atoi(req.URL.Query().Get("position"))
	if e != nil {
		rw.Write([]byte(e.Error()))
		return
	}
	sec := req.URL.Query().Get("section")

	//spec.

	p := spec.Sections[sec][pos]
	p.EstimatedHours = v.EstimatedHours
	p.Deadline = v.Deadline
	p.LastUpdate = v.LastUpdate
	p.Responsible = v.Responsible
	p.State = v.State
	p.Text = v.Text
	p.Title = v.Title
	spec.Sections[sec][pos] = p
	//ps[pos] = p
	//spec.Sections[sec] = ps
	//fmt.Printf("section: %s \nposition: %d\n %#v\n", sec, pos, *(v.Paragraph))
	Save()
	rw.Write([]byte("ok"))
}

var layout = HTML5(
	ATTR("ng-app", "app"),
	HEAD(
		CharsetUtf8(),
		META(ATTR("name", "viewport", "content", "width=device-width, initial-scale=1.0")),
		HTML(`
			<!--[if lt IE 9]>
      			<script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
      			<script src="https://oss.maxcdn.com/libs/respond.js/1.3.0/respond.min.js"></script>
    		<![endif]-->
    	`),
		CssHref("/css/bootstrap.css"),
		CssHref("/css/app.css"),
		JsSrc("/js/jquery-1.10.2.min.js"),
		JsSrc("/angular-1.2.1/angular.js"),
		JsSrc("/js/bootstrap.js"),
		JsSrc("/js/app.js"),
	),

	BODY(
		ATTR("ng-controller", "Spec"),
		DIV(CLASS("container"),

			H1("SPECS"),

			UL(CLASS("nav"), CLASS("nav-tabs"),
				LI(AHref("#info", ATTR("data-toggle", "tab"), "INFO")),
				LI(AHref("#overview", ATTR("data-toggle", "tab"), "OVERVIEW")),
				dropDownSection("SCENARIO"),
				dropDownSection("NONGOAL"),
				dropDownSection("UNDECIDED"),
				dropDownSection("DEFINITION"),
				dropDownSection("FEATURE"),
			),

			DIV(CLASS("tab-content"),
				DIV(CLASS("tab-pane"), CLASS("active"), ID("info"), "info"),
				DIV(CLASS("tab-pane"), ID("overview"), "overview"),
				contentSection("SCENARIO"),
				contentSection("NONGOAL"),
				contentSection("UNDECIDED"),
				contentSection("DEFINITION"),
				contentSection("FEATURE"),
			),
			//DIV(ATTR("ng-view", "ng-view")),
		),
	),
)

type Paragraph struct {
	*speclib.Paragraph
	Position int
	Section  string
}

func mkSectionJson(section string) []byte {
	ps := spec.Sections[section]
	psWithPos := make([]*Paragraph, len(ps))
	for pos, _ := range ps {
		psWithPos[pos] = &Paragraph{&ps[pos], int(pos), section}
	}
	res, err := json.Marshal(psWithPos)
	if err != nil {
		panic(err.Error())
	}
	return res
}

func specHandler(rw http.ResponseWriter, req *http.Request) {
	section := req.URL.Query().Get("section")
	switch section {
	case "":
	case "overview":
	case "FEATURE":
		rw.Write(mkSectionJson("FEATURE"))
	case "DEFINITION":
		rw.Write(mkSectionJson("DEFINITION"))
	case "SCENARIO":
		rw.Write(mkSectionJson("SCENARIO"))
	case "NONGOAL":
		rw.Write(mkSectionJson("NONGOAL"))
	case "UNDECIDED":
		rw.Write(mkSectionJson("UNDECIDED"))
	}
	/*
		dropDownSection("scenarios", "SCENARIOS"),
					dropDownSection("nongoals", "NONGOALS"),
					dropDownSection("undecided", "UNDECIDED"),
					dropDownSection("definitions", "DEFINITIONS"),
					dropDownSection("features", "FEATURES"),
	*/
}

var spec = &speclib.Spec{}

func Save() {
	// fmt.Println(spec.Json())
	err := ioutil.WriteFile(os.Args[1], []byte(spec.Json()), 0644)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

func main() {
	fileserver := http.FileServer(http.Dir(`/home/benny/neue_platte/benny/Entwicklung/gopath/src/github.com/metakeule/speccer/webspec/public`))
	if len(os.Args) != 2 {
		fmt.Println("no file given")
		os.Exit(1)
	}
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err.Error())
	}
	//fmt.Printf("%s\n", file)
	err = spec.LoadJson(fmt.Sprintf("%s", file))
	if err != nil {
		panic(err.Error())
	}

	http.Handle("/", layout)
	http.HandleFunc("/spec.json", specHandler)
	http.HandleFunc("/updateSection", updateSection)
	http.Handle("/css/", fileserver)
	http.Handle("/js/", fileserver)
	http.Handle("/img/", fileserver)
	http.Handle("/tpl/", fileserver)
	http.Handle("/fonts/", fileserver)
	http.Handle("/angular-1.2.1/", fileserver)
	http.ListenAndServe(":8181", nil)
}
