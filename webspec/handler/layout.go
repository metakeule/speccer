package handler

import (
	"github.com/metakeule/goh4"
	. "github.com/metakeule/goh4/tag"
	. "github.com/metakeule/goh4/tag/short"
	"github.com/metakeule/ng"
)

func (h *handler) layout(content ...interface{}) *DocType {
	container := DIV(CLASS("container"))
	container.Add(content...)
	return HTML5(
		ATTR("ng-app", "app"),
		HEAD(
			CharsetUtf8(),
			META(ATTR("name", "viewport", "content", "width=device-width, initial-scale=1.0")),
			HTML(`
			<!--[if lt IE 9]>
      			<script src="public/js/html5shiv.js"></script>
      			<script src="public/js/1.3.0/respond.min.js"></script>
    		<![endif]-->
    	`),
			CssHref("public/css/bootstrap.css"),
			CssHref("public/css/codemirror.css"),
			CssHref("public/css/codemirror-webspec.theme.css"),
			CssHref("public/css/app.css"),
			JsSrc("public/js/jquery-1.10.2.min.js"),
			JsSrc("public/js/codemirror.js"),
			JsSrc("public/js/codemirror-mode/markdown/markdown.js"),
			JsSrc("public/js/angular.min.js"),
			JsSrc("public/js/angular-animate.min.js"),
			JsSrc("public/js/angular-route.min.js"),
			JsSrc("public/js/markdown.min.js"),
			JsSrc("public/js/bootstrap.js"),
			JsSrc("public/js/app.js"),
			TITLE(h.translate("SPE_C")+" '"+h.Spec.OVERVIEW.Title+"' ["+h.Spec.OVERVIEW.UUID+"]"),
		),
		BODY(ATTR("ng-controller", "Spec"), container),
	)
}

func (h *handler) adminlayout() *DocType {
	return h.layout(
		DIV(CLASS("row"),
			H1(h.translate("SPE_C")+" '{{Sections.OVERVIEW.Title}}'"),
			P(CLASS("well"), CLASS("well-sm"), STRONG(h.translate("LASTUPDATE")+" "),
				"{{Sections.OVERVIEW.LastUpdate}}",
				HTML("&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"),
				STRONG(h.translate("UUID")+" "), "{{Sections.OVERVIEW.UUID}} "),
		),

		DIV(ID("ok-message"), CLASS("alert"), CLASS("alert-success"), ng.Show("ok_message"), "{{ok_message}}"),
		DIV(ID("fail-message"), CLASS("alert"), CLASS("alert-danger"),
			CLASS("alert-dismissable"),
			ng.Show("fail_message"), "{{fail_message}}",
			BUTTON(ATTR("type", "button", "data-dismiss", "alert", "aria-hidden", "true"), HTML("&times;"), CLASS("close")),
		),

		DIV(CLASS("row"),
			CLASS("navbar-default"),
			UL(CLASS("nav"),
				CLASS("navbar-nav"),
				//CLASS("nav-tabs"),
				//CLASS("nav-pills"),
				LI(A(ng.Href("#info"), ng.Click("unsetCurrentSection()"), ATTR("data-toggle", "tab"), h.translate("INFO"))),
				LI(A(ng.Href("#section"), ATTR("data-toggle", "tab"), h.translate("OVERVIEW"),
					BR(),
					SPAN(CLASS("badge"), CLASS("badge-warning"), "{{sectionNumbers['OVERVIEW']}}"),
				), ng.Click("setParagraphs('OVERVIEW')")),
				h.dropDownSection("SCENARIO"),
				h.dropDownSection("NONGOAL"),
				h.dropDownSection("UNDECIDED"),
				h.dropDownSection("DEFINITION"),
				h.dropDownSection("CONTRADICTION"),
				h.dropDownSection("FEATURE"),
				LI(
					DIV(CLASS("navbar-form "),
						A(
							ng.Href("#validation"),
							CLASS("btn"), CLASS("btn-default"),
							ATTR("data-toggle", "tab"), h.translate("VALIDATE")),
						ng.Click("validate()"),
					),
				),
				LI(
					DIV(CLASS("navbar-form "),
						AHref("spec.html", ATTR("target", "_blank"),
							// A(
							// ng.Href("#validation"),
							CLASS("btn"), CLASS("btn-default"),
							h.translate("SHOW"),
						),
						// ATTR("data-toggle", "tab"), "Validate"),
						// ng.Click("validate()"),
					),
				),
			),
		),

		BR(),
		DIV(CLASS("panel"), CLASS("panel-default"), ID("filter-bar"),
			ng.Show("currentSection"),
			DIV(CLASS("panel-heading"),
				DIV(
					CLASS("progress"),
					CLASS("progress-striped"),
					DIV(CLASS("progress-bar"), CLASS("progress-bar-danger"),
						goh4.Style{"width", "{{progressNotAgreed}}%"},
						SPAN(h.translate("PLANNING")),
					),
					DIV(CLASS("progress-bar"), CLASS("progress-bar-success"),
						goh4.Style{"width", "{{progressFinished}}%"},
						SPAN(h.translate("FINISHED")),
					),
				),
			),
			DIV(CLASS("panel-body"),
				ng.Show("currentSection"),
				DIV(CLASS("col-xs-8"),
					ID("state-filters"),
					CLASS("btn-group"),
					h.stateFilter("PLANNING"),
					h.stateFilter("AGREED"),
					h.stateFilter("IMPLEMENTING"),
					h.stateFilter("FINISHED"),
					h.stateFilter("OBSOLET"),
				),
				h.personFilter(),
			),
		),

		DIV(CLASS("row"),
			DIV(CLASS("col-xs-3"),
				DIV(CLASS("panel"), CLASS("panel-default"),
					ng.Show("paragraphs.length > 0"),
					ATTR("droppable", "droppable"),
					ID("paragraph-list"),
					DIV(CLASS("panel-body"),
						UL(CLASS("nav"), CLASS("nav-pills"), CLASS("nav-stacked"),
							CLASS("list-group"),
							goh4.Style{"padding-bottom", "30px"},
							LI(
								CLASS("list-group-item"),
								ng.Repeat("p", "paragraphs"),
								ng.Hide("isFiltered(p.State, p.Responsible)"),
								ID("{{$index}}"),
								ng.Class("sectionIndex == $index ? 'active' : ''"),
								ATTR("draggable", "draggable", "item", "$index", "drop", "handleDropped"),
								A(
									ATTR("data-toggle", "tab"),
									ng.Href("#section"),
									"{{currentSection.charAt(0)}}{{$index+1}} {{p.Title}}",
									ng.Click("setPara()")),
							),
						),
					),
				),
			),
			DIV(CLASS("col-xs-9"),
				DIV(CLASS("tab-content"),
					h.contentInfo(),
					h.contentSection(),
					DIV(ID("validation"), CLASS("tab-pane"),
						H4(h.translate("VALIDAT_ION_RES_PONSE")),
						PRE("{{translate(validationMessage)}}"),
					),
				),
			),
		),
	)
}
