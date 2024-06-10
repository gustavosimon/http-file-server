package handler

import (
	"http-file-server/src/model"
	"os"
)

// Handler para retornar a primeira página da aplicação
func FirstPage() model.HttpResponse {
	var page, err = os.ReadFile("./public/page1.html")
	if err != nil {
		return model.From(500, make(map[string]string), "")
	}
	return model.From(200, make(map[string]string), string(page[:]))
}
