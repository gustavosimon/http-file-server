package handler

import (
	"http-file-server/src/model"
	"os"
)

// Handler para retornar o arquivo informado por parãmetro
func ListFile(resource string) model.HttpResponse {
	file, err := os.ReadFile("./public/" + resource)
	if err != nil {
		return model.From(500, make(map[string]string), "")
	}
	return model.From(200, make(map[string]string), string(file[:]))
}
