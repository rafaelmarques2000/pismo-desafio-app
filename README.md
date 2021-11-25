# Pismo-Desafio-Api

## Detalhes Técnicos

- Container Engine: **Docker 20.10.7**
- EnvManager: **Docker Compose 20.10.7**
- Linguagem: **Golang 1.7**
- Db: **Sqllite3(embbeded**
- ORM: **goorm**
- Router **Api: Mux**


## Como executar o projeto
Siga as instruções abaixo para iniciar o projeto

- Execute **make serve**
- Para acessar a api acesse **http://localhost:8081**
- Para rebuild do projeto: **make rebuild**
- Para executar tests: **make run-all-tests**



## Estrutura Projeto

- config: configurações (Banco de dados, seed e outros)
- pkg: Toda a base de codigo se encontra neste pacote
- server: Configurações do servidor



## Endpoints

Criar Conta
<br>
Endpoint: http://localhost:8081/accounts <br>
Method: POST<br>

Request Body:<br>

```json
{
  "document_number":"559998989898"
}
```

Response:<br>
HttpStatus: 201(Created)
Criado com sucesso
```json
{
  "msg": "Conta criada com sucesso",
  "data": {
    "id": 3,
    "document_number": "3434343"
  }
}
```
HttpStatus: 422(Unprocessable Entity) Caso a conta já exista
```json
{
  "msg": "conta já cadastrada"
}
```

---

Consultar Conta
<br>
Endpoint: http://localhost:8081/accounts/{account_id} <br>
Method: GET<br>


HttpStatus: 200(OK)
Obteve com sucesso a conta cadastrada
```json
{
  "msg": "",
  "data": {
    "id": 1,
    "doc_number": "559998989898"
  }
}
```
HttpStatus: 422(Unprocessable Entity) Caso a conta Não exista
```json
{
  "msg": "record not found"
}
```

---

Efetuar transação
<br>
Endpoint: http://localhost:8081/transactions <br>
Method: POST<br>

Request Body:<br>

```json
{
  "account_id": 1,
  "operation_type_id": 1,
  "amount": -120.80
}
```
Response:<br>
HttpStatus: 201(Created)
Transação realizada com sucesso
```json
{
  "msg": "Transação realizada com sucesso",
  "data": {
    "id": 1,
    "account_id": 1,
    "operation_type_id": 3,
    "amount": -120.8,
    "event_date": "2021-11-25T03:30:45.057658681Z"
  }
}
```
HttpStatus: 422(Unprocessable Entity) Caso a conta Não exista ou Operação
```json
{
  "msg": "record not found"
}
```

HttpStatus: 422(Unprocessable Entity) Caso uma transação seja inválida
```json
{
  "msg": "transação invalida para a operação selecionada"
}
```