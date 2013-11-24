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
		A(
			ID(section+"-menu"),
			// CLASS("btn"),
			// CLASS("btn-default"),
			//CLASS("btn-success"),
			ATTR("data-toggle", "tab"),
			ATTR("droppable", "droppable"),
			ATTR("section", section),
			ng.Href("#section"),
			section,
			BR(),
			SPAN(CLASS("badge"), CLASS("badge-warning"), "{{sectionNumbers['"+section+"']}}"),
			HTML("&nbsp;"),
			ng.Click("setParagraphs('"+section+"')")),
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
					CLASS("btn-sm"),
					ng.Click("set"+model+"("+keyD+", "+valD+")"), "{{"+keyD+"}}", CLASS("btn"), CLASS("btn-default"),
				),
				BUTTON(
					ATTR("type", "button"),
					goh4.Style{"float", "none"},
					ng.Click("remove"+model+"("+keyD+")"),
					//CLASS("close"),
					CLASS("btn"),
					CLASS("btn-danger"),
					CLASS("btn-sm"),
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
					CLASS("btn-sm"),
				),
				BUTTON(
					ATTR("type", "button"),
					goh4.Style{"float", "none"},
					ng.Click("remove"+model+"("+valD+")"),
					//CLASS("close"),
					CLASS("btn"),
					CLASS("btn-danger"),
					CLASS("btn-sm"),
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
		),
	)
}

func infoRow(name, value string) *goh4.Element {
	return DIV(CLASS("row"),
		DIV(CLASS("col-xs-2"), H3(name)),
		DIV(CLASS("col-xs-10"), value),
	)
}

func infoArray(name, collection string) *goh4.Element {
	return DIV(CLASS("row"),
		DIV(CLASS("col-xs-2"), H3(name)),
		DIV(CLASS("col-xs-10"),
			SPAN(
				ng.Repeat("v", collection),
				CLASS("btn"), CLASS("btn-default"),
				goh4.Style{"margin-right", "12px"},
				"{{v}}"),
		),
	)
}

func infoMap(name, collection string) *goh4.Element {
	return DIV(CLASS("row"),
		DIV(CLASS("col-xs-2"), H3(name)),
		DIV(CLASS("col-xs-10"),
			DL(CLASS("dl-horizontal"),
				ng.RepeatKeyVal("k", "v", collection),
				DT("{{k}}"),
				DD("{{v}}"),
			),
		),
	)
}

func contentInfoView() *goh4.Element {
	return DIV(
		ID("info"), CLASS("tab-pane"),
		BR(),
		infoRow("Company:", "{{INFO.Company}}"),
		infoRow("Project:", "{{INFO.Project}}"),

		infoRow("URL:", "{{INFO.URL}}"),
		infoRow("Parent:", "{{INFO.Parent}}"),

		infoRow("Language:", "{{INFO.Language}}"),
		infoRow("DateFormat:", "{{INFO.DateFormat}}"),
		infoRow("Approved:", "{{INFO.Approved}}"),

		infoArray("RequestedBy:", "INFO.RequestedBy"),
		infoMap("Translations", "INFO.Translations"),

		infoMap("Related", "INFO.Related"),
		infoMap("SupersededBy", "INFO.SupersededBy"),
		infoMap("Resources", "INFO.Resources"),
		infoMap("Persons", "INFO.Persons"),
	)
}

