# API feita em GO e com framework Gin

Esta é uma simples API escrita em GO com objetivo de estudos da Linguagem e do framework [Gin](https://gin-gonic.com/). Para versionamento e consultas no Banco de Dados utilizei o ORM [Gorm](https://gorm.io/).

## Funcionalidades

Basicamente, esta API realiza um CRUD de alunos.

## Documentação :book:

A documentação foi feita com Swagger. Após rodar o projeto localmente, a doc ficará disponível neste [link](http://localhost:3000/docs/index.html#/).

## Rodar projeto localmente :computer:

### Pré-requisitos:

- Go (1.21)
- Postgres (necessário criar o arquivo `.env`, seguindo o examplo do `.env.example`)

---

Para iniciar o serviço:

```sh
go run main.go
```

Por padrão, o serviço iniará na porta 3000. Caso seja necessário mudar, é possível alterar no arquivo `main.go`.

## Testes :hammer:

Utilizo testes de integração com o banco de dados, portanto é necessário a aplicação estar conectada ao banco para rodar os testes. Na raíz, rodar o comando:


```sh
go test
```