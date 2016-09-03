package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

const (
	master  = `Names:{{block "list11" .}}{{"\n"}}{{range .}}{{println "124123412-" .}}{{end}}{{end}}`
	overlay = `{{define "list11"}} {{join . ", "}}{{end}} `
)

var (
	funcs     = template.FuncMap{"join": strings.Join}
	guardians = []string{"Gamora", "Groot", "Nebula", "Rocket", "Star-Lord"}
)

func main() {

	masterTmpl, err := template.New("maste2341r").Funcs(funcs).Parse(master)
	if err != nil {
		log.Fatal(err)
	}
	overlayTmpl, err := template.Must(masterTmpl.Clone()).Parse(overlay)
	if err != nil {
		log.Fatal(err)
	}
	if err := masterTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n\n okkk  \n\n")
	if err := overlayTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
}
