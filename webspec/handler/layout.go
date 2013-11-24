package handler

import (
	"github.com/metakeule/goh4"
	. "github.com/metakeule/goh4/tag"
	. "github.com/metakeule/goh4/tag/short"
	"github.com/metakeule/ng"
)

func personFilter() *goh4.Element {
	return DIV(CLASS("col-xs-4"),
		ng.Show("currentSection"),
		// goh4.Style{"padding-left", "0px"},
		DIV(
			CLASS("row"),
			DIV(CLASS("label"), CLASS("label-default"), "Responsible", CLASS("col-xs-4")),

			DIV(
				CLASS("form-group"),
				CLASS("col-xs-8"),
				SELECT(ng.Model("filteredPerson"),
					CLASS("form-control"),
					CLASS("input-sm"),
					ID("filteredPerson"),
					ng.Change("setCounts()"),
					ATTR("ng-options", "key as key + ' (' + value + ')' for (key , value) in countResponsibles"),
				),
			),
		),
	)
}

func stateFilter(state string) *goh4.Element {
	return LABEL(CLASS("btn"), CLASS("btn-sm"), ng.Class("filterClass('"+state+"')"),
		ng.Show("currentSection"),
		InputCheckbox("filter-"+state, ng.Model("filter."+state+""), goh4.Style{"display", "none"},
			ng.Click("setFilter()"),
		),
		SPAN(CLASS("badge"), "{{count."+state+"}}"),
		" "+state,
	)
}

func layout(content ...interface{}) *DocType {
	container := DIV(CLASS("container"))
	container.Add(content...)
	return HTML5(
		ATTR("ng-app", "app"),
		HEAD(
			CharsetUtf8(),
			META(ATTR("name", "viewport", "content", "width=device-width, initial-scale=1.0")),
			HTML(`
			<!--[if lt IE 9]>
      			<script src="/public/js/html5shiv.js"></script>
      			<script src="/public/js/1.3.0/respond.min.js"></script>
    		<![endif]-->
    	`),
			CssHref("/public/css/bootstrap.css"),
			CssHref("/public/css/codemirror.css"),
			CssHref("/public/css/codemirror-webspec.theme.css"),
			CssHref("/public/css/app.css"),
			JsSrc("/public/js/jquery-1.10.2.min.js"),
			JsSrc("/public/js/codemirror.js"),
			JsSrc("/public/js/codemirror-mode/markdown/markdown.js"),
			JsSrc("/public/js/angular.min.js"),
			JsSrc("/public/js/angular-animate.min.js"),
			JsSrc("/public/js/angular-route.min.js"),
			JsSrc("/public/js/markdown.min.js"),
			JsSrc("/public/js/bootstrap.js"),
			JsSrc("/public/js/app.js"),
			TITLE("Spec for: "+spec.OVERVIEW.Title+" ["+spec.OVERVIEW.UUID+"]"),
		),
		BODY(ATTR("ng-controller", "Spec"), container),
	)
}

func viewLayout() *DocType {
	return layout(
		DIV(CLASS("row"),
			DIV(CLASS("col-xs-1"), H4("SPEC")),
			personFilter(),
			DIV(CLASS("col-xs-9"),
				CLASS("btn-group"),
				BR(),
				stateFilter("PLANNING"),
				stateFilter("AGREED"),
				stateFilter("IMPLEMENTING"),
				stateFilter("FINISHED"),
				stateFilter("OBSOLET"),
			),
		),
		DIV(ID("ok-message"), CLASS("alert"), CLASS("alert-success"), ng.Show("ok_message"), "{{ok_message}}"),
		DIV(ID("fail-message"), CLASS("alert"), CLASS("alert-danger"),
			CLASS("alert-dismissable"),
			ng.Show("fail_message"), "{{fail_message}}",
			BUTTON(ATTR("type", "button", "data-dismiss", "alert", "aria-hidden", "true"), HTML("&times;"), CLASS("close")),
		),
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
					ID("paragraph-list"),
					ATTR("droppable", "droppable"),
					goh4.Style{"padding-bottom", "30px"},
					LI(
						ng.Repeat("p", "paragraphs"),
						ID("{{$index}}"),
						ng.Class("sectionIndex == $index ? 'active' : ''"),
						ATTR("draggable", "draggable", "item", "$index", "drop", "handleDropped"),
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
					// contentInfo(),
					contentInfoView(),
					contentSection(),
					DIV(ID("validation"), CLASS("tab-pane"),
						H4("Validation Response"),
						PRE("{{validationMessage}}"),
					),
				),
			),
		),
	)
}

