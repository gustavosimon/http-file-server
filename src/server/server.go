package server

import (
	"fmt"
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
	fmt.Println(request)
}
