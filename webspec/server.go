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
	"strings"
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
					ATTR("data-toggle", "tab"),
					ng.Href("#section"),
					"new",
					ng.Click("newContent('"+section+"')")),
			),
			LI(
				ng.Repeat("s", "Sections['"+section+"']"),
				A(
					ATTR("data-toggle", "tab"),
					ng.Href("#section"),
					"{{s.Title}}",
					ng.Click("setContent('"+section+"')")),
			),
		),
	)
}

func input(fieldName string, label string, width int, i ...interface{}) *goh4.Element {
	i = append(i,
		ng.Model(fieldName),
		CLASS("form-control"),
		ID(fieldName))
	return DIV(
		CLASS("form-group"),
		CLASS(fmt.Sprintf("col-xs-%d", width)),
		LabelFor(fieldName, label),
		InputText(fieldName, i...),
	)
}

func static(label string, width int, i ...interface{}) *goh4.Element {
	i = append(i,
		CLASS("form-control-static"))
	return DIV(
		CLASS("form-group"),
		CLASS(fmt.Sprintf("col-xs-%d", width)),
		LABEL(label),
		P(i...),
	)
}

func selectbox(fieldName string, label string, width int, arraySrc string, i ...interface{}) *goh4.Element {
	i = append(i,
		ng.Model(fieldName),
		CLASS("form-control"),
		ID(fieldName), ATTR("ng-options", "v for v in "+arraySrc))

	return DIV(
		CLASS("form-group"),
		CLASS(fmt.Sprintf("col-xs-%d", width)),
		LabelFor(fieldName, label),
		SELECT(i...),
	)
}

func textarea(fieldName string, label string, width int, rows int, i ...interface{}) *goh4.Element {
	i = append(i,
		ng.Model(fieldName),
		CLASS("form-control"),
		ID(fieldName),
		ATTR("rows", fmt.Sprintf("%d", rows), "value", "{{fieldName}}", "name", fieldName))
	return DIV(CLASS("form-group"),
		CLASS(fmt.Sprintf("col-xs-%d", width)),
		LabelFor(fieldName, label),
		TEXTAREA(i...),
	)
}

func mapList(model, key, value string) *goh4.Element {
	var keyD, valD = strings.ToLower(key), strings.ToLower(value)
	return DIV(
		BR(),
		CLASS("row"),
		DIV(CLASS("form-group"), CLASS("col-xs-6"),
			LABEL(model),
			BR(),
			DIV(CLASS("btn-group"),
				goh4.Style{"margin-right", "15px"},
				ng.RepeatKeyVal(keyD, valD, "INFO."+model),
				BUTTON(
					ATTR("type", "button"),
					SPAN(CLASS("glyphicon"), CLASS("glyphicon-pencil")),
					HTML("&nbsp;"),
					ng.Click("set"+model+"("+keyD+", "+valD+")"), "{{"+keyD+"}}", CLASS("btn"), CLASS("btn-default"),
				),
				BUTTON(
					ATTR("type", "button"),
					goh4.Style{"float", "none"},
					ng.Click("remove"+model+"("+keyD+")"),
					//CLASS("close"),
					CLASS("btn"),
					CLASS("btn-danger"),
					ATTR("type", "button", "aria-hidden", "true"), HTML("&times;")),
			),
		),
		input(model+key, key, 2, ng.Blur("save"+model+"()")),
		input(model+value, value, 4, ng.Change("save"+model+"()")),
	)
}

func arrayList(model, value string) *goh4.Element {
	var valD = strings.ToLower(value)
	return DIV(
		BR(),
		CLASS("row"),
		DIV(CLASS("form-group"), CLASS("col-xs-6"),
			LABEL(model),
			BR(),
			DIV(CLASS("btn-group"),
				goh4.Style{"margin-right", "15px"},
				ng.Repeat(valD, "INFO."+model),
				BUTTON(
					ATTR("type", "button"),
					SPAN(CLASS("glyphicon"), CLASS("glyphicon-pencil")),
					HTML("&nbsp;"),
					"{{"+valD+"}}", CLASS("btn"), CLASS("btn-default"),
				),
				BUTTON(
					ATTR("type", "button"),
					goh4.Style{"float", "none"},
					ng.Click("remove"+model+"("+valD+")"),
					//CLASS("close"),
					CLASS("btn"),
					CLASS("btn-danger"),
					ATTR("type", "button", "aria-hidden", "true"), HTML("&times;")),
			),
		),
		input(model+value, value, 4, ng.Blur("add"+model+"()")),
	)
}

