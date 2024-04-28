package model

// Estrutura de uma requisição HTTP
type HttpRequest struct {
	Method      string
	Resource    string
	HttpVersion string
	Headers     map[string]string
	Body        string
}

// Método para criar um objeto HttpRequest a partir de um payload recebido pela requisição
func CreateFromPayload(payload string) HttpRequest {

	// TODO: implementar a regra para construir o objeto HttpRequest

	var headers = make(map[string]string)
	headers["Host"] = "localhost:8080"
	return HttpRequest{
		Method:      "GET",
		Resource:    "/",
		HttpVersion: "1.1",
		Headers:     headers,
	}
}