func contentSection() *goh4.Element {
	return DIV(
		ID("section"), ATTR("role", "form"), CLASS("tab-pane"),
		FORM(
			ID("section-form"),
			ATTR("name", "sectionform"),
			CLASS("form-inline"),
			ng.Submit("saveSection()"),
			DIV(CLASS("row"),
				DIV(CLASS("col-xs-8"),
					DIV(CLASS("panel"), CLASS("panel-default"),
						//DIV(CLASS("panel-heading"), STRONG("UUID: {{paragraph.UUID}}"), ng.If("sectionIndex != -1 || currentSection == 'OVERVIEW'")),
						DIV(CLASS("panel-heading"),
							//
							DIV(
								CLASS("form-group"),
								CLASS("col-xs-12"),
								LabelFor("paragraph.Title", "Title of UUID {{paragraph.UUID}}", ng.If("sectionIndex != -1 || currentSection == 'OVERVIEW'")),
								LabelFor("paragraph.Title", "new Title", ng.If("sectionIndex == -1 && currentSection != 'OVERVIEW'")),
								InputText("paragraph.Title", ng.Model("paragraph.Title"),
									CLASS("form-control"),
									ID("paragraph.Title")),
							),
						),
						//DIV(CLASS("panel-heading"), STRONG("New"), ng.If("sectionIndex == -1 && currentSection != 'OVERVIEW'")),
						DIV(CLASS("panel-body"),
							DIV(CLASS("row"),
								selectbox("paragraph.Responsible", "Responsible", 6, "persons"),
								selectbox("paragraph.State", "State", 6, "states"),
							),
							BR(),
							DIV(CLASS("row"),
								// input("paragraph.Title", "Title", 8),
								//static("UUID", 8, HTML("{{paragraph.UUID}}")),
								input("paragraph.Deadline", "Deadline", 6),
								input("paragraph.EstimatedHours", "Est. Hours", 3, ATTR("type", "number")),
								static("LastUpdate", 3, HTML("{{paragraph.LastUpdate}}")),
							),
						),
					),

					DIV(CLASS("panel"), CLASS("panel-primary"),
						DIV(CLASS("panel-heading"), "Text (Markdown)"),
						DIV(ID("Text"), CLASS("panel-body")),
					),
					DIV(CLASS("panel"), CLASS("panel-default"),
						DIV(CLASS("panel-heading"), "HTML-Preview"),
						DIV(ID("preview"), CLASS("panel-body")),
					),
				),

				DIV(CLASS("col-xs-4"),
					DIV(CLASS("row"),
						DIV(
							CLASS("form-group"),
							CLASS("col-xs-6"),
							LABEL(
								ng.Click("forcestore()"),
								CLASS("btn"),
								CLASS("btn-warning"),
								CLASS("form-control"),
								"Save",
							),
						),
						DIV(
							CLASS("form-group"),
							CLASS("col-xs-6"),
							LABEL(
								ng.Click("deleteSection()"),
								ng.Hide("sectionIndex == -1"),
								CLASS("btn"),
								CLASS("btn-danger"),
								CLASS("form-control"),
								"delete",
							),
						),
					),
					BR(),
					DIV(CLASS("panel"), CLASS("panel-info"),
						DIV(CLASS("panel-heading"), "Comments"),
						DIV(CLASS("panel-body"),
							DIV(
								CLASS("row"),
								ng.RepeatKeyVal("author", "comment", "paragraph.Comments"),
								DIV(CLASS("form-group"), CLASS("col-xs-12"),
									DIV(CLASS("btn-group"),
										goh4.Style{"margin-right", "15px"},
										BUTTON(
											ATTR("type", "button"),
											//<span class="glyphicon glyphicon-star"></span>
											SPAN(CLASS("glyphicon"), CLASS("glyphicon-pencil")),
											HTML("&nbsp;"),
											ng.Click("setComment(author, comment)"), "{{author}}",
											CLASS("btn-sm"),
											CLASS("btn"), CLASS("btn-default"),
										),
										BUTTON(
											ATTR("type", "button"),
											goh4.Style{"float", "none"},
											ng.Click("removeComment(author)"),
											//CLASS("close"),
											CLASS("btn"),
											CLASS("btn-danger"),
											CLASS("btn-sm"),
											ATTR("type", "button", "aria-hidden", "true"), HTML("&times;")),
									),
								),
							),
							BR(),

							DIV(
								CLASS("row"),
								selectbox("CommentAuthor", "Author", 12, "persons", ng.Change("setCommentForAuthor()")),
							),
							BR(),
							DIV(
								CLASS("row"),
								textarea("CommentText", "Comment", 12, 20, ng.Change("saveComment()")),
							),
						),
					),
					BR(),
				),
			),
		),
	)
}
