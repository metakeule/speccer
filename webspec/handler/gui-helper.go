package handler

import (
	"fmt"
	"github.com/metakeule/goh4"
	. "github.com/metakeule/goh4/tag"
	. "github.com/metakeule/goh4/tag/short"
	"github.com/metakeule/ng"
	"strings"
)

func (h *handler) personFilter() *goh4.Element {
	return DIV(CLASS("col-xs-4"),
		ng.Show("currentSection"),
		DIV(
			CLASS("row"),
			DIV(CLASS("label"), CLASS("label-default"), h.translate("RESPONSIBLE"), CLASS("col-xs-4")),

			DIV(
				CLASS("form-group"),
				CLASS("col-xs-8"),
				SELECT(ng.Model("filteredPerson"),
					CLASS("form-control"),
					CLASS("input-sm"),
					ID("filteredPerson"),
					ng.Change("setCounts()"),
					ATTR("ng-options", "key as translate(key) + ' (' + value + ')' for (key , value) in countResponsibles"),
				),
			),
		),
	)
}

func (h *handler) stateFilter(state string) *goh4.Element {
	return LABEL(CLASS("btn"), CLASS("btn-sm"), ng.Class("filterClass('"+state+"')"),
		ng.Show("currentSection"),
		InputCheckbox("filter-"+state, ng.Model("filter."+state+""), goh4.Style{"display", "none"},
			ng.Click("setFilter()"),
		),
		SPAN(CLASS("badge"), "{{count."+state+"}}"),
		" "+h.translate(state),
	)
}

func (h *handler) dropDownSection(section string) *goh4.Element {
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
			h.translate(section),
			BR(),
			SPAN(CLASS("badge"), CLASS("badge-warning"), "{{sectionNumbers['"+section+"']}}"),
			//HTML("&nbsp;"),
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
		ID(fieldName), ATTR("ng-options", "translate(v) for v in "+arraySrc))

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

func mapList(model, key, value string, labelmodel, labelkey, labelvalue string) *goh4.Element {
	var keyD, valD = strings.ToLower(key), strings.ToLower(value)
	return DIV(
		BR(),
		CLASS("row"),
		DIV(CLASS("form-group"), CLASS("col-xs-6"),
			LABEL(labelmodel),
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
		input(model+key, labelkey, 2, ng.Blur("save"+model+"()")),
		input(model+value, labelvalue, 4, ng.Change("save"+model+"()")),
	)
}

func arrayList(model, value string, labelproperty, label string) *goh4.Element {
	var valD = strings.ToLower(value)
	return DIV(
		BR(),
		CLASS("row"),
		DIV(CLASS("form-group"), CLASS("col-xs-6"),
			LABEL(labelproperty),
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
		input(model+value, label, 4, ng.Blur("add"+model+"()")),
	)
}

func (h *handler) contentInfo() *goh4.Element {
	return DIV(
		ID("info"), ATTR("role", "form"), CLASS("tab-pane"),
		BR(),
		FORM(
			ID("info-form"),
			ATTR("name", "infoform"),
			CLASS("form-inline"),
			ng.Submit("saveInfo()"),
			DIV(CLASS("row"),
				input("INFO.Company", h.translate("COMPANY"), 6),
				input("INFO.Project", h.translate("PROJECT"), 6),
			),
			DIV(CLASS("row"),
				input("INFO.URL", h.translate("URL"), 6),
				input("INFO.Parent", h.translate("PARENT"), 6),
			),
			DIV(CLASS("row"),
				input("INFO.Language", h.translate("LANGUAGE"), 4),
				input("INFO.DateFormat", h.translate("DATEFORMAT"), 4),
				selectbox("INFO.Approved", h.translate("APPROVED"), 4, "[true,false]"),
			),
			arrayList("RequestedBy", "Name", h.translate("REQUESTED_BY"), h.translate("NAME")),
			mapList("Translations", "Language", "URL", h.translate("TRANSLATIONS"), h.translate("LANGUAGE"), h.translate("URL")),
			mapList("Related", "Name", "URL", h.translate("RELATED"), h.translate("NAME"), h.translate("URL")),
			mapList("SupersededBy", "Name", "URL", h.translate("SUPERSEDED_BY"), h.translate("NAME"), h.translate("URL")),
			mapList("Resources", "Name", "URL", h.translate("RESOURCES"), h.translate("NAME"), h.translate("URL")),
			mapList("Persons", "Short", "Full", h.translate("PERSONS"), h.translate("SHORTN_AME"), h.translate("NAME")),
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

func (h *handler) contentSection() *goh4.Element {
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
						DIV(CLASS("panel-heading"),
							DIV(
								CLASS("form-group"),
								CLASS("col-xs-12"),
								LabelFor("paragraph.Title", h.translate("TIT_TLE_OF_UUI_D")+" {{paragraph.UUID}}", ng.If("sectionIndex != -1 || currentSection == 'OVERVIEW'")),
								LabelFor("paragraph.Title", h.translate("NE_W_TI_T_TLE"), ng.If("sectionIndex == -1 && currentSection != 'OVERVIEW'")),
								InputText("paragraph.Title", ng.Model("paragraph.Title"),
									CLASS("form-control"),
									ID("paragraph.Title")),
							),
						),
						DIV(CLASS("panel-body"),
							DIV(CLASS("row"),
								selectbox("paragraph.Responsible", h.translate("RESPONSIBLE"), 6, "persons"),
								selectbox("paragraph.State", h.translate("STATE"), 6, "states"),
							),
							BR(),
							DIV(CLASS("row"),
								input("paragraph.Deadline", h.translate("DEADLINE"), 6),
								input("paragraph.EstimatedHours", h.translate("ESTIMATEDHOURS"), 3, ATTR("type", "number")),
								static(h.translate("LASTUPDATE"), 3, HTML("{{paragraph.LastUpdate}}")),
							),
						),
					),

					DIV(CLASS("panel"), CLASS("panel-primary"),
						DIV(CLASS("panel-heading"), h.translate("TEX_T_(MARKDOWN)")),
						DIV(ID("Text"), CLASS("panel-body")),
					),
					DIV(CLASS("panel"), CLASS("panel-default"),
						DIV(CLASS("panel-heading"), h.translate("HTM_L_PREVIEW")),
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
								h.translate("SAVE"),
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
								h.translate("DELETE"),
							),
						),
					),
					BR(),
					DIV(CLASS("panel"), CLASS("panel-info"),
						DIV(CLASS("panel-heading"), h.translate("COMMEN_TS")),
						DIV(CLASS("panel-body"),
							DIV(
								CLASS("row"),
								ng.RepeatKeyVal("author", "comment", "paragraph.Comments"),
								DIV(CLASS("form-group"), CLASS("col-xs-12"),
									DIV(CLASS("btn-group"),
										goh4.Style{"margin-right", "15px"},
										BUTTON(
											ATTR("type", "button"),
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

								DIV(
									CLASS("form-group"),
									CLASS("col-xs-12"),
									LabelFor("CommentAuthor", h.translate("AUTHOR")),
									SELECT(ng.Model("CommentAuthor"),
										CLASS("form-control"),
										ng.Change("setCommentForAuthor()"),
										ID("CommentAuthor"), ATTR("ng-options", "v for v in persons")),
								),
							),
							BR(),
							DIV(
								CLASS("row"),
								textarea("CommentText", h.translate("COMMENT"), 12, 20, ng.Change("saveComment()")),
							),
						),
					),
					BR(),
				),
			),
		),
	)
}
