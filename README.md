# golang-udemy-cod3r-pratice
Aulas práticas do curso de Golang da Cod3r pela Udemy

## Concorrência vs Paralelismo

Go é uma linguagem concorrente

Paralelismo - Executar código simultaneamente em processadores físicos diferentes.

Concorrência - intercalar (administrar) vários processos ao mesmo tempo e isso pode ocorrer em um único processador físico

Concorrência viabiliza paralelismo

É possível que a concorrência seja melhor que o paralelismo!
Parelelismo exige muito mais do SO e do hardware.

Concorrência - é a forma de estruturar o seu programa. É a composição de processos (tipicamente funções) que executam de forma independente.

---
## Github

Atualizar todos pacotes do github
```bash
go get -u all
```
---
## Testes

```bash
$ go test
$ go test ./...

$ go test -cover
$ go test -coverprofile=resultado.out
$ go tool cover -html=resultado.out
```
---
## DATABASE
mysql
user: robb
pass: cr

```go
db, err := sql.Open("mysql", "robb:crmrjr@/cursogo")
defer db.Close()
stmt, _ := db.Prepare("insert into usuarios(nome) values (?)")
res, _ := stmt.Exec("Pedro")
```
---
## HTTP

```go
http.HandleFunc("/usuarios/", UsuarioHandler)
```

## Genial

```go
s := time.now().format("02/01/2006 03:04:05")
```

