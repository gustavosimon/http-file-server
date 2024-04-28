package model

import "strings"

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
	lines := strings.Split(payload, "\n")
	method, resource, httpVersion := parseRequestLine(lines[0])
	headers := parseHeaders(lines[1:])

	body := ""
	if len(lines) > 0 {
		body = lines[len(lines)-1]
	}

	return HttpRequest{
		Method:      method,
		Resource:    resource,
		HttpVersion: httpVersion,
		Headers:     headers,
		Body:        body,
	}
}

func parseRequestLine(requestLine string) (method, resource, httpVersion string) {
	parts := strings.Split(requestLine, " ")
	return parts[0], parts[1], parts[2]
}

func parseHeaders(headerLines []string) map[string]string {
	headers := make(map[string]string)
	for _, line := range headerLines {
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) < 2 {
			continue
		}
		headers[parts[0]] = parts[1]
	}
	return headers
}
