package main

import (
	"encoding/json"
	"fmt"
	"github.com/metakeule/speclib"
	"io/ioutil"
	"os"
	"path/filepath"
)

type args struct {
	Language         string
	StrictValidation bool
	Name             string
	Set              string
	Unset            bool
	Section          string
	Position         int
	Responsible      string
	Author           string
	TargetPosition   int
	Filter           string
	CMD              string
	Usage            string
	IncludeComments  bool
	IncludeMeta      bool
}

type speccer struct {
	PWD         string
	File        string
	Sections    map[string]bool
	Spec        *speclib.Spec
	Filter      *speclib.Filter
	Args        *args
	Paragraph   *speclib.Paragraph
	IsOverview  bool
	shouldSave  bool
	shouldPrint string
}

var Speccer = &speccer{
	Sections: map[string]bool{"OVERVIEW": true},
	Spec:     &speclib.Spec{},
	Filter:   &speclib.Filter{},
	Args:     &args{},
}

func (s *speccer) setJsonFileToCreate() {
	if s.Args.Language != "" {
		s.File = filepath.Join(s.PWD, "spec_"+s.Args.Language+".json")
	} else {
		s.File = filepath.Join(s.PWD, "spec.json")
	}
}

func (s *speccer) findJsonFile() error {
	if s.Args.Language != "" {
		s.File = filepath.Join(s.PWD, "spec_"+s.Args.Language+".json")
		info, err := os.Stat(s.File)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return fmt.Errorf("is a dir, but should be a file: %s", s.File)
		}
		return nil
	}
	matches, err := filepath.Glob(filepath.Join(s.PWD, "spec*.json"))
	if err != nil {
		return err
	}
	if len(matches) == 0 {
		return fmt.Errorf("can't find any 'spec*.json' in %#v", s.PWD)
	}

	if len(matches) == 1 {
		s.File = matches[0]
		return nil
	}

	return fmt.Errorf("there is more than one spec*.json file in %s. Please specify the language.", s.PWD)
}

func (s *speccer) loadSpec() error {
	data, err := ioutil.ReadFile(s.File)
	if err != nil {
		return err
	}
	return s.Spec.LoadJson(fmt.Sprintf("%s", data))
}

func (s *speccer) setPWD() (err error) {
	s.PWD, err = os.Getwd()
	return
}

func (s *speccer) setFilter() error {
	if s.Args.Filter != "" {
		err := json.Unmarshal([]byte(s.Args.Filter), s.Filter)
		if err != nil {
			return err
		}
	}
	switch s.Args.Usage {
	case "documentation":
		s.Filter.PROPERTIES = true
		s.Filter.META = true
		s.Filter.COMMENTS = true
		s.Filter.NONGOALS = true
		s.Filter.CONTRADICTIONS = true
		s.Filter.UNDECIDED = true
		s.Filter.SPEC_END = true
		s.Filter.RESOURCES = true
		s.Filter.PLANNING = true
		s.Filter.APPROVED = true
		s.Filter.PARTLY_IMPLEMENTED = true
		s.Filter.OBSOLET = true
	case "approval":
		s.Filter.META = true
		s.Filter.COMMENTS = true
		s.Filter.CONTRADICTIONS = true
		s.Filter.UNDECIDED = true
		s.Filter.SPEC_END = true
		s.Filter.RESOURCES = true
		s.Filter.PLANNING = true
		s.Filter.OBSOLET = true
	case "discussion":
		s.Filter.OVERVIEW = true
		s.Filter.APPROVED = true
		s.Filter.PARTLY_IMPLEMENTED = true
		s.Filter.OBSOLET = true
		s.Filter.SPEC_END = true
	case "implementation":
		s.Filter.UNDECIDED = true
		s.Filter.PLANNING = true
		s.Filter.FULLY_IMPLEMENTED = true
		s.Filter.OBSOLET = true
	}
	return nil
}

func (s *speccer) checkSection() {
	if s.Args.Section != "" {
		if _, has := s.Sections[s.Args.Section]; !has {
			panic("unknown section: " + s.Args.Section + " available sections: " + fmt.Sprintf("%#v\n", s.Sections))
		}
	}
}

func (s *speccer) Run() (err error) {
	err = s.init()
	if err != nil {
		return
	}
	err = s.runCMD()
	if err != nil {
		return
	}
	if s.shouldSave {
		err = s.save()
		if err != nil {
			return
		}
	}
	if s.shouldPrint != "" {
		fmt.Println(s.shouldPrint)
	}
	return
}

func main() {
	err := Speccer.Run()
	if err != nil {
		panic(err.Error())
	}
}
