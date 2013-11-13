package main

import (
	"flag"
	"fmt"
	"github.com/metakeule/speclib"
)

func (s *speccer) initFlags() {
	flag.StringVar(&s.Args.Language, "language", "",
		"language of the spec that is to be "+
			"edited. Should be something like en_US. "+
			"is only necessary, if there are different languages in use. "+
			"if set, a file with the name spec_[language].json is loaded")
	flag.BoolVar(&s.Args.StrictValidation, "strict", false, "if set, validation will be done in strict mode")
	flag.StringVar(&s.Args.Name, "name", "", "name of the property that should be shown / set")
	flag.StringVar(&s.Args.Set, "set", "", "the value to which the property should be set to")
	flag.BoolVar(&s.Args.Unset, "uset", false, "unset removes the property")
	flag.StringVar(&s.Args.Section, "sec", "", "the section to operate on")
	flag.IntVar(&s.Args.Position, "at", -1, "the position of a paragraph in a session")
	flag.StringVar(&s.Args.Responsible, "resp", "", "the responsible for a paragraph")
	flag.StringVar(&s.Args.Author, "author", "", "the author for a comment")
	flag.IntVar(&s.Args.TargetPosition, "to", -1, "the target position of a paragraph movement")
	flag.StringVar(&s.Args.Usage, "usage", "", "filters for typical use cases as 'documentation', 'approval', 'discussion', 'implementation'")
	flag.StringVar(&s.Args.Filter, "filter", "", "filter to filter out things from html or markdown")
	flag.StringVar(&s.Args.CMD, "cmd", "", "command that should be run")
	flag.BoolVar(&s.Args.IncludeComments, "with-comments", false, "include comments in paragraph output")
	flag.BoolVar(&s.Args.IncludeMeta, "with-meta", false, "include meta information in paragraph output")
}

func (s *speccer) initSections() {
	for _, sec := range speclib.AllSections {
		s.Sections[sec.String()] = true
	}
}

func initV1() {
	Speccer.initFlags()
	Speccer.initSections()
}

func init() { initV1() }

func (s *speccer) init() (err error) {
	err = s.setPWD()
	if err != nil {
		return
	}

	flag.Parse()
	// count 1 down to match ordinals to index 1 => 0, 2 => 1 etc.
	if s.Args.Position > -1 {
		s.Args.Position--
	}
	// count 1 down to match ordinals to index 1 => 0, 2 => 1 etc.
	if s.Args.TargetPosition > -1 {
		s.Args.TargetPosition--
	}

	/*
		if s.Args.CMD == "create" && s.Args.Language == "" {
			s.Args.Language = "en_US"
		}
	*/
	if s.Args.CMD != "create" {
		err = s.findJsonFile()
		if err != nil {
			return
		}

		err = s.loadSpec()
		if err != nil {
			return
		}
	} else {
		s.setJsonFileToCreate()
	}
	err = s.setFilter()
	if err != nil {
		fmt.Println("error while setting filter")
		return
	}
	s.checkSection()
	return nil
}
