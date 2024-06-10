# Sistemas Distribuídos 2024/01
## Trabalho 1

### Descrição

Implemente a base de um servidor web onde seu servidor deverá,
obrigatoriamente, responder requisições do tipo GET. Desta maneira, deverá ser
possível acessar um site composto por 3 páginas (uma home e duas internas)
ilustradas por um conjunto de imagens. O acesso ao site poderá ocorrer em
qualquer sequência, independente da página inicial. Este site será armazenado
em um diretório disponibilizado por seu servidor (equivalente ao DocumentRoot
do Apache). Caso nenhum arquivo seja apontado na requisição enviada pelo
browser, seu servidor deverá listar o conteúdo de arquivos disponível no
DocumentRoot no momento do acesso, permitindo que arquivos com extensão
html sejam clicados a partir de links.

### Formato
Apresentação individual feita diretamente ao professor. Você rodará seu
servidor, permitindo a execução de testes de navegação a partir de um browser
rodando localmente em sua máquina. O browser deve ser capaz de acessar
todas as 3 páginas, exibindo todo o conteúdo das mesmas. Após a observação
do funcionamento de seu servidor, o professor fará perguntas a respeito de sua
implementação. Estas perguntas poderão, eventualmente, incluir explicações de
partes de seu código-fonte, que deverá ser exibido ao professor.
Critérios de Avaliação (assumindo 10 como nota máxima):
- Implementação das funcionalidades de acesso: 50% da nota (5 pontos)
- Explicação e respostas aos questionamentos do professor relacionados
ao código-fonte: 50% da nota (5 pontos).


## Como executar esse projeto?

### Pré-requisitos
Ter o Go instalado. Versão utilizada go1.19.1.

Estando no diretório root do projeto, executar:

```shell
go run src/main.go
```

Deve gerar um output de Servidor TCP iniciado!