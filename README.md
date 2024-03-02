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

Antes de começar a aplicação, lembre de criar a tabela no banco de dados, utilizando por exemplo a _sql query_:

```sql
CREATE TABLE orders (
    id varchar(255) NOT NULL,
    price float NOT NULL,
    tax float NOT NULL,
    final_price float NOT NULL,
    PRIMARY KEY (id)
)
```

Finalmente, para iniciar o projeto, execute o comando `go run main.go wire_gen.go` a partir da pasta `./cmd/ordersystem`
