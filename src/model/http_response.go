package model

import "strconv"

// Estrutura de uma resposta HTTP
type HttpResponse struct {
	HttpVersion       string
	StatusCode        int
	StatusDescription string
	Headers           map[string]string
	Body              string
}

// Constrói um objeto HttpResponse assumindo defaults
func From(status int, headers map[string]string, body string) HttpResponse {
	return HttpResponse{
		HttpVersion:       "HTTP/1.1",
		StatusCode:        status,
		StatusDescription: getStatusCodeDescription(status),
		Headers:           headers,
		Body:              body,
	}
}

// Retorna a descrição do status code HTTP informado
func getStatusCodeDescription(status int) string {
	switch status {
	case 200:
		return "OK"
	case 404:
		return "Not Found"
	case 405:
		return "Method Not Allowed"
	case 500:
		return "Internal Server Error"
	default:
		return "Internal Server Error"
	}
}

// Converte o objeto HttpResponse em uma string reconhecida pelo protocolo HTTP
func (response HttpResponse) String() string {
	var rawResponse string
	rawResponse = response.HttpVersion + " " + strconv.Itoa(response.StatusCode) + " " + response.StatusDescription + "\n"
	for key, value := range response.Headers {
		rawResponse += key + ": " + value + "\n"
	}
	rawResponse += "\n" + response.Body
	return rawResponse
}
