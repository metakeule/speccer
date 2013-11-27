package main

import (
	tr "github.com/metakeule/speccer/translations"
	"strings"
)

var (
	translations = map[string]map[string]string{
		"de_DE": tr.Trans_de_DE,
		"en_US": tr.Trans_en_US,
	}
)

// it does not translate, if there is no translation
func (s *speccer) translate(in string) (out string) {
	tr, has := translations[s.Spec.INFO.Language]
	if !has {
		return in
	}
	out = in
	for k, v := range tr {
		out = strings.Replace(out, k, v, -1)
	}
	return
}
