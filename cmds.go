package main

import (
	"bytes"
	"fmt"
	"github.com/metakeule/speclib"
	"io/ioutil"
	"regexp"
	"strings"
)

func (s *speccer) runCMD() error {
	switch s.Args.CMD {
	// operations on the spec
	case "create":
		return s.createCMD()
	case "property":
		return s.propertyCMD()
	case "validate":
		return s.validateCMD()
	case "html":
		return s.htmlCMD()
	case "markdown":
		return s.markdownCMD()
	case "save":
		return s.saveCMD()
	// operations on a section / paragraph of a section
	case "text":
		return s.textCMD()
	case "meta":
		return s.metaCMD()
	case "add-multi":
		return s.addMultiCMD()
	case "comment-multi":
		return s.commentMultiCMD()
	case "pack":
		return s.packCMD()
	case "add":
		return s.addCMD()
	case "rm":
		return s.removeCMD()
	case "comment":
		return s.commentCMD()
	case "move":
		return s.moveCMD()
	case "positions":
		return s.positionsCMD()
	default:
		return fmt.Errorf("unknown command %#v", s.Args.CMD)
	}
}

func (s *speccer) createCMD() error {
	s.shouldSave = true
	if s.Args.Language == "" {
		s.Args.Language = "en_US"
		//	return fmt.Errorf("no language set")
	}
	s.Spec = speclib.New("YYYY-MM-DD", s.Args.Language, "")
	return nil
}

func (s *speccer) propertyCMD() error {
	if s.Args.Unset || s.Args.Set != "" {
		s.shouldSave = true
		if s.Args.Name == "" {
			return fmt.Errorf("no name set")
		}
		if !properties[s.Args.Name] {
			return fmt.Errorf("unknown name, spell all caps")
		}
	}
	if s.Args.Unset {
		if !unsettableProperties[s.Args.Name] {
			return fmt.Errorf("can't unset %s", s.Args.Name)
		}
		return s.unsetProperty(s.Args.Name)
	}

	if s.Args.Set != "" {
		return s.setProperty(s.Args.Name, s.Args.Set)
	}

	if s.Args.Name != "" {
		s.shouldPrint = s.getProperty(s.Args.Name)
	} else {
		s.shouldPrint = s.allProperties()
	}
	return nil
}

func (s *speccer) validateCMD() error {
	return s.Spec.Validate(s.Args.StrictValidation)
}

func (s *speccer) htmlCMD() error {
	s.shouldPrint = s.translate(s.Spec.HtmlFromMarkdown(s.Filter))
	return nil
}

func (s *speccer) markdownCMD() error {
	s.shouldPrint = s.translate(s.Spec.Markdown(s.Filter))
	return nil
}

func (s *speccer) setParagraph() error {
	if s.Args.Section == "" {
		return fmt.Errorf("no section given")
	}
	if s.Args.Section == "OVERVIEW" {
		s.Paragraph = s.Spec.Paragraph
		s.IsOverview = true
		return nil
	}
	if s.Args.Position < 0 {
		return fmt.Errorf("no position given")
	}
	ps := s.Spec.GetSection(speclib.SectionObj[s.Args.Section])
	if s.Args.Position >= len(ps) {
		return fmt.Errorf("no such position in %#v (too large)", s.Args.Section)
	}
	s.Paragraph = ps[s.Args.Position]
	return nil
}

func (s *speccer) mustSetParagraph() {
	err := s.setParagraph()
	if err != nil {
		panic(err.Error())
	}
}

func (s *speccer) readSetFile() (string, error) {
	data, err := ioutil.ReadFile(s.Args.Set)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", data), nil
}

func (s *speccer) removeCMD() error {
	s.shouldSave = true
	if s.Args.Section == "" {
		return fmt.Errorf("no section given")
	}
	if s.Args.Section == "OVERVIEW" {
		return fmt.Errorf("OVERVIEW section can't be removed, try to set it")
	}
	if s.Args.Position < 0 {
		return fmt.Errorf("no position given")
	}
	s.Spec.RemoveParagraph(speclib.SectionObj[s.Args.Section], s.Args.Position)
	return nil
}

