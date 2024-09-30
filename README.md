# Backend Engineer Challenge - API Challenge Pismo

### Ferramentas Utilizadas

- **Golang** (>= 1.7) - Testes e desenvolvimento
- **MongoDB** (>= 2.6.10) - Testes e desenvolvimento
- **Docker** (>= 17.03.2-ce) - Desenvolvimento e produção
- **Redis** (>= 3.0.6) - Desenvolvimento e produção
- **Nginx** (>= 1.0.15) - Produção

### Estrutura de Diretórios

- **/cmd**: Contém o arquivo principal (`main.go`) para execução das chamadas de métodos e testes de conexão.
- **/pkg/api/model**: Contém os arquivos `account.go` e `transaction.go` com as structs usadas no projeto.

### Rotas Principais

- **/router/routes.go**: Define as rotas para os caminhos `/accounts`, `/transactions`, `/health`, e `/home`.

### Integração com MongoDB

- **/util/mg/mongo.go**: Usa o pacote `"go.mongodb.org/mongo-driver/mongo"` para conexão com o banco de dados.

### Controladores

- **/controller/accountController.go**:
    - `CreateAccount()`: Criação de uma nova conta.
    - `GetAccount()`: Busca uma conta pelo ID.
    - `ValidateAccount()`: Validação dos dados da conta.

- **/controller/transactionController.go**:
    - `CreateTransaction()`: Criação de uma nova transação.
    - `ValidateTransaction()`: Validação dos dados da transação.

### Executando o Projeto

#### Com Docker Compose

- Para desenvolvimento: `docker-compose up`
- Para produção: `docker-compose -f docker-compose-prod.yml up`

#### Manualmente

Requisitos: Banco de dados MongoDB em execução.

1. Compile o projeto: `go build ./cmd/main.go`
2. Execute o binário: `./main`

### Testes

Para rodar os testes automatizados:

- Execute na raiz do projeto: `go test ./...`

---

### Exemplos de Uso da API

#### Criar uma Conta

**POST** `/accounts`

- **Request Body**:
  ```json
  {
    "document_number": "12345678900"
  }
  ```

#### Buscar Informações de uma Conta

**GET** `/accounts/:accountId`

- **Response Body**:
  ```json
  {
    "account_id": "66f9ac2a1b1586137fa1561b",
    "document_number": "12345678900"
  }
  ```

#### Criar uma Transação

**POST** `/transactions`

- **Request Body**:
  ```json
  {
    "account_id": "66f9ac2a1b1586137fa1561b",
    "operation_type_id": 4,
    "amount": 123.45
  }
  ```



# Especificações
go version 17 linux/amd64

Mongo

Sistema Operacional: linux Ubuntu 18.4

# Importações
`go build ./...`

### Swagger Documentation ###

  
Swagger disponibilizado em: `{{URL}}:{{PORT}}/swagger/index.html`

# Contatos
thg.mnzs@gmail.com
