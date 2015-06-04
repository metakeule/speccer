package handler

import (
	"fmt"
	"strings"

	. "gopkg.in/go-on/lib.v2/html"
	. "gopkg.in/go-on/lib.v2/html/internal/element"
	"gopkg.in/go-on/lib.v2/internal/ng"
	. "gopkg.in/go-on/lib.v2/types"
)

func (h *handler) personFilter() *Element {
	return DIV(Class("col-xs-4"),
		ng.Show("currentSection"),
		DIV(
			Class("row"),
			DIV(Class("label"), Class("label-default"), h.translate("RESPONSIBLE"), Class("col-xs-4")),

			DIV(
				Class("form-group"),
				Class("col-xs-8"),
				SELECT(ng.Model("filteredPerson"),
					Class("form-control"),
					Class("input-sm"),
					Id("filteredPerson"),
					ng.Change("setCounts()"),
					Attrs_("ng-options", "key as translate(key) + ' (' + value + ')' for (key , value) in countResponsibles"),
				),
			),
		),
	)
}

func (h *handler) stateFilter(state string) *Element {
	return LABEL(Class("btn"), Class("btn-sm"), ng.Class("filterClass('"+state+"')"),
		ng.Show("currentSection"),
		InputCheckbox("filter-"+state, ng.Model("filter."+state+""), Style{"display", "none"},
			ng.Click("setFilter()"),
		),
		SPAN(Class("badge"), "{{count."+state+"}}"),
		" "+h.translate(state),
	)
}

func (h *handler) dropDownSection(section string) *Element {
	return LI(
		A(
			Id(section+"-menu"),
			// Class("btn"),
			// Class("btn-default"),
			//Class("btn-success"),
			Attrs_("data-toggle", "tab"),
			Attrs_("droppable", "droppable"),
			Attrs_("section", section),
			ng.Href("#section"),
			h.translate(section),
			BR(),
			SPAN(Class("badge"), Class("badge-warning"), "{{sectionNumbers['"+section+"']}}"),
			//HTMLString("&nbsp;"),
			ng.Click("setParagraphs('"+section+"')")),
	)
}

func input(fieldName string, label string, width int, i ...interface{}) *Element {
	i = append(i,
		ng.Model(fieldName),
		Class("form-control"),
		Id(fieldName))
	return DIV(
		Class("form-group"),
		Class(fmt.Sprintf("col-xs-%d", width)),
		LabelFor(fieldName, label),
		InputText(fieldName, i...),
	)
}

func static(label string, width int, i ...interface{}) *Element {
	i = append(i,
		Class("form-control-static"))
	return DIV(
		Class("form-group"),
		Class(fmt.Sprintf("col-xs-%d", width)),
		LABEL(label),
		P(i...),
	)
}

func selectbox(fieldName string, label string, width int, arraySrc string, i ...interface{}) *Element {
	i = append(i,
		ng.Model(fieldName),
		Class("form-control"),
		Id(fieldName), Attrs_("ng-options", "translate(v) for v in "+arraySrc))

	return DIV(
		Class("form-group"),
		Class(fmt.Sprintf("col-xs-%d", width)),
		LabelFor(fieldName, label),
		SELECT(i...),
	)
}

func textarea(fieldName string, label string, width int, rows int, i ...interface{}) *Element {
	i = append(i,
		ng.Model(fieldName),
		Class("form-control"),
		Id(fieldName),
		Attrs_("rows", fmt.Sprintf("%d", rows), "value", "{{fieldName}}", "name", fieldName))
	return DIV(Class("form-group"),
		Class(fmt.Sprintf("col-xs-%d", width)),
		LabelFor(fieldName, label),
		TEXTAREA(i...),
	)
}