func (s *speccer) textCMD() error {
	err := s.setParagraph()
	if err != nil {
		return err
	}

	if s.IsOverview && s.Args.Unset {
		return fmt.Errorf("can't remove OVERVIEW")
	}
	if s.Args.Unset {
		s.shouldSave = true
		s.Spec.RemoveParagraph(speclib.SectionObj[s.Args.Section], s.Args.Position)
		return nil
	}
	if s.Args.Set != "" {
		s.shouldSave = true
		text, e := s.readSetFile()
		if e != nil {
			return e
		}
		s.Paragraph.Text = text
		s.Paragraph.Update()
		return nil
	}
	filter := &speclib.Filter{}
	filter.META = true
	if s.Args.IncludeMeta {
		filter.META = false
	}
	filter.COMMENTS = true
	if s.Args.IncludeComments {
		filter.COMMENTS = false
	}
	fmt.Println(s.Paragraph.Markdown(filter))
	return nil
}

func (s *speccer) metaCMD() error {
	err := s.setParagraph()
	if err != nil {
		return err
	}

	if s.Args.Set != "" {
		s.shouldSave = true
		if s.Args.Name == "" {
			return fmt.Errorf("no name set")
		}

		if !metas[s.Args.Name] {
			return fmt.Errorf("unknown name, spell all caps")
		}
		err = s.setMeta(s.Args.Name, s.Args.Set)
		if err != nil {
			return err
		}
		s.Paragraph.Update()
		return nil
	}
	if s.Args.Name != "" {
		s.shouldPrint = s.getMeta(s.Args.Name)
	} else {
		s.shouldPrint = s.allMetas()
	}
	return nil
}

var hashMatch = regexp.MustCompile(`^(#+\s)$`)

func (s *speccer) addCMD() error {
	s.shouldSave = true
	if s.Args.Section == "" {
		return fmt.Errorf("no section given")
	}
	if s.Args.Section == "OVERVIEW" {
		return fmt.Errorf("can't add to OVERVIEW, try spec text instead")
	}
	if s.Args.Responsible == "" {
		return fmt.Errorf("no responsible given")
	}
	if s.Args.Set == "" {
		return fmt.Errorf("no set given")
	}
	all, err := s.readSetFile()
	if err != nil {
		return err
	}
	all = speclib.NormalizeLineFeeds(all)
	return s.addParagraph(all)
}

func (s *speccer) saveCMD() error {
	s.shouldSave = true
	return nil
}

func (s *speccer) addParagraph(def string) error {
	sep := strings.Index(def, "\n")
	if sep == -1 {
		return fmt.Errorf("can't find linefeed")
	}
	title := def[:sep]
	title = hashMatch.ReplaceAllString(title, "")
	title = "## " + title
	text := def[sep+1:]
	text = speclib.NormalizeLineFeeds(text)
	text = strings.Trim(text, "\n")
	p := s.Spec.NewParagraph(s.Args.Responsible, title, text)
	s.Spec.AddParagraph(speclib.SectionObj[s.Args.Section], p)
	return nil
}

func (s *speccer) addComment(c string) error {
	sep := strings.Index(c, "\n")
	if sep == -1 {
		return fmt.Errorf("can't find linefeed")
	}
	author := c[:sep]
	text := c[sep+1:]
	text = speclib.NormalizeLineFeeds(text)
	text = strings.Trim(text, "\n")
	s.Paragraph.SetComment(author, text)
	return nil
}

func (s *speccer) addMultiCMD() error {
	s.shouldSave = true
	if s.Args.Section == "" {
		return fmt.Errorf("no section given")
	}
	if s.Args.Section == "OVERVIEW" {
		return fmt.Errorf("can't add to OVERVIEW, try spec text instead")
	}
	if s.Args.Responsible == "" {
		return fmt.Errorf("no responsible given")
	}
	if s.Args.Set == "" {
		return fmt.Errorf("no set given")
	}
	all, err := s.readSetFile()
	if err != nil {
		return err
	}
	all = speclib.NormalizeLineFeeds(all)
	pDefs := strings.Split(all, "\n***")
	for _, pdef := range pDefs {
		e := s.addParagraph(pdef)
		if e != nil {
			return e
		}
	}
	return nil
}

