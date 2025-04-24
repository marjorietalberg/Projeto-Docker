# 🚀 Projeto: Site Estático com Docker + NGINX + Material Kit (Creative Tim)
---
Este projeto tem como objetivo criar uma imagem Docker personalizada, baseada no servidor NGINX, que hospeda um site HTML/CSS moderno. Utilizamos como base a landing page open-source Material Kit da Creative Tim, para garantir um design elegante e responsivo.

---
### 📌 Objetivos
Utilizar o NGINX como servidor web dentro de um container.

Inserir uma página HTML/CSS estática moderna no container.

Criar uma imagem Docker personalizada com esses arquivos.

Publicar essa imagem no Docker Hub.

Executar o site localmente via Docker, acessando pelo navegador.

---
### 🛠️ Passo a Passo Completo
🔽 1. Clonar o repositório do Material Kit
```bash
git clone https://github.com/creativetimofficial/material-kit.git
cd material-kit

```
2. Acessar o diretório do projeto:
```bash
cd material-kit

```

🐳 Criar o Dockerfile
Crie um arquivo chamado Dockerfile dentro da pasta material-kit :
```bash

FROM nginx:latest

COPY . /usr/share/nginx/html/

EXPOSE 80

```
Explicação:

FROM nginx:latest: Estamos utilizando a imagem oficial do Nginx.

COPY . /usr/share/nginx/html/: Copia todos os arquivos da página estática (que estão no diretório atual) para o diretório onde o Nginx espera os arquivos HTML.

EXPOSE 80: O Nginx vai servir o site na porta 80.


### Construir a imagem Docker:

Agora, no diretório onde está o Dockerfile, execute o comando para construir a imagem:

```bash
docker build -t meu-site-nginx .

```
Esse comando irá construir a imagem usando o Dockerfile e nomeá-la como meu-site-nginx.


Passo 4: Rodar o Container com a Imagem Criada
Rodar o container com a imagem criada:

Após a imagem ser criada, você pode rodar o container. Use o seguinte comando:
```bash
docker run -d -p 8080:80 --name site-material-kit meu-site-nginx

```
Explicação:

-d: Rodar o container em segundo plano (modo desanexado).

-p 8080:80: Mapeia a porta 80 do container para a porta 8080 da sua máquina local.

--name site-material-kit: Define o nome do container como site-material-kit.

meu-site-nginx: Nome da imagem que você criou.

Verificar se o container está rodando:

Para verificar se o container foi iniciado corretamente, use:
```bash
docker ps

```


Passo 5: Acessar a Página no Navegador
Agora, a página estática está hospedada no seu container. Para acessá-la, abra o navegador e digite o endereço:
```bash
http://localhost:8080

```
Você verá a landing page do Creative Tim rodando no seu container Nginx.


Passo 6: Concluindo - Subindo a Imagem para o Docker Hub
Se você deseja compartilhar sua imagem Docker com outras pessoas ou armazená-la no Docker Hub, siga os seguintes passos:

Faça login no Docker Hub (se ainda não estiver logado):

```bash
docker login

```


2 Taguear a imagem para envio:

Antes de fazer o push da imagem para o Docker Hub, você precisa tagueá-la com seu nome de usuário do Docker Hub:
```bash
docker tag meu-site-nginx seu_usuario/meu-site-nginx:latest

```
Subir a imagem para o Docker Hub:

Agora, você pode enviar a imagem para o Docker Hub:
```bash
docker push seu_usuario/meu-site-nginx:latest

```


