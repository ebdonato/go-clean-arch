# Desafio Clean Architecture

>Olá devs!
>Agora é a hora de botar a mão na massa. Para este desafio, você precisará criar o usecase de listagem das orders.
>Esta listagem precisa ser feita com:
>
>- Endpoint REST (GET /order)
>- Service ListOrders com GRPC
>- Query ListOrders GraphQL
>    Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.
>
>Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.
>Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.

## Executando o projeto

Primeiramente baixe as dependências do projeto, execute o comando `go mod tidy`.

Para subir os serviços usando o Docker, execute o comando `docker compose -f "docker-compose.yaml" up -d`.

Os seguintes serviços serão executados:

- MySQL
- RabbitMQ
- webBD

> O serviço [webDB](https://webdb.app/) é opcional pois se trata de um  Sistema de Gerenciamento de Banco de Dados.

Finalmente, para iniciar o projeto, execute o comando `go run main.go wire_gen.go` a partir da pasta `./cmd/ordersystem`

> O banco de dados será configurado por migrações automaticas no início da aplicação.

## Portas

Por padrão, o projeto utiliza as seguintes portas:

- web server na porta 8000
- gRPC server na porta 50051
- GraphQL server na porta 8080

> Essas portas podem ser ajustas no arquivo `cmd/ordersystem/.env`

## Extras

Para testar as chamadas REST ou GrahpQL, pode se usar a extensão REST Client no VSCode e utilizar as chamadas descritas no arquivo `api/ordersystem.http`

Para testar gRPC pode se utilizar o [Evans](https://github.com/ktr0731/evans) (instalando localmente) ou [grpcui](https://github.com/fullstorydev/grpcui) (rodando o repositorio go localmente)