func mapList(model, key, value string, labelmodel, labelkey, labelvalue string) *Element {
	var keyD, valD = strings.ToLower(key), strings.ToLower(value)
	return DIV(
		BR(),
		Class("row"),
		DIV(Class("form-group"), Class("col-xs-6"),
			LABEL(labelmodel),
			BR(),
			DIV(Class("btn-group"),
				Style{"margin-right", "15px"},
				ng.RepeatKeyVal(keyD, valD, "INFO."+model),
				BUTTON(
					Attrs_("type", "button"),
					SPAN(Class("glyphicon"), Class("glyphicon-pencil")),
					HTMLString("&nbsp;"),
					Class("btn-sm"),
					ng.Click("set"+model+"("+keyD+", "+valD+")"), "{{"+keyD+"}}", Class("btn"), Class("btn-default"),
				),
				BUTTON(
					Attrs_("type", "button"),
					Style{"float", "none"},
					ng.Click("remove"+model+"("+keyD+")"),
					//Class("close"),
					Class("btn"),
					Class("btn-danger"),
					Class("btn-sm"),
					Attrs_("type", "button", "aria-hidden", "true"), HTMLString("&times;")),
			),
		),
		input(model+key, labelkey, 2, ng.Blur("save"+model+"()")),
		input(model+value, labelvalue, 4, ng.Change("save"+model+"()")),
	)
}

func arrayList(model, value string, labelproperty, label string) *Element {
	var valD = strings.ToLower(value)
	return DIV(
		BR(),
		Class("row"),
		DIV(Class("form-group"), Class("col-xs-6"),
			LABEL(labelproperty),
			BR(),
			DIV(Class("btn-group"),
				Style{"margin-right", "15px"},
				ng.Repeat(valD, "INFO."+model),
				BUTTON(
					Attrs_("type", "button"),
					SPAN(Class("glyphicon"), Class("glyphicon-pencil")),
					HTMLString("&nbsp;"),
					"{{"+valD+"}}", Class("btn"), Class("btn-default"),
					Class("btn-sm"),
				),
				BUTTON(
					Attrs_("type", "button"),
					Style{"float", "none"},
					ng.Click("remove"+model+"("+valD+")"),
					//Class("close"),
					Class("btn"),
					Class("btn-danger"),
					Class("btn-sm"),
					Attrs_("type", "button", "aria-hidden", "true"), HTMLString("&times;")),
			),
		),
		input(model+value, label, 4, ng.Blur("add"+model+"()")),
	)
}

func (h *handler) contentInfo() *Element {
	return DIV(
		Id("info"), Attrs_("role", "form"), Class("tab-pane"),
		BR(),
		FORM(
			Id("info-form"),
			Attrs_("name", "infoform"),
			Class("form-inline"),
			ng.Submit("saveInfo()"),
			DIV(Class("row"),
				input("INFO.Company", h.translate("COMPANY"), 6),
				input("INFO.Project", h.translate("PROJECT"), 6),
			),
			DIV(Class("row"),
				input("INFO.URL", h.translate("URL"), 6),
				input("INFO.Parent", h.translate("PARENT"), 6),
			),
			DIV(Class("row"),
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

func infoRow(name, value string) *Element {
	return DIV(Class("row"),
		DIV(Class("col-xs-2"), H3(name)),
		DIV(Class("col-xs-10"), value),
	)
}

func infoArray(name, collection string) *Element {
	return DIV(Class("row"),
		DIV(Class("col-xs-2"), H3(name)),
		DIV(Class("col-xs-10"),
			SPAN(
				ng.Repeat("v", collection),
				Class("btn"), Class("btn-default"),
				Style{"margin-right", "12px"},
				"{{v}}"),
		),
	)
}

func infoMap(name, collection string) *Element {
	return DIV(Class("row"),
		DIV(Class("col-xs-2"), H3(name)),
		DIV(Class("col-xs-10"),
			DL(Class("dl-horizontal"),
				ng.RepeatKeyVal("k", "v", collection),
				DT("{{k}}"),
				DD("{{v}}"),
			),
		),
	)
}

func (h *handler) contentSection() *Element {
	return DIV(
		Id("section"), Attrs_("role", "form"), Class("tab-pane"),
		FORM(
			Id("section-form"),
			Attrs_("name", "sectionform"),
			Class("form-inline"),
			ng.Submit("saveSection()"),
			DIV(Class("row"),
				DIV(Class("col-xs-8"),
					DIV(Class("panel"), Class("panel-default"),
						DIV(Class("panel-heading"),
							DIV(
								Class("form-group"),
								Class("col-xs-12"),
								LabelFor("paragraph.Title", h.translate("TIT_TLE_OF_UUI_D")+" {{paragraph.UUID}}", ng.If("sectionIndex != -1 || currentSection == 'OVERVIEW'")),
								LabelFor("paragraph.Title", h.translate("NE_W_TI_T_TLE"), ng.If("sectionIndex == -1 && currentSection != 'OVERVIEW'")),
								InputText("paragraph.Title", ng.Model("paragraph.Title"),
									Class("form-control"),
									Id("paragraph.Title")),
							),
						),
						DIV(Class("panel-body"),
							DIV(Class("row"),
								selectbox("paragraph.Responsible", h.translate("RESPONSIBLE"), 6, "persons"),
								selectbox("paragraph.State", h.translate("STATE"), 6, "states"),
							),
							BR(),
							DIV(Class("row"),
								input("paragraph.Deadline", h.translate("DEADLINE"), 6),
								input("paragraph.EstimatedHours", h.translate("ESTIMATEDHOURS"), 3, Attrs_("type", "number")),
								static(h.translate("LASTUPDATE"), 3, HTMLString("{{paragraph.LastUpdate}}")),
							),
						),
					),

					DIV(Class("panel"), Class("panel-primary"),
						DIV(Class("panel-heading"), h.translate("TEX_T_(MARKDOWN)")),
						DIV(Id("Text"), Class("panel-body")),
					),
					DIV(Class("panel"), Class("panel-default"),
						DIV(Class("panel-heading"), h.translate("HTM_L_PREVIEW")),
						DIV(Id("preview"), Class("panel-body")),
					),
				),

				DIV(Class("col-xs-4"),
					DIV(Class("row"),
						DIV(
							Class("form-group"),
							Class("col-xs-6"),
							LABEL(
								ng.Click("forcestore()"),
								Class("btn"),
								Class("btn-warning"),
								Class("form-control"),
								h.translate("SAVE"),
							),
						),
						DIV(
							Class("form-group"),
							Class("col-xs-6"),
							LABEL(
								ng.Click("deleteSection()"),
								ng.Hide("sectionIndex == -1"),
								Class("btn"),
								Class("btn-danger"),
								Class("form-control"),
								h.translate("DELETE"),
							),
						),
					),
					BR(),
					DIV(Class("panel"), Class("panel-info"),
						DIV(Class("panel-heading"), h.translate("COMMEN_TS")),
						DIV(Class("panel-body"),
							DIV(
								Class("row"),
								ng.RepeatKeyVal("author", "comment", "paragraph.Comments"),
								DIV(Class("form-group"), Class("col-xs-12"),
									DIV(Class("btn-group"),
										Style{"margin-right", "15px"},
										BUTTON(
											Attrs_("type", "button"),
											SPAN(Class("glyphicon"), Class("glyphicon-pencil")),
											HTMLString("&nbsp;"),
											ng.Click("setComment(author, comment)"), "{{author}}",
											Class("btn-sm"),
											Class("btn"), Class("btn-default"),
										),
										BUTTON(
											Attrs_("type", "button"),
											Style{"float", "none"},
											ng.Click("removeComment(author)"),
											//Class("close"),
											Class("btn"),
											Class("btn-danger"),
											Class("btn-sm"),
											Attrs_("type", "button", "aria-hidden", "true"), HTMLString("&times;")),
									),
								),
							),
							BR(),
							DIV(
								Class("row"),

								DIV(
									Class("form-group"),
									Class("col-xs-12"),
									LabelFor("CommentAuthor", h.translate("AUTHOR")),
									SELECT(ng.Model("CommentAuthor"),
										Class("form-control"),
										ng.Change("setCommentForAuthor()"),
										Id("CommentAuthor"), Attrs_("ng-options", "v for v in persons")),
								),
							),
							BR(),
							DIV(
								Class("row"),
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
