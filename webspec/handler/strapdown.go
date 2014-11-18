package handler

import (
	. "github.com/go-on/lib/html"
	. "github.com/go-on/lib/html/internal/element"
	. "github.com/go-on/lib/types"

	"github.com/metakeule/speclib"
)

func xmp(i ...interface{}) (e *Element) {
	e = NewElement("xmp", WithoutEscaping)
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

func themeSelect(selected string) *Element {
	s := SELECT(Id("selecttheme"))
	for i := 0; i < len(themes); i++ {
		t := themes[i]
		if t == selected {
			s.Add(OPTION(t, Attrs_("selected", "selected")))
		} else {
			s.Add(OPTION(t))
		}
	}
	return s
}

func (h *handler) usageSelect(selected string) *Element {
	s := SELECT(Id("selectusage"))
	h.SpecLock.RLock()
	defer h.SpecLock.RUnlock()
	for i := 0; i < len(usages); i++ {
		t := usages[i]
		if t == selected {
			s.Add(OPTION(translate(h.Spec.INFO.Language, t), Attrs_("selected", "selected", "value", t)))
		} else {
			s.Add(OPTION(translate(h.Spec.INFO.Language, t), Attrs_("value", t)))
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
