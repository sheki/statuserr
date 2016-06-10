package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type status struct {
	Code int
	Name string
}

func (s status) String() string {
	return fmt.Sprintf("%d %s", s.Code, s.Name)
}

const templateStr = `{{.Name}} = err{ Code: {{.Code}}, Message:  "{{.Message}}" }

`

func main() {
	resp, err := http.Get("https://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	r, err := regexp.Compile(`\<h3\>.*(\d{3}) ([\w+\s?]*)\</h3\>`)
	if err != nil {
		log.Fatal(err)
	}
	matches := r.FindAllStringSubmatch(string(b), -1)
	var result []status

	for _, row := range matches {
		i, err := strconv.Atoi(row[1])
		if err != nil {
			log.Fatal(err)
		}
		s := status{
			Code: i,
			Name: row[2],
		}
		result = append(result, s)
	}

	templ, err := template.New("struct").Parse(templateStr)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range result {
		if r.Code < 400 {
			continue
		}
		data := struct {
			Name    string
			Message string
			Code    int
		}{
			Name:    strings.Replace(r.Name, " ", "", -1),
			Message: r.Name,
			Code:    r.Code,
		}
		if err := templ.Execute(os.Stdout, data); err != nil {
			log.Fatal(err)
		}
	}
}
