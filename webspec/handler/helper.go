package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func j(v interface{}) []byte {
	res, err := json.Marshal(v)
	if err != nil {
		panic(err.Error())
	}
	return res
}

func mkSectionJson(section string) []byte {
	return j(spec.Sections[section])
}

func Save() {
	spec.Update()
	err := ioutil.WriteFile(os.Args[1], []byte(spec.Json()), 0644)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

func writeError(rw http.ResponseWriter, err error) {
	rw.WriteHeader(400)
	rw.Write([]byte(err.Error()))

}
