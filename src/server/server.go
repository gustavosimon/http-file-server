package server

import (
	"fmt"
	"http-file-server/src/handler"
	"http-file-server/src/model"
	"net"
)

// Endereço em que o serviço irá responder
const HttpServerAddress = "127.0.0.1:8080"

// Protocolo utilizado para a comunicação
const Protocol = "tcp"

// Cria o servidor TCP que irá intepretar as requisições HTTP recebidas
func Serve() {
	listener, err := net.Listen(Protocol, HttpServerAddress)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor: ", err)
	}
	defer listener.Close()
	fmt.Println("Servidor TCP iniciado!")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erro ao receber conexão: ", err)
			return
		}
		handleConnection(conn)
	}
}

// Trata a requisição HTTP recebida
func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	bytesRead, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Falha na leitura dos bytes da requisição: ")
		return
	}
	payload := string(buffer[:bytesRead])
	request := model.CreateFromPayload(payload)
	//
	// Se a requisição for de um método diferente de GET,
	// consideramos um método não permitido
	//
	if request.Method != "GET" {
		var response = model.From(405, make(map[string]string), "")
		conn.Write([]byte(response.String()))
		return
	}
	//
	// Avalia o recurso que foi requisitado e trata com o handler adequado
	//
	var response model.HttpResponse
	switch request.Resource {
	case "/":
		response = handler.ListRootDirectory()
	case "/home":
		response = handler.HomePage()
	case "/page1":
		response = handler.FirstPage()
	case "/page2":
		response = handler.SecondPage()
	default:
		response = handler.ListRootDirectory()
	}
	conn.Write([]byte(response.String()))
}
