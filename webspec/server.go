package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-contrib/uuid"
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
		// paragraphs
		A(
			ID(section+"-menu"),
			ATTR("data-toggle", "tab"),
			ng.Href("#section"),
			section,
			ng.Click("setParagraphs('"+section+"')")),
		/*
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
						ng.Hide("isFiltered(s.State)"),
						"{{s.Title}}",
						ng.Click("setContent('"+section+"')")),
				),
			),
		*/

		/*
					<ul class="nav nav-pills">
			  <li class="active"><a href="#">Home</a></li>
			  <li><a href="#">Profile</a></li>
			  <li><a href="#">Messages</a></li>
			</ul>
		*/
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
			ID("info-form"),
			ATTR("name", "infoform"),
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
			/*
				DIV(CLASS("row"),
					BR(),
					DIV(CLASS("form-group"), CLASS("col-xs-6"),
						BUTTON(ATTR("type", "submit"), CLASS("btn"), CLASS("btn-primary"), "Save"),
					),
				),
			*/
		),
	)
}

func contentSection() *goh4.Element {
	return DIV(
		ID("section"), ATTR("role", "form"), CLASS("tab-pane"),
		BR(),
		FORM(
			ID("section-form"),
			ATTR("name", "sectionform"),
			CLASS("form-inline"),
			ng.Submit("saveSection()"),
			DIV(CLASS("row"),
				static("Uuid", 5, HTML("{{paragraph.Uuid}}")),
				input("paragraph.Title", "Title", 7),
			),
			DIV(CLASS("row"),
				BR(),
				//input("paragraph.Responsible", "Responsible", 6),
				//input("paragraph.State", "State", 2),
				selectbox("paragraph.Responsible", "Responsible", 3, "persons"),
				selectbox("paragraph.State", "State", 3, "states"),
				static("LastUpdate", 2, HTML("{{paragraph.LastUpdate}}")),
				//input("paragraph.LastUpdate", "LastUpdate", 2),
				input("paragraph.Deadline", "Deadline", 2),
				input("paragraph.EstimatedHours", "Est. Hours", 2, ATTR("type", "number")),

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
				//input("CommentAuthor", "Author", 4, ng.Blur("saveComment()")),
				selectbox("CommentAuthor", "Author", 4, "persons", ng.Change("setCommentForAuthor()")),
				textarea("CommentText", "Comment", 8, 4, ng.Change("saveComment()")),
			),
			DIV(CLASS("row"),
				BR(),
				/*
					DIV(CLASS("form-group"), CLASS("col-xs-6"),
						BUTTON(ATTR("type", "submit"), CLASS("btn"), CLASS("btn-primary"), "Save"),
					),
				*/
				DIV(CLASS("form-group"), CLASS("col-xs-6"),
					A(
						// goh4.Style{"float", "right"},
						ng.Hide("sectionIndex == -1"),
						ng.Click("deleteSection()"),
						"delete", CLASS("btn"), CLASS("btn-danger")),
				),
			),
		),
	)
}

/*
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
*/
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
	// fmt.Printf("%#v\n", i)
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

