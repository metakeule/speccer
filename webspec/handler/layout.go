package handler

import (
	"github.com/metakeule/goh4"
	. "github.com/metakeule/goh4/tag"
	. "github.com/metakeule/goh4/tag/short"
	"github.com/metakeule/ng"
)

var layout = HTML5(
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
		CssHref("/public/css/app.css"),
		CssHref("/public/css/codemirror.css"),
		CssHref("/public/css/codemirror-webspec.theme.css"),
		JsSrc("/public/js/jquery-1.10.2.min.js"),
		JsSrc("/public/js/codemirror.js"),
		JsSrc("/public/js/codemirror-mode/markdown/markdown.js"),
		JsSrc("/public/js/angular.min.js"),
		JsSrc("/public/js/angular-animate.min.js"),
		JsSrc("/public/js/angular-route.min.js"),
		JsSrc("/public/js/markdown.min.js"),
		JsSrc("/public/js/bootstrap.js"),
		JsSrc("/public/js/app.js"),
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
					// ['PLANNING', 'AGREED', 'IMPLEMENTING', 'FINISHED', 'OBSOLET' ];
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

					//LABEL(CLASS("btn"), ng.Class("filter.AGREED ? 'active btn-danger' : 'btn-success'"),
					LABEL(CLASS("btn"), ng.Class("filterClass('AGREED')"),
						InputCheckbox("filter-AGREED", ng.Model("filter.AGREED"), goh4.Style{"display", "none"},
							ng.Click("setFilter()"),
						),
						"{{count.AGREED}} AGREED",
					),

					//LABEL(CLASS("btn"), ng.Class("filter.IMPLEMENTING ? 'active btn-danger' : 'btn-success'"),
					LABEL(CLASS("btn"), ng.Class("filterClass('IMPLEMENTING')"),
						InputCheckbox("filter-IMPLEMENTING", ng.Model("filter.IMPLEMENTING"), goh4.Style{"display", "none"},
							ng.Click("setFilter()"),
						),
						"{{count.IMPLEMENTING}} IMPLEMENTING",
					),

					//LABEL(CLASS("btn"), ng.Class("filter.FINISHED ? 'active btn-danger' : 'btn-success'"),
					LABEL(CLASS("btn"), ng.Class("filterClass('FINISHED')"),
						InputCheckbox("filter-FINISHED", ng.Model("filter.FINISHED"), goh4.Style{"display", "none"},
							ng.Click("setFilter()"),
						),
						"{{count.FINISHED}} FINISHED",
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
						ID("paragraph-list"),
						/*
							handleDropped($index)
						*/
						//ATTR("droppable", "droppable", "bin", "paragraphBin"),
						ATTR("droppable", "droppable"),
						goh4.Style{"padding-bottom", "30px"},
						/*
							LI(
								A(
									ATTR("data-toggle", "tab"),
									ng.Href("#section"),
									"new",
									ng.Show("(currentSection && currentSection != 'OVERVIEW')"),
									ng.Href("#section"),
									ng.Click("newPara()")),
							),
						*/
						LI(
							ng.Repeat("p", "paragraphs"),
							ID("{{$index}}"),
							ng.Class("sectionIndex == $index ? 'active' : ''"),
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
