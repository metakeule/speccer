package main

import (
	"fmt"
	"github.com/metakeule/speccer/webspec/handler"
	utils "github.com/metakeule/utils/http"
	"net/http"
	"os"
)

func my(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(req.URL.String()))
}

var _ = handler.New

func main() {
	_ = os.Args
	/*
		if len(os.Args) != 2 {
			fmt.Println("no file given")
			os.Exit(1)
		}
		spec, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err.Error())
		}
	*/

	/*
		spec1file := "../examples/de/spec_de_DE.json"
		spec2file := "public/js/spec.json"

		var port = "8282"
		//_ = spec

		utils.Mount(
			"/spec1",
			handler.New(
				handler.FileLoader(spec1file),
				handler.FileSaver(spec1file),
			),
		)
		utils.Mount(
			"/spec2",
			handler.New(
				handler.FileLoader(spec2file),
				handler.FileSaver(spec2file),
			),
		)
	*/

	if len(os.Args) != 2 {
		fmt.Println("no file given")
		os.Exit(1)
	}
	var port = "8282"

	specfile := os.Args[1]
	mountpoint := "/spec"

	utils.Mount(
		mountpoint,
		handler.New(
			handler.FileLoader(specfile),
			handler.FileSaver(specfile),
		),
	)

	fmt.Printf("listening on http://localhost:%s%s/admin\n", mountpoint, port)
	err := http.ListenAndServe(":"+port, nil)
	//err = http.ListenAndServe(":"+port, handler.New(fmt.Sprintf("%s", spec)))
	if err != nil {
		panic(err.Error())
	}
}
