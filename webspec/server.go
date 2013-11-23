package main

import (
	"fmt"
	"github.com/metakeule/speccer/webspec/handler"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("no file given")
		os.Exit(1)
	}
	spec, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err.Error())
	}
	var port = "8182"
	fmt.Printf("listening on http://localhost:%s\n", port)
	err = http.ListenAndServe(":"+port, handler.New(fmt.Sprintf("%s", spec)))
	if err != nil {
		panic(err.Error())
	}
}
