package handler

import (
	"http-file-server/src/model"
	"os"
	"strings"
)

const html_header = `<!DOCTYPE html>
					 <html lang="en">
					 <head>
    					<meta charset="UTF-8">
    					<meta name="viewport" content="width=device-width, initial-scale=1.0">
    					<title>Teste de site</title>
					 </head>
					 <body`

const html_footer = `</body>
   					 </html>`

// Handler para retornar todos os arquivos disponíveis para serem acessados pelo serviço
func ListRootDirectory() model.HttpResponse {
	files, err := os.ReadDir("./public")
	if err != nil {
		return model.From(500, make(map[string]string), "")
	}
	var body string
	body = html_header
	for _, file := range files {
		if file.IsDir() {
			body += "<a> [DIR] " + file.Name() + "</a><br>"
		} else {
			if strings.HasSuffix(file.Name(), ".html") {
				body += "<a href='http://localhost:8080/" + file.Name() + "'>[FILE] " + file.Name() + "</a><br>"
			} else {
				body += "<a> [FILE] " + file.Name() + "</a><br>"
			}
		}
	}
	body += html_footer
	return model.From(200, make(map[string]string), body)
}

// Handler para retornar o arquivo informado por parãmetro
func ListFile(resource string) model.HttpResponse {
	file, err := os.ReadFile("./public/" + resource)
	if err != nil {
		return model.From(500, make(map[string]string), "")
	}
	return model.From(200, make(map[string]string), string(file[:]))
}
