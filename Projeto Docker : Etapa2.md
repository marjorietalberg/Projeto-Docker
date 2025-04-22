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