func contentInfo() *goh4.Element {
	return DIV(
		ID("info"), ATTR("role", "form"), CLASS("tab-pane"),
		BR(),
		FORM(
			CLASS("form-inline"),
			ng.Submit("saveInfo()"),
			DIV(CLASS("row"),
				input("INFO.Company", "Company", 6),
				input("INFO.Project", "Project", 6),
			),
			DIV(CLASS("row"),
				input("INFO.URL", "URL", 6),
				input("INFO.Parent", "Parent", 6),
			),
			DIV(CLASS("row"),
				input("INFO.Language", "Language", 4),
				input("INFO.DateFormat", "DateFormat", 4),
				selectbox("INFO.Approved", "Approved", 4, "[true,false]"),
			),
			arrayList("RequestedBy", "Name"),
			mapList("Translations", "Language", "URL"),
			mapList("Related", "Name", "URL"),
			mapList("SupersededBy", "Name", "URL"),
			mapList("Resources", "Name", "URL"),
			mapList("Persons", "Short", "Full"),

			DIV(CLASS("row"),
				BR(),
				DIV(CLASS("form-group"), CLASS("col-xs-6"),
					BUTTON(ATTR("type", "submit"), CLASS("btn"), CLASS("btn-primary"), "Save"),
				),
			),
		),
	)
}

func contentSection() *goh4.Element {
	return DIV(
		ID("section"), ATTR("role", "form"), CLASS("tab-pane"),
		BR(),
		FORM(
			CLASS("form-inline"),
			ng.Submit("saveSection()"),
			DIV(CLASS("row"),
				input("paragraph.Title", "Title", 6),
				input("paragraph.Responsible", "Responsible", 6),
				//input("paragraph.State", "State", 2),
				selectbox("paragraph.State", "State", 3, "states"),
				static("LastUpdate", 3, HTML("{{paragraph.LastUpdate}}")),
				//input("paragraph.LastUpdate", "LastUpdate", 2),
				input("paragraph.Deadline", "Deadline", 3),
				input("paragraph.EstimatedHours", "EstimatedHours", 3, ATTR("type", "number")),

				//Responsible
			),
			BR(),
			DIV(CLASS("row"),
				DIV(CLASS("form-group"),
					CLASS("col-xs-7"),
					LABEL("Text"),
					DIV(ID("Text"), CLASS("form-control"), goh4.Style{"height", "auto"}),
				),
				// ),
				// DIV(CLASS("row"),
				DIV(CLASS("form-group"),
					CLASS("col-xs-5"),
					LABEL("Preview"),
					DIV(ID("preview"), CLASS("form-control"), goh4.Style{"height", "300px"}, goh4.Style{"overflow-y", "scroll"}),
				),
				//textarea("paragraph.Text", "Text", 12, 6),
			),

			DIV(
				CLASS("row"),
				DIV(CLASS("form-group"), CLASS("col-xs-10"),
					BR(),
					LABEL("Comments"),
					BR(),
					DIV(CLASS("btn-group"),
						// SPAN(
						// goh4.Style{"float", "left"},
						//goh4.Style{"display", "inline-block"},
						goh4.Style{"margin-right", "15px"},
						ng.RepeatKeyVal("author", "comment", "paragraph.Comments"),
						BUTTON(
							ATTR("type", "button"),
							//<span class="glyphicon glyphicon-star"></span>
							SPAN(CLASS("glyphicon"), CLASS("glyphicon-pencil")),
							HTML("&nbsp;"),
							ng.Click("setComment(author, comment)"), "{{author}}", CLASS("btn"), CLASS("btn-default"),
						),
						BUTTON(
							ATTR("type", "button"),
							goh4.Style{"float", "none"},
							ng.Click("removeComment(author)"),
							//CLASS("close"),
							CLASS("btn"),
							CLASS("btn-danger"),
							ATTR("type", "button", "aria-hidden", "true"), HTML("&times;")),
					),
				),
			),
			DIV(
				CLASS("row"),
				//goh4.Style{"clear", "both"},
				BR(),
				input("CommentAuthor", "Author", 4, ng.Blur("saveComment()")),
				textarea("CommentText", "Comment", 12, 4, ng.Change("saveComment()")),
			),
			DIV(CLASS("row"),
				BR(),
				DIV(CLASS("form-group"), CLASS("col-xs-6"),
					BUTTON(ATTR("type", "submit"), CLASS("btn"), CLASS("btn-primary"), "Save"),
				),
				DIV(CLASS("form-group"), CLASS("col-xs-6"),
					A(goh4.Style{"float", "right"},
						ng.Hide("sectionIndex == -1"),
						ng.Click("deleteSection()"),
						"delete", CLASS("btn"), CLASS("btn-danger")),
				),
			),
		),
	)
}

func saveAll(rw http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rw.WriteHeader(400)
		rw.Write([]byte(err.Error()))
		return
	}

	err = json.Unmarshal(b, spec)
	if err != nil {
		rw.WriteHeader(400)
		rw.Write([]byte(err.Error()))
		return
	}
	Save()
	rw.Write([]byte("ok"))
}