func adminlayout() *DocType {
	return layout(
		DIV(CLASS("row"),
			H1("SPEC: '{{Sections.OVERVIEW.Title}}'"),
			P(STRONG("LastUpdate: "), "{{Sections.OVERVIEW.LastUpdate}}", HTML("&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"), STRONG("UUID: "), "{{Sections.OVERVIEW.UUID}} "),
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
				LI(A(ng.Href("#info"), ng.Click("unsetCurrentSection()"), ATTR("data-toggle", "tab"), "INFO")),
				LI(A(ng.Href("#section"), ATTR("data-toggle", "tab"), "OVERVIEW"), ng.Click("setParagraphs('OVERVIEW')")),
				dropDownSection("SCENARIO"),
				dropDownSection("NONGOAL"),
				dropDownSection("UNDECIDED"),
				dropDownSection("DEFINITION"),
				dropDownSection("CONTRADICTION"),
				dropDownSection("FEATURE"),
				LI(
					DIV(CLASS("navbar-form "),
						A(
							ng.Href("#validation"),
							CLASS("btn"), CLASS("btn-default"),
							ATTR("data-toggle", "tab"), "Validate"),
						ng.Click("validate()"),
					),
				),
			),
			//A(ng.Href("#validation"), CLASS("btn"), CLASS("btn-default"), ATTR("data-toggle", "tab"), "Validate"), ng.Click("validate()"),
			// A(ng.Href("#validation"), CLASS("btn"), CLASS("btn-default"), ATTR("data-toggle", "tab"), "Validate"), ng.Click("validate()"),
		),

		/*
		   <div class="progress">
		     <div class="progress-bar progress-bar-success" style="width: 35%">
		       <span class="sr-only">35% Complete (success)</span>
		     </div>
		     <div class="progress-bar progress-bar-warning" style="width: 20%">
		       <span class="sr-only">20% Complete (warning)</span>
		     </div>
		     <div class="progress-bar progress-bar-danger" style="width: 10%">
		       <span class="sr-only">10% Complete (danger)</span>
		     </div>
		   </div>
		*/
		BR(),
		DIV(CLASS("panel"), CLASS("panel-default"), ID("filter-bar"),
			ng.Show("currentSection"),
			DIV(CLASS("panel-heading"),
				DIV(
					CLASS("progress"),
					CLASS("progress-striped"),
					DIV(CLASS("progress-bar"), CLASS("progress-bar-danger"),
						goh4.Style{"width", "{{progressNotAgreed}}%"},
						SPAN("planning"),
					),
					DIV(CLASS("progress-bar"), CLASS("progress-bar-success"),
						goh4.Style{"width", "{{progressFinished}}%"},
						SPAN("finished"),
					),
				),
			),
			DIV(CLASS("panel-body"),
				ng.Show("currentSection"),
				// BR(),
				/*
					DIV(CLASS("col-xs-1"),
						LABEL(CLASS("btn"), CLASS("btn-warning"), ng.Click("forcestore()"), "Save"),
					),
				*/
				/*
					DIV(CLASS("col-xs-3"),
						//goh4.Style{"padding-right", "0px"},
						//goh4.Style{"padding-left", "0px"},
						DIV(CLASS("row"),
							//goh4.Style{"margin-right", "0px"},
							DIV(CLASS("col-xs-6"),
								//		goh4.Style{"margin-left", "15px"},
								DIV(CLASS("label"), CLASS("label-danger"), "not agreed: {{progressNotAgreed}}%"),
							),
							DIV(CLASS("col-xs-6"),
								//goh4.Style{"padding-right", "0px"},
								DIV(CLASS("label"), CLASS("label-success"), "finished: {{progressFinished}}%"),
							),
						),
					),
				*/

				DIV(CLASS("col-xs-8"),
					ID("state-filters"),
					//goh4.Style{"background-color", "yellow"},
					// ng.Show("currentSection"),
					CLASS("btn-group"),
					stateFilter("PLANNING"),
					stateFilter("AGREED"),
					stateFilter("IMPLEMENTING"),
					stateFilter("FINISHED"),
					stateFilter("OBSOLET"),
				),
				personFilter(),
			),
		),

		DIV(CLASS("row"),
			DIV(CLASS("col-xs-3"),
				//ng.Show("currentSection"),
				// BR(),
				DIV(CLASS("panel"), CLASS("panel-default"),
					//ng.Show("currentSection && currentSection != 'OVERVIEW' "),
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
									"{{p.Title}}",
									ng.Click("setPara()")),
							),
						),
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
	)
}
