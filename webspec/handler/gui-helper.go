package handler

import (
	"fmt"
	"github.com/metakeule/goh4"
	. "github.com/metakeule/goh4/tag"
	. "github.com/metakeule/goh4/tag/short"
	"github.com/metakeule/ng"
	"strings"
)

func dropDownSection(section string) *goh4.Element {
	return LI(
		// paragraphs
		A(
			ID(section+"-menu"),
			ATTR("data-toggle", "tab"),
			ATTR("droppable", "droppable"),
			ATTR("section", section),
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
				static("UUID", 5, HTML("{{paragraph.UUID}}")),
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