func saveInfo(rw http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rw.WriteHeader(400)
		rw.Write([]byte(err.Error()))
		return
	}

	i := speclib.Info{}
	err = json.Unmarshal(b, &i)
	if err != nil {
		rw.WriteHeader(400)
		rw.Write([]byte(err.Error()))
		return
	}
	spec.INFO = i
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
		rw.WriteHeader(400)
		rw.Write([]byte(e.Error()))
		return
	}

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rw.WriteHeader(400)
		rw.Write([]byte(err.Error()))
		return
	}

	p := speclib.Paragraph{}
	err = json.Unmarshal(b, &p)
	if err != nil {
		rw.WriteHeader(400)
		rw.Write([]byte(err.Error()))
		return
	}
	p.Spec = spec

	if pos == -1 {
		if p.Comments == nil {
			p.Comments = map[string]string{}
		}
		if sec == "OVERVIEW" {
			//spec.Paragraph = &p
			spec.OVERVIEW = &p
		} else {
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
		rw.WriteHeader(400)
		rw.Write([]byte(e.Error()))
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
		CssHref("/codemirror-3.19/lib/codemirror.css"),
		JsSrc("/js/jquery-1.10.2.min.js"),
		JsSrc("/codemirror-3.19/lib/codemirror.js"),
		JsSrc("/codemirror-3.19/mode/markdown/markdown.js"),
		JsSrc("/angular-1.2.1/angular.min.js"),
		JsSrc("/angular-1.2.1/angular-route.min.js"),
		//JsSrc("/js/ui-codemirror.js"),
		JsSrc("/js/markdown.min.js"),
		JsSrc("/js/bootstrap.js"),
		JsSrc("/js/app.js"),
	),

	BODY(
		ATTR("ng-controller", "Spec"),
		DIV(CLASS("container"),

			H1("SPECS"),
			FORM(
				CLASS("form-inline"),
				ng.Submit("saveAll()"),
				DIV(CLASS("form-group"), CLASS("col-xs-6"),
					BUTTON(ATTR("type", "submit"), CLASS("btn"), CLASS("btn-primary"), "Save"),
				),
			),
			UL(CLASS("nav"), CLASS("nav-tabs"),
				LI(AHref("#info", ATTR("data-toggle", "tab"), "INFO")),
				LI(A(ng.Href("#section"), ATTR("data-toggle", "tab"), "OVERVIEW"), ng.Click("setContent('OVERVIEW')")),
				dropDownSection("SCENARIO"),
				dropDownSection("NONGOAL"),
				dropDownSection("UNDECIDED"),
				dropDownSection("DEFINITION"),
				dropDownSection("CONTRADICTION"),
				dropDownSection("FEATURE"),
			),

			DIV(ID("ok-message"), CLASS("alert"), CLASS("alert-success"), ng.Show("ok_message"), "{{ok_message}}"),
			DIV(ID("fail-message"), CLASS("alert"), CLASS("alert-danger"), ng.Show("fail_message"), "{{fail_message}}"),

			DIV(CLASS("tab-content"),
				contentInfo(),
				contentSection(),
			),
		),
	),
)

func j(v interface{}) []byte {
	res, err := json.Marshal(v)
	if err != nil {
		panic(err.Error())
	}
	return res
}

func mkSectionJson(section string) []byte {
	ps := spec.Sections[section]
	return j(ps)
}

func specHandler(rw http.ResponseWriter, req *http.Request) {
	section := req.URL.Query().Get("section")
	switch section {
	case "":
		rw.Write(j(spec))
		/*
			case "OVERVIEW":
				rw.Write(j(spec.OVERVIEW))
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
			case "CONTRADICTION":
				rw.Write(mkSectionJson("CONTRADICTION"))
		*/
	}
	/*
		dropDownSection("scenarios", "SCENARIOS"),
					dropDownSection("nongoals", "NONGOALS"),
					dropDownSection("undecided", "UNDECIDED"),
					dropDownSection("definitions", "DEFINITIONS"),
					dropDownSection("features", "FEATURES"),
	*/
}

var spec = &speclib.Spec{Sections: map[string][]speclib.Paragraph{}}

func Save() {
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
	err = spec.LoadJson(fmt.Sprintf("%s", file))
	if err != nil {
		panic(err.Error())
	}

	http.Handle("/", layout)
	http.HandleFunc("/spec.json", specHandler)
	http.HandleFunc("/saveSection", saveSection)
	http.HandleFunc("/saveAll", saveAll)
	http.HandleFunc("/saveInfo", saveInfo)
	http.HandleFunc("/deleteSection", deleteSection)
	http.Handle("/css/", fileserver)
	http.Handle("/js/", fileserver)
	http.Handle("/img/", fileserver)
	http.Handle("/tpl/", fileserver)
	http.Handle("/fonts/", fileserver)
	http.Handle("/codemirror-3.19/", fileserver)
	http.Handle("/angular-1.2.1/", fileserver)
	http.ListenAndServe(":8181", nil)
}
