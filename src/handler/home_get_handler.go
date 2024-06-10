package handler

import (
	"http-file-server/src/model"
	"os"
)

// Handler para retornar a página home da aplicação
func HomePage() model.HttpResponse {
	var page, err = os.ReadFile("./public/home.html")
	if err != nil {
		return model.From(500, make(map[string]string), "")
	}
	return model.From(200, make(map[string]string), string(page[:]))
}
