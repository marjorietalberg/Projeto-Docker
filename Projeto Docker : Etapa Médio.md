#  Criando e utilizando volumes para persistência de dados

### ✅ Objetivo:
Executar um container MySQL e conectar com um backend (Express) e frontend (React), salvando os dados em um volume persistente, para que eles não se percam se o container parar ou for recriado.

---

# 1️⃣ Criando e utilizando volumes para persistência de dados com MySQL

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
# 2️⃣ Criando e rodando um container multi-stage com Go

### 🎯 Objetivo
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
 <img src="https://github.com/user-attachments/assets/d4b98571-a2cd-4edd-b8b3-54abe59b1fff" alt="Image" style="max-width: 100%; height: auto;">
 

---
# 3️⃣  React + Node.js + MongoDB com Docker Compose

Este repositório documenta o Desafio 7 do curso de Docker, cujo objetivo é executar uma aplicação fullstack com frontend em React, backend em Node.js/Express e banco de dados MongoDB, utilizando o Docker Compose para orquestração de containers.

---
## 📌 Objetivo do desafio
O desafio, proposto no repositório oficial da Docker, disponível em:

🔗 https://github.com/docker/awesome-compose/tree/master/react-express-mongodb

consiste em:

✅ Subir os containers necessários com Docker Compose
✅ Garantir a comunicação entre o frontend, o backend e o MongoDB
✅ Verificar que a aplicação está funcional via navegador
---

🧰 Tecnologias utilizadas
Docker: para conteinerização

Docker Compose: para orquestrar os containers

React: frontend da aplicação

Node.js/Express: backend da aplicação

MongoDB: banco de dados da aplicação

---
Etapas realizadas
1. Clonagem do repositório oficial:
```bash
git clone https://github.com/docker/awesome-compose.git
cd awesome-compose/react-express-mongodb
```
### 📌Instalação do Docker e Docker Compose
No meu ambiente Linux (Ubuntu), instalei o Docker e o Docker Compose:
```bash
sudo apt update
sudo apt install docker.io
sudo apt install docker-compose
```

2. Subida dos containers com Docker Compose:
```bash
docker-compose up -d
```

3.Verificação do status dos serviços:
```bash
docker-compose ps
```

4.Logs do backend para confirmar conexão:
```bash
docker-compose logs backend
```

  <img src="https://github.com/user-attachments/assets/827f632a-d626-4299-9e7e-61791b514939" alt="Imagem 2">

6. Acesso à aplicação
A aplicação ficou disponível localmente nas seguintes portas:

🔗 Backend: http://localhost:3000

🔗 Frontend: http://localhost:3000 (mesma porta, pois React e backend estavam unificados)

  ---
  
  <img src="https://github.com/user-attachments/assets/6c691732-ffba-45e1-a480-cb0ef10e4472" alt="Imagem 1">

  <img src="https://github.com/user-attachments/assets/9d27dd2d-ef26-4159-97e5-62a2a7addee8" alt="Imagem 3">


  ---
 ### 📌 Observações importantes
Foi necessário instalar o distutils manualmente, pois o docker-compose estava dando erro por falta dele:
```bash
sudo apt install python3-distutils
```
O ambiente utilizado foi Ubuntu 22.04 com Python 3.12, o que exige atenção extra para evitar conflitos com pacotes Python (por conta do PEP 668).

---
## 📍 Conclusão
Este desafio teve como objetivo testar a capacidade de rodar uma aplicação full stack com múltiplos serviços via Docker Compose, e tudo foi realizado com sucesso.

🔗 Link do desafio original:
https://github.com/docker/awesome-compose/tree/master/react-express-mongodb

---
# PostgreSQL com pgAdmin via Docker Compose — Ambiente de Banco de Dados Rápido e Gerenciável
🧠 O que vamos fazer?

✅ Clonar um projeto com Docker Compose

✅ Subir um ambiente com PostgreSQL e pgAdmin

✅ Acessar o pgAdmin via navegador

✅ Conectar-se ao banco PostgreSQL

✅ (Opcional) Customizar com um Dockerfile

---
### 🚀 Passo a Passo
🔹 1. Clonar o repositório oficial
Vamos usar um repositório de exemplo da própria Docker Inc., que está no GitHub:

```bash
git clone https://github.com/docker/awesome-compose.git
cd awesome-compose/postgresql-pgadmin
```
🧾 Isso vai baixar o código e entrar na pasta específica do exemplo postgresql-pgadmin.

🔹 2. Verificar a estrutura da pasta
Veja os arquivos com:
```bash
ls

```