func (s *speccer) packCMD() error {
	return nil
}

func (s *speccer) commentMultiCMD() error {
	err := s.setParagraph()
	if err != nil {
		return err
	}

	if s.Args.Set != "" {
		return fmt.Errorf("no set given")
	}
	text, e := s.readSetFile()
	if e != nil {
		return e
	}
	text = speclib.NormalizeLineFeeds(text)
	cs := strings.Split(text, "\n***")
	for _, c := range cs {
		e := s.addComment(c)
		if e != nil {
			return e
		}
	}
	return nil
}

func (s *speccer) commentCMD() error {
	err := s.setParagraph()
	if err != nil {
		return err
	}

	if s.Args.Set != "" || s.Args.Unset {
		s.shouldSave = true
		if s.Args.Author == "" {
			return fmt.Errorf("no author given")
		}
	}

	if s.Args.Unset {
		s.Paragraph.RemoveComment(s.Args.Author)
		return nil
	}
	if s.Args.Set != "" {
		text, e := s.readSetFile()
		if e != nil {
			return e
		}
		s.Paragraph.SetComment(s.Args.Author, text)
		return nil
	}

	if s.Args.Author != "" {
		s.shouldPrint = s.Paragraph.Comments[s.Args.Author]
	} else {
		s.shouldPrint = s.commentsForParagraph()
	}
	return nil
}

func (s *speccer) commentsForParagraph() string {
	var buffer bytes.Buffer
	var first = true
	for author, comment := range s.Paragraph.Comments {
		pre := "***\n"
		if first {
			pre = ""
			first = false
		}
		fmt.Fprintf(&buffer, "%s%s  \n\n%s\n", pre, author, comment)
	}
	return buffer.String()
}

func (s *speccer) moveCMD() error {
	s.shouldSave = true
	if s.Args.Section == "" {
		return fmt.Errorf("no section given")
	}
	if s.Args.Position < 0 {
		return fmt.Errorf("no position given")
	}

	if s.Args.TargetPosition < 0 {
		return fmt.Errorf("no target_position given")
	}

	if s.Args.Section == "OVERVIEW" {
		return fmt.Errorf("can't move within OVERVIEW (single paragraph)")
	}
	return s.Spec.MoveParagraph(speclib.SectionObj[s.Args.Section], s.Args.Position, s.Args.TargetPosition)
}

func (s *speccer) positionsCMD() error {
	if s.Args.Section == "" {
		s.shouldPrint = s.allPositionsAllSections()
	} else {
		s.shouldPrint = s.allPositions(s.Args.Section)
	}
	return nil
}

func (s *speccer) allPositionsAllSections() string {
	var buffer bytes.Buffer
	for _, sec := range speclib.AllSections {
		// buffer.WriteString("\n" + sec.String() + "\n")
		s._allPositions(&buffer, sec)
	}
	return buffer.String()
}

func (s *speccer) _allPositions(buffer *bytes.Buffer, sec speclib.Section) {
	ps := s.Spec.GetSection(sec)
	for i, p := range ps {
		/*
			end := strings.Index(p.Text, "\n")
			if end == -1 {
				fmt.Fprintf(buffer, "%d %s (%s)\n", i+1, p.Text, p.State)
			} else {
				fmt.Fprintf(buffer, "%d %s (%s)\n", i+1, p.Text[0:end], p.State)
			}
		*/
		fmt.Fprintf(buffer, "%s %d %s (%s)\n", sec.String(), i+1, p.Title, p.State)
	}
}

func (s *speccer) allPositions(sec string) string {
	var buffer bytes.Buffer
	s._allPositions(&buffer, speclib.SectionObj[sec])
	return buffer.String()
}
