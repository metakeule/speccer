package handler

import (
	"encoding/json"
	// "fmt"
	tr "github.com/metakeule/speccer/translations"

	"net/http"
	"strings"
)

func j(v interface{}) []byte {
	res, err := json.Marshal(v)
	if err != nil {
		panic(err.Error())
	}
	return res
}

func (h *handler) mkSectionJson(section string) []byte {
	return j(h.Spec.Sections[section])
}

func writeError(rw http.ResponseWriter, err error) {
	rw.WriteHeader(400)
	rw.Write([]byte(err.Error()))

}

var (
	translations = map[string]map[string]string{
		"de_DE": tr.Trans_de_DE,
		"en_US": tr.Trans_en_US,
	}
)

// it does not translate, if there is no translation
func translate(language string, in string) (out string) {
	tr, has := translations[language]
	if !has {
		return in
	}
	out = in
	for k, v := range tr {
		out = strings.Replace(out, k, v, -1)
	}
	return
}

func (h *handler) translate(in string) (out string) {
	return translate(h.Spec.INFO.Language, in)
}
