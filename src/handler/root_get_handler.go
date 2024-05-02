package handler

import (
	"fmt"
	"http-file-server/src/model"
	"io/ioutil"
)

// Handler para retornar todos os arquivos disponíveis para serem acessados pelo serviço
func ListRootDirectory() model.HttpResponse {
	files, err := ioutil.ReadDir("./public")
	if err != nil {
		fmt.Println(err)
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
