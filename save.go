package main

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
)

func (s *speccer) SetExtension(targetExt string) string {
	ext := filepath.Ext(s.File)
	r := regexp.MustCompile("(" + regexp.QuoteMeta(ext) + ")$")
	return r.ReplaceAllString(s.File, targetExt)
}

func (s *speccer) writeFile(file, content string) error {
	return ioutil.WriteFile(file, []byte(content), 0644)
}

func (s *speccer) writeTranslatedFile(file, content string) error {
	return s.writeFile(file, s.translate(content))
}

func (s *speccer) saveHtml() error {
	return s.writeTranslatedFile(
		s.SetExtension(".html"),
		s.Spec.HtmlFromMarkdown(nil))
}

func (s *speccer) saveMarkdown() error {
	return s.writeTranslatedFile(
		s.SetExtension(".md"),
		s.Spec.Markdown(nil))
}

func (s *speccer) save() error {
	s.Spec.Update()
	err := s.writeFile(s.File, s.Spec.Json())
	if err != nil {
		return err
	}
	err = s.saveMarkdown()
	if err != nil {
		return err
	}
	return s.saveHtml()
}
