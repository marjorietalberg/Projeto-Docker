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
🔹 1. Clone o Repositório GitHub Localmente
Com o repositório criado no GitHub, você pode cloná-lo no seu sistema local para começar a trabalhar. Execute o seguinte comando no seu terminal (substitua seu-usuario pelo seu nome de usuário no GitHub):
```bash
git clone https://github.com/seu-usuario/landingpage-nginx.git
cd landingpage-nginx

```
Esse comando vai clonar o repositório no seu diretório local e entrar na pasta do projeto

🔹 2. Clone o Repositório do Material Kit
Agora, clone o repositório do Material Kit da Creative Tim:
Esse comando vai criar a pasta material-kit com todos os arquivos do template.

```bash
git clone https://github.com/creativetimofficial/material-kit.git

```
🔹 4. Copie os Arquivos do Material Kit para o Repositório Local
Agora, vamos copiar os arquivos do Material Kit para o diretório landingpage-nginx no seu repositório local.
```bash
cp -r material-kit/* landingpage-nginx/

```

🔹 5. Crie o Dockerfile
Dentro do diretório landingpage-nginx, crie um arquivo chamado Dockerfile para configurar o servidor Nginx.

Execute:


```bash
nano Dockerfile

```
No editor de texto, cole o seguinte conteúdo:
```bash

FROM nginx:alpine


RUN rm -rf /usr/share/nginx/html/*


COPY . /usr/share/nginx/html/


EXPOSE 80


CMD ["nginx", "-g", "daemon off;"]

```
Salve e saia (no nano, pressione CTRL + X, depois Y e Enter).

🔹 6. (Opcional) Crie o Arquivo .dockerignore
Para evitar copiar arquivos desnecessários para o container, crie um arquivo .dockerignore:
```bash
nano .dockerignore
```
Adicione o seguinte conteúdo:
```bash
node_modules
*.git
*.zip

```
🔹 9. Construa a Imagem Docker Localmente
Agora, vamos construir a imagem Docker a partir do Dockerfile criado.

Execute o comando abaixo no diretório onde está o Dockerfile (o diretório landingpage-nginx):
```bash
docker build -t landingpage-nginx .

```

Esse comando vai criar a imagem Docker com o nome landingpage-nginx.



🔹 10. Execute o Container Docker Localmente
Agora que você tem a imagem Docker, podemos rodar o container com o seguinte comando:

```bash
docker run -d -p 8080:80 --name site-landingpage landingpage-nginx

```
Esse comando vai:

Iniciar o container em segundo plano (-d).

Mapear a porta 8080 do host para a porta 80 do container (onde o Nginx está ouvindo).

Nomear o container como site-landingpage.


🔹 11. Acesse a Landing Page no Navegador
Agora, abra o navegador e acesse:
```bash
http://localhost:8080

```
 <img src="https://github.com/user-attachments/assets/4754eefc-c6b5-4e4c-baf1-c96ba5c53de2" alt="Exemplo de Imagem">

🔹 12. (Opcional) Subir a Imagem Docker para o Docker Hub
Se você deseja subir a imagem Docker para o Docker Hub para compartilhar ou reutilizar, siga os seguintes passos:

Faça login no Docker Hub:
```bash
docker login

```

Tag a imagem Docker com o nome do seu usuário no Docker Hub (substitua seu-usuario pelo seu nome de usuário no Docker Hub):

```bash
docker tag landingpage-nginx seu-usuario/landingpage-nginx

```

Suba a imagem para o Docker Hub:
```bash
docker push seu-usuario/landingpage-nginx

```
