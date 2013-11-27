package handler

import (
	"github.com/metakeule/goh4"
	. "github.com/metakeule/goh4/tag"
	"github.com/metakeule/speclib"
)

func xmp(i ...interface{}) (e *goh4.Element) {
	e = goh4.NewElement("xmp", goh4.WithoutEscaping)
	e.Add(i...)
	return
}

var themes = []string{
	"amelia",
	"cerulean",
	// "cosmo",
	"cyborg",
	// "flatly",
	"journal",
	"readable",
	"simplex",
	"slate",
	"spacelab",
	"spruce",
	"superhero",
	"united",
	// "yeti",
}

var usages = []string{
	"NOFILTER",
	"DOCFILTER",
	"APPROVALFILTER",
	"DISCUSSIONFILTER",
	"IMPLEMENTATIONFILTER",
	"PREVIEWFILTER",
}

func themeSelect(selected string) *goh4.Element {
	s := SELECT(ID("selecttheme"))
	for i := 0; i < len(themes); i++ {
		t := themes[i]
		if t == selected {
			s.Add(OPTION(t, ATTR("selected", "selected")))
		} else {
			s.Add(OPTION(t))
		}
	}
	return s
}

func (h *handler) usageSelect(selected string) *goh4.Element {
	s := SELECT(ID("selectusage"))
	h.SpecLock.RLock()
	defer h.SpecLock.RUnlock()
	for i := 0; i < len(usages); i++ {
		t := usages[i]
		if t == selected {
			s.Add(OPTION(translate(h.Spec.INFO.Language, t), ATTR("selected", "selected", "value", t)))
		} else {
			s.Add(OPTION(translate(h.Spec.INFO.Language, t), ATTR("value", t)))
		}
	}
	return s
}

func filterForUsage(usage string) (f *speclib.Filter) {
	f = &speclib.Filter{}
	switch usage {
	case "DOCFILTER":
		f.PROPERTIES = true
		f.META = true
		f.COMMENTS = true
		f.NONGOAL = true
		f.CONTRADICTION = true
		f.UNDECIDED = true
		f.RESOURCES = true
		f.PLANNING = true
		f.AGREED = true
		f.IMPLEMENTING = true
		f.OBSOLET = true
	case "APPROVALFILTER":
		f.META = true
		f.COMMENTS = true
		f.CONTRADICTION = true
		f.UNDECIDED = true
		f.RESOURCES = true
		f.PLANNING = true
		f.OBSOLET = true
	case "PREVIEWFILTER":
		f.OVERVIEW = true
		f.PROPERTIES = true
		f.META = true
		f.COMMENTS = true
		f.NONGOAL = true
		f.CONTRADICTION = true
		f.UNDECIDED = true
		f.RESOURCES = true
		f.PLANNING = true
		f.FINISHED = true
		f.OBSOLET = true
	case "DISCUSSIONFILTER":
		f.META = true
		f.OVERVIEW = true
		f.PROPERTIES_EXTENDED = true
		f.AGREED = true
		f.IMPLEMENTING = true
		f.FINISHED = true
		f.OBSOLET = true
	case "IMPLEMENTATIONFILTER":
		f.UNDECIDED = true
		f.PROPERTIES_EXTENDED = true
		f.PLANNING = true
		f.FINISHED = true
		f.OBSOLET = true
	default:
		return f
	}
	return f
}
