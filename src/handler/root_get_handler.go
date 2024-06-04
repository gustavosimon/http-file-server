package handler

import (
	"http-file-server/src/model"
	"os"
)

// Handler para retornar todos os arquivos disponíveis para serem acessados pelo serviço
func ListRootDirectory() model.HttpResponse {
	files, err := os.ReadDir("./public")
	if err != nil {
		return model.From(500, make(map[string]string), "")
	}
	var body string
	for _, file := range files {
		if file.IsDir() {
			body += "[DIR] "
		} else {
			body += "[FILE] "
		}
		body += file.Name() + "\n"
	}
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