// /moveParagraph?section='+$scope.currentSection+"&from="+from+"&to="+to
func moveParagraph(rw http.ResponseWriter, req *http.Request) {
	sec := req.URL.Query().Get("section")

	if sec == "" {
		rw.WriteHeader(400)
		rw.Write([]byte("section missing"))
		return
	}
	from, e := strconv.Atoi(req.URL.Query().Get("from"))
	if e != nil {
		rw.WriteHeader(400)
		rw.Write([]byte(e.Error()))
		return
	}

	to, e := strconv.Atoi(req.URL.Query().Get("to"))
	if e != nil {
		rw.WriteHeader(400)
		rw.Write([]byte(e.Error()))
		return
	}

	err := spec.MoveParagraph(speclib.SectionObj[sec], from, to)
	if err != nil {
		rw.WriteHeader(400)
		rw.Write([]byte(err.Error()))
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

	p.Update()

	if pos == -1 {
		if p.Comments == nil {
			p.Comments = map[string]string{}
		}
		if sec == "OVERVIEW" {
			//spec.Paragraph = &p
			spec.OVERVIEW = &p
		} else {
			p.Uuid = uuid.NewV4().String()
			spec.AddParagraph(speclib.SectionObj[sec], &p)
		}
		Save()
		rw.Write([]byte("ok"))
		return
	}
	//fmt.Printf("sec: %s pos: %d  %#v\n", sec, pos, p)
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
		//CssHref("/codemirror-3.19/theme/monokai.css"),
		//CssHref("/codemirror-3.19/theme/xq-light.css"),
		//CssHref("/codemirror-3.19/theme/eclipse.css"),
		//CssHref("/codemirror-3.19/theme/elegant.css"),
		CssHref("/css/codemirror-webspec.theme.css"),
		JsSrc("/js/jquery-1.10.2.min.js"),
		JsSrc("/codemirror-3.19/lib/codemirror.js"),
		JsSrc("/codemirror-3.19/mode/markdown/markdown.js"),
		JsSrc("/angular-1.2.1/angular.min.js"),
		JsSrc("/angular-1.2.1/angular-animate.min.js"),
		JsSrc("/angular-1.2.1/angular-route.min.js"),
		//JsSrc("/js/ui-codemirror.js"),
		JsSrc("/js/markdown.min.js"),
		JsSrc("/js/bootstrap.js"),
		JsSrc("/js/app.js"),
	),

	BODY(
		ATTR("ng-controller", "Spec"),
		DIV(CLASS("container"),

			DIV(CLASS("row"),
				DIV(CLASS("col-xs-1"),
					H4("SPEC"),
				),
				DIV(CLASS("col-xs-2"),
					DIV(
						CLASS("form-group"),
						BR(),
						SELECT(ng.Model("filteredPerson"),
							CLASS("form-control"),
							// select as label for (key , value) in object
							// countResponsibles
							//ID("filteredPerson"), ATTR("ng-options", "countResponsibles[v] for v in personFilter"),
							//
							ID("filteredPerson"), ATTR("ng-options", "key as key + ' (' + value + ')' for (key , value) in countResponsibles"),
						),
					),
				),
				DIV(CLASS("col-xs-9"),
					CLASS("btn-group"),
					BR(),
					//ATTR("data-toggle", "buttons"),
					//FORM(
					// ['PLANNING', 'APPROVED', 'PARTLY_IMPLEMENTED', 'FULLY_IMPLEMENTED', 'OBSOLET' ];
					LABEL(CLASS("btn"), CLASS("btn-warning"),
						ng.Click("forcestore()"),
						"Save",
					),
					//LABEL(CLASS("btn"), ng.Class("filter.PLANNING ? 'active btn-danger' : 'btn-success'"),
					LABEL(CLASS("btn"), ng.Class("filterClass('PLANNING')"),
						InputCheckbox("filter-PLANNING", ng.Model("filter.PLANNING"), goh4.Style{"display", "none"},
							ng.Click("setFilter()"),
						),
						"{{count.PLANNING}} PLANNING",
					),

					//LABEL(CLASS("btn"), ng.Class("filter.APPROVED ? 'active btn-danger' : 'btn-success'"),
					LABEL(CLASS("btn"), ng.Class("filterClass('APPROVED')"),
						InputCheckbox("filter-APPROVED", ng.Model("filter.APPROVED"), goh4.Style{"display", "none"},
							ng.Click("setFilter()"),
						),
						"{{count.APPROVED}} APPROVED",
					),

					//LABEL(CLASS("btn"), ng.Class("filter.PARTLY_IMPLEMENTED ? 'active btn-danger' : 'btn-success'"),
					LABEL(CLASS("btn"), ng.Class("filterClass('PARTLY_IMPLEMENTED')"),
						InputCheckbox("filter-PARTLY_IMPLEMENTED", ng.Model("filter.PARTLY_IMPLEMENTED"), goh4.Style{"display", "none"},
							ng.Click("setFilter()"),
						),
						"{{count.PARTLY_IMPLEMENTED}} PARTLY_IMPLEMENTED",
					),

					//LABEL(CLASS("btn"), ng.Class("filter.FULLY_IMPLEMENTED ? 'active btn-danger' : 'btn-success'"),
					LABEL(CLASS("btn"), ng.Class("filterClass('FULLY_IMPLEMENTED')"),
						InputCheckbox("filter-FULLY_IMPLEMENTED", ng.Model("filter.FULLY_IMPLEMENTED"), goh4.Style{"display", "none"},
							ng.Click("setFilter()"),
						),
						"{{count.FULLY_IMPLEMENTED}} FULLY_IMPLEMENTED",
					),

					// LABEL(CLASS("btn"), ng.Class("filter.OBSOLET ? 'active btn-danger' : 'btn-success'"),
					LABEL(CLASS("btn"), ng.Class("filterClass('OBSOLET')"),
						InputCheckbox("filter-OBSOLET", ng.Model("filter.OBSOLET"), goh4.Style{"display", "none"},
							ng.Click("setFilter()"),
						),
						"{{count.OBSOLET}} OBSOLET",
					),
				),
			),
			/*
							<div class="btn-group" data-toggle="buttons">
				  <label class="btn btn-primary">
				    <input type="checkbox"> Option 1
				  </label>
				  <label class="btn btn-primary">
				    <input type="checkbox"> Option 2
				  </label>
				  <label class="btn btn-primary">
				    <input type="checkbox"> Option 3
				  </label>
				</div>
			*/

			//BUTTON(CLASS("btn"), CLASS("btn-info"), ng.Click("setFilter(s)"), ng.Repeat("s", "states"), "{{s}}"),
			// ng-animate="'animate'"
			DIV(ID("ok-message"), CLASS("alert"), CLASS("alert-success"), ng.Show("ok_message"), "{{ok_message}}"),
			DIV(ID("fail-message"), CLASS("alert"), CLASS("alert-danger"), ng.Show("fail_message"), "{{fail_message}}"),
			/*
				FORM(
					CLASS("form-inline"),
					ng.Submit("saveAll()"),
					DIV(CLASS("form-group"), CLASS("col-xs-6"),
						BUTTON(ATTR("type", "submit"), CLASS("btn"), CLASS("btn-primary"), "Save"),
					),
				),
			*/
			UL(CLASS("nav"), CLASS("nav-tabs"),
				LI(A(ng.Href("#info"), ng.Click("unsetCurrentSection()"), ATTR("data-toggle", "tab"), "INFO")),
				LI(A(ng.Href("#section"), ATTR("data-toggle", "tab"), "OVERVIEW"), ng.Click("setParagraphs('OVERVIEW')")),
				dropDownSection("SCENARIO"),
				dropDownSection("NONGOAL"),
				dropDownSection("UNDECIDED"),
				dropDownSection("DEFINITION"),
				dropDownSection("CONTRADICTION"),
				dropDownSection("FEATURE"),
				LI(A(ng.Href("#validation"), ATTR("data-toggle", "tab"), "Validate"), ng.Click("validate()")),
			),

			DIV(CLASS("row"),
				DIV(CLASS("col-xs-3"),
					BR(),
					UL(CLASS("nav"), CLASS("nav-pills"), CLASS("nav-stacked"),
						/*
							handleDropped($index)
						*/
						ATTR("droppable", "droppable", "bin", "paragraphBin"),
						goh4.Style{"padding-bottom", "30px"},
						LI(
							A(
								ATTR("data-toggle", "tab"),
								ng.Href("#section"),
								"new",
								ng.Show("(currentSection && currentSection != 'OVERVIEW')"),
								ng.Href("#section"),
								ng.Click("newPara()")),
						),
						LI(
							ng.Repeat("p", "paragraphs"),
							ID("{{$index}}"),
							//draggable item="item"
							ATTR("draggable", "draggable", "item", "$index", "drop", "handleDropped"),
							//ng.Class("$index == sectionIndex ? 'ng-activated' : ''"),
							A(
								ATTR("data-toggle", "tab"),
								ng.Href("#section"),
								ng.Hide("isFiltered(p.State, p.Responsible)"),
								"{{p.Title}}",
								ng.Click("setPara()")),
						),
					),
				),
				DIV(CLASS("col-xs-9"),
					DIV(CLASS("tab-content"),
						contentInfo(),
						contentSection(),
						DIV(ID("validation"), CLASS("tab-pane"),
							H4("Validation Response"),
							PRE("{{validationMessage}}"),
						),
					),
				),
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
	spec.Update()
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
	//	http.HandleFunc("/saveAll", saveAll)
	http.HandleFunc("/moveParagraph", moveParagraph)
	http.HandleFunc("/saveInfo", saveInfo)
	http.HandleFunc("/deleteSection", deleteSection)
	http.HandleFunc("/validate", validate)
	http.Handle("/css/", fileserver)
	http.Handle("/js/", fileserver)
	http.Handle("/img/", fileserver)
	http.Handle("/tpl/", fileserver)
	http.Handle("/fonts/", fileserver)
	http.Handle("/codemirror-3.19/", fileserver)
	http.Handle("/angular-1.2.1/", fileserver)
	http.ListenAndServe(":8181", nil)
}
