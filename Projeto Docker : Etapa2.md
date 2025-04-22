#  Criando e utilizando volumes para persistência de dados

✅ Objetivo:
Executar um container MySQL e conectar com um backend (Express) e frontend (React), salvando os dados em um volume persistente, para que eles não se percam se o container parar ou for recriado.

---

## Criando e utilizando volumes para persistência de dados com MySQL
### Etapas :
### ✅ 1. Criar volume:
```bash
docker volume create mysql_data
```
🔍 Cria um volume nomeado mysql_data que será usado para armazenar os dados do banco em /var/lib/mysql.
### ✅ 2. Executar o container MySQL com volume:
```bash
docker run -d \
  --name mysql-container \
  -e MYSQL_ROOT_PASSWORD=rootpass \
  -e MYSQL_DATABASE=meubanco \
  -v mysql_data:/var/lib/mysql \
  -p 3307:3306 \
  mysql:latest
```
### 🧠 Explicação:

-d: modo detached (em segundo plano)
-name: nome do container
-e: define variáveis de ambiente
-v: mapeia volume local para o container
-p 3307:3306: redireciona porta da máquina (3307) para o MySQL (3306)

<img src="https://github.com/user-attachments/assets/3a09f049-7a80-41e8-99ca-3b126f7c5805" alt="Image" style="max-width: 100%; height: auto;">

---

### ✅ 3. Verificar se o volume está funcionando:
```bash 
docker volume inspect mysql_data
```
  <img src="https://github.com/user-attachments/assets/839fb2bd-4105-445f-a332-0eb0b5a589ea" alt="Image" style="max-width: 100%; height: auto;">

---
   
### 📌 Extra: Remova o container mas mantenha os dados
```bash
docker stop mysql-container
docker rm mysql-container
```
E recrie com o mesmo volume. O banco continuará com os dados!

---
### Criando e rodando um container multi-stage com Go

🎯 Objetivo
Utilizar multi-stage builds para otimizar uma aplicação Go, reduzindo o tamanho da imagem final. Vamos usar o projeto GS PING, que você pode ter desenvolvido em Golang, como exemplo.

### Passo a Passo
1. Criar o projeto Go básico
Primeiro, se você não tem o projeto Go pronto, vamos criar um exemplo simples de aplicação, chamado GS PING.

a. Criando a estrutura do projeto:
```bash
mkdir gs-ping
cd gs-ping
touch main.go

```
  <img src="https://github.com/user-attachments/assets/39d1a3ea-a080-40f2-abf9-dca22207383d" alt="Image" style="max-width: 100%; height: auto;">

---
### Para abrir o editor use :
```bash
nano main.go

```
### b. Escrever o código básico da aplicação:

Abra o arquivo main.go e insira o seguinte código de exemplo: 
```bash


import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ping OK!")
}

func main() {
	http.HandleFunc("/ping", handler)
	http.ListenAndServe(":8080", nil)
}

```
 <img src="https://github.com/user-attachments/assets/a4e98d26-cf9a-4faa-83c3-bd7708ce664a" alt="Image" style="max-width: 100%; height: auto;">

---
### 2. Criar o Dockerfile com multi-stage build
O multi-stage build permite que você construa sua aplicação Go em uma etapa (build stage), e depois copie apenas o que for necessário para uma imagem mais limpa e leve na etapa final (runtime stage).

a. Criar o arquivo Dockerfile:

```bash
touch Dockerfile

```
b. Escrever o conteúdo do Dockerfile:
```bash

FROM golang:1.18 AS builder

WORKDIR /go/src/app
COPY . .

RUN go mod init gs-ping && \
    go mod tidy && \
    go build -o gs-ping .

FROM debian:bullseye-slim

WORKDIR /app
COPY --from=builder /go/src/app/gs-ping .

EXPOSE 8080
CMD ["./gs-ping"]

```

<img src="https://github.com/user-attachments/assets/b9bbd9c0-423f-4449-bdd3-8708b15292ec" alt="Image" style="max-width: 100%; height: auto;">

### 🧠 Explicação:
Stage 1 (Build):

FROM golang:1.18 AS builder: Usa a imagem oficial do Go para construir a aplicação.

WORKDIR /go/src/app: Define o diretório de trabalho no container.

COPY . .: Copia todos os arquivos do diretório local para o container.

RUN go mod init gs-ping && go mod tidy && go build -o gs-ping .: Inicializa o módulo Go, baixa as dependências e compila o aplicativo Go.

Stage 2 (Runtime):

FROM debian:bullseye-slim: Usa uma imagem mais leve baseada no Debian.

WORKDIR /app: Define o diretório de trabalho.

COPY --from=builder /go/src/app/gs-ping .: Copia a aplicação compilada do estágio de build.

EXPOSE 8080: Expõe a porta onde o servidor Go estará escutando.

CMD ["./gs-ping"]: Executa o binário compilado na inicialização do container.

---

### 3. Criar a imagem Docker
No diretório onde está o seu Dockerfile, execute o comando abaixo para construir a imagem:
```bash
docker build -t gs-ping .

```
Esse comando cria uma imagem Docker chamada gs-ping.

### 4. Rodar o container a partir da imagem
Agora que você tem a imagem, você pode rodá-la em um container. Use o seguinte comando:
```bash
docker run -d -p 8080:8080 --name gs-ping-container gs-ping

```

### 🧠 Explicação:
-d: Roda o container em segundo plano (modo detached).

-p 8080:8080: Mapeia a porta 8080 do container para a porta 8080 do host.

--name gs-ping-container: Nomeia o container como gs-ping-container.

gs-ping: Nome da imagem que acabamos de construir.

### 5. Testar o funcionamento da aplicação
Abra o navegador ou use o curl para acessar a aplicação rodando no container:
```bash
curl http://localhost:8080/ping
```
 <img src="https://github.com/user-attachments/assets/b36aea10-844c-4d6c-a907-2d803f28e1ab" alt="Imagem 1" style="max-width: 100%; height: auto; display: block; margin-bottom: 20px;">
 
  <img src="https://github.com/user-attachments/assets/a4960611-4475-4fb1-b76c-fd17995ea02c" alt="Imagem 2" style="max-width: 100%; height: auto; display: block;">

---
