package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func readHtmlFromFile(fileName string) (string, error) {

	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func parse(text string) (data []string) {

	tkn := html.NewTokenizer(strings.NewReader(text))

	var vals []string

	var isLi bool

	for {

		tt := tkn.Next()

		switch {

		case tt == html.ErrorToken:
			return vals

		case tt == html.StartTagToken:

			t := tkn.Token()
			isLi = t.Data == "li"

		case tt == html.TextToken:

			t := tkn.Token()

			if isLi {
				vals = append(vals, t.Data)
			}

			isLi = false
		}
	}
}

func main() {
	thema := "dimep"
	style := ""
	var base64Encoding = "data:image/png;base64,"

	if thema == "madis" {
		bytes, err := ioutil.ReadFile("images/madis/Logo.png")
		if err != nil {
			log.Fatal(err)
		}

		base64Encoding += base64.StdEncoding.EncodeToString(bytes)

		style = "<text>\n" +
			"	background: #0074a7;\n" +
			"	background: -moz-linear-gradient(top, #0074a7 0%, #003d7a 100%); /* FF3.6+ */\n" +
			"	background: -webkit-gradient(linear, left top, left bottom, color-stop(0%,#0074a7), color-stop(100%,#003d7a)); /* Chrome,Safari4+ */\n" +
			"	background: -webkit-linear-gradient(top, #0074a7 0%,#003d7a 100%); /* Chrome10+,Safari5.1+ */\n" +
			"	background: -o-linear-gradient(top, #0074a7 0%,#003d7a 100%); /* Opera 11.10+ */\n" +
			"	background: -ms-linear-gradient(top, #0074a7 0%,#003d7a 100%); /* IE10+ */\n" +
			"	background: linear-gradient(to bottom, #0074a7 0%,#003d7a 100%); /* W3C */\n" +
			"	filter: progid:DXImageTransform.Microsoft.gradient( startColorstr='#0074a7', endColorstr='#003d7a',GradientType=0 ); /* IE6-9 */\n" +
			"</text>"
	} else {
		bytes, err := ioutil.ReadFile("images/default/soft_logo.png")
		if err != nil {
			log.Fatal(err)
		}

		base64Encoding += base64.StdEncoding.EncodeToString(bytes)
		style = "background-color: #55ab2e;"
	}

	//fmt.Println(style)
	fileName := "html/comprovanteMarcacao.html"
	text, err := readHtmlFromFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	text = strings.ReplaceAll(text, "/*@Style*/", style)
	text = strings.ReplaceAll(text, "@Model.UrlLogo", base64Encoding)
	text = strings.ReplaceAll(text, "@LabelComprovanteDePonto", "Comprovante de Ponto")
	text = strings.ReplaceAll(text, "@LabelOla", "Olá")
	text = strings.ReplaceAll(text, "@Model.NomeFuncionario", "Alexandre Fernandes")
	text = strings.ReplaceAll(text, "@LabelTituloComprovante", "Segue em anexo o seu comprovante de ponto para a seguinte marcação:")
	text = strings.ReplaceAll(text, "@Model.Marcacao", "19-09-2022 19:10")

	fmt.Println(text)

	//data := parse(text)
	//fmt.Println(data)
}
