# desafio-amparo

- ## Descrição
  - Construir uma API para gerenciar uma lista de tarefas específicas para pessoas que perderam um ente querido. A API deve permitir que os usuários possam ler, atualizar e arquivar tarefas, além de visualizar as próximas 3 tarefas a serem realizadas.
- ## Como executar o projeto
  - Para rodar o projeto é necessário ter o Go instalado na máquina. Para instalar o Go, siga as instruções do site oficial: https://golang.org/doc/install
  - Após instalar o Go, clone o repositório e execute o comando `go run .` na raiz do projeto.
  - A API estará disponível em `http://localhost:8080`
  - ### Endpoints
    - `GET /tasks` - Retorna todas as tarefas cadastradas
    - `GET /tasks?limit=3` - Retorna as próximas 3 tarefas a serem realizadas
    - `PUT /tasks/:id` - Atualiza uma tarefa
    - `DELETE /tasks/:id` - Arquiva uma tarefa
- ### Testes
  - Para rodar os testes, execute o comando `go test ./handlers` na raiz do projeto.
- ## Tecnologias utilizadas
  - Golang
  - Gin
  - Go Test