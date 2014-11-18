package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/metakeule/speccer/webspec/handler"
	utils "github.com/metakeule/utils/http"
)

func my(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(req.URL.String()))
}

var server string

func init() {
	flag.StringVar(&server, "server", "localhost:8282", "host and port where the webserver should run, e.g. localhost:8282")
}

var _ = handler.New

func main() {
	_ = os.Args
	flag.Parse()
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

	//specfile := "spec.json"
	specfile := "spec"
	/*
		if len(os.Args) > 1 {
			//specfile = os.Args[1]
			specfile = os.Args[1]

			//fmt.Println("no file given")
			//os.Exit(1)
		}

		_ = specfile
	*/
	pwd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	//fmt.Printf("%#v\n", pwd)
	// fileserver := http.FileServer(http.Dir(pwd))

	mountpoint := "/spec"

	utils.Mount(
		mountpoint,
		handler.New(
			handler.FileLoader(specfile+".json"),
			handler.OmniSaver(specfile),
		),
	)

	fileDir := filepath.Join(pwd, "file")
	d, e := os.Stat(fileDir)
	if e != nil {
		fmt.Printf("warning: no file directory found\n")
	} else {
		if !d.IsDir() {
			fmt.Printf("warning: %s is no directory\n", fileDir)
		} else {
			fileserver := http.FileServer(http.Dir(fileDir))
			http.Handle("/spec/file/", http.StripPrefix("/spec/file", fileserver))
		}
	}

	fmt.Printf("listening on http://%s%s/admin\n", server, mountpoint)
	err = http.ListenAndServe(server, nil)
	//err = http.ListenAndServe(":"+port, handler.New(fmt.Sprintf("%s", spec)))
	if err != nil {
		panic(err.Error())
	}
}
