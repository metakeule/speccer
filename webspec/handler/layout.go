package handler

import (
	. "github.com/go-on/lib/html"
	"github.com/go-on/lib/internal/ng"
	. "github.com/go-on/lib/types"
)

func (h *handler) layout(content ...interface{}) *DocType {
	container := DIV(Class("container"))
	container.Add(content...)
	return HTML5(
		HTML(
			Attrs_("ng-app", "app"),

			HEAD(
				CharsetUtf8(),
				META(Attrs_("name", "viewport", "content", "width=device-width, initial-scale=1.0")),
				HTMLString(`
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
			BODY(Attrs_("ng-controller", "Spec"), container),
		),
	)
}

func (h *handler) adminlayout() *DocType {
	return h.layout(
		DIV(Class("row"),
			H1(h.translate("SPE_C")+" '{{Sections.OVERVIEW.Title}}'"),
			P(Class("well"), Class("well-sm"), STRONG(h.translate("LASTUPDATE")+" "),
				"{{Sections.OVERVIEW.LastUpdate}}",
				HTMLString("&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"),
				STRONG(h.translate("UUID")+" "), "{{Sections.OVERVIEW.UUID}} "),
		),

		DIV(Id("ok-message"), Class("alert"), Class("alert-success"), ng.Show("ok_message"), "{{ok_message}}"),
		DIV(Id("fail-message"), Class("alert"), Class("alert-danger"),
			Class("alert-dismissable"),
			ng.Show("fail_message"), "{{fail_message}}",
			BUTTON(Attrs_("type", "button", "data-dismiss", "alert", "aria-hidden", "true"), HTMLString("&times;"), Class("close")),
		),

		DIV(Class("row"),
			Class("navbar-default"),
			UL(Class("nav"),
				Class("navbar-nav"),
				//Class("nav-tabs"),
				//Class("nav-pills"),
				LI(A(ng.Href("#info"), ng.Click("unsetCurrentSection()"), Attrs_("data-toggle", "tab"), h.translate("INFO"))),
				LI(A(ng.Href("#section"), Attrs_("data-toggle", "tab"), h.translate("OVERVIEW"),
					BR(),
					SPAN(Class("badge"), Class("badge-warning"), "{{sectionNumbers['OVERVIEW']}}"),
				), ng.Click("setParagraphs('OVERVIEW')")),
				h.dropDownSection("SCENARIO"),
				h.dropDownSection("NONGOAL"),
				h.dropDownSection("UNDECIDED"),
				h.dropDownSection("DEFINITION"),
				h.dropDownSection("CONTRADICTION"),
				h.dropDownSection("FEATURE"),
				LI(
					DIV(Class("navbar-form "),
						A(
							ng.Href("#validation"),
							Class("btn"), Class("btn-default"),
							Attrs_("data-toggle", "tab"), h.translate("VALIDATE")),
						ng.Click("validate()"),
					),
				),
				LI(
					DIV(Class("navbar-form "),
						AHref("spec.html", Attrs_("target", "_blank"),
							// A(
							// ng.Href("#validation"),
							Class("btn"), Class("btn-default"),
							h.translate("SHOW"),
						),
						// Attrs_("data-toggle", "tab"), "Validate"),
						// ng.Click("validate()"),
					),
				),
			),
		),

		BR(),
		DIV(Class("panel"), Class("panel-default"), Id("filter-bar"),
			ng.Show("currentSection"),
			DIV(Class("panel-heading"),
				DIV(
					Class("progress"),
					Class("progress-striped"),
					DIV(Class("progress-bar"), Class("progress-bar-danger"),
						Style{"width", "{{progressNotAgreed}}%"},
						SPAN(h.translate("PLANNING")),
					),
					DIV(Class("progress-bar"), Class("progress-bar-success"),
						Style{"width", "{{progressFinished}}%"},
						SPAN(h.translate("FINISHED")),
					),
				),
			),
			DIV(Class("panel-body"),
				ng.Show("currentSection"),
				DIV(Class("col-xs-8"),
					Id("state-filters"),
					Class("btn-group"),
					h.stateFilter("PLANNING"),
					h.stateFilter("AGREED"),
					h.stateFilter("IMPLEMENTING"),
					h.stateFilter("FINISHED"),
					h.stateFilter("OBSOLET"),
				),
				h.personFilter(),
			),
		),

		DIV(Class("row"),
			DIV(Class("col-xs-3"),
				DIV(Class("panel"), Class("panel-default"),
					ng.Show("paragraphs.length > 0"),
					Attrs_("droppable", "droppable"),
					Id("paragraph-list"),
					DIV(Class("panel-body"),
						UL(Class("nav"), Class("nav-pills"), Class("nav-stacked"),
							Class("list-group"),
							Style{"padding-bottom", "30px"},
							LI(
								Class("list-group-item"),
								ng.Repeat("p", "paragraphs"),
								ng.Hide("isFiltered(p.State, p.Responsible)"),
								Id("{{$index}}"),
								ng.Class("sectionIndex == $index ? 'active' : ''"),
								Attrs_("draggable", "draggable", "item", "$index", "drop", "handleDropped"),
								A(
									Attrs_("data-toggle", "tab"),
									ng.Href("#section"),
									"{{currentSection.charAt(0)}}{{$index+1}} {{p.Title}}",
									ng.Click("setPara()")),
							),
						),
					),
				),
			),
			DIV(Class("col-xs-9"),
				DIV(Class("tab-content"),
					h.contentInfo(),
					h.contentSection(),
					DIV(Id("validation"), Class("tab-pane"),
						H4(h.translate("VALIDAT_ION_RES_PONSE")),
						PRE("{{translate(validationMessage)}}"),
					),
				),
			),
		),
	)
}
