# 🚀 Guia Completo: Trabalhando com Docker e Segurança</h1>

## 📌 Menu 

- [🚀 Projeto: Site Estático com Docker + NGINX + Material Kit (Creative Tim)](#projeto-site-estatico-com-docker-nginx-material-kit-creative-tim)

- [🔐 Evitar Execução como Root em Containers](#evitar-execucao-como-root)



## 🚀 Projeto: Site Estático com Docker + NGINX + Material Kit (Creative Tim)
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


---

<h2 id="evitar-execucao-como-root">🔐 Evitar Execução como Root em Containers</h2>

## 🎯 Objetivo do Projeto
O principal objetivo deste projeto é demonstrar como criar um ambiente de execução Dockerizado seguro, onde a aplicação é rodando com um usuário não-root, evitando vulnerabilidades associadas ao uso de privilégios elevados dentro do container.

---

###  1️⃣  Criar o Diretório para o Projeto
Primeiro, crie um diretório onde você armazenará o seu código e o Dockerfile. No terminal, execute:

```bash
mkdir ~/meu-app-desf10
cd ~/meu-app-desf10

```
Este comando cria uma nova pasta chamada meu-app-desf10 e navega até ela.


### 2️⃣ Criar o Script Python
Dentro da pasta criada, crie o seu script Python, por exemplo, exe10.py. Você pode usar o nano para editar o arquivo:

```bash
nano exe10.py

```
Adicione o código do seu script Python. Exemplo de conteúdo:
```bash
                           
import os
from datetime import datetime

print("✅ Container rodando como userTeste!")
print(f"👤 UID/GID: {os.getuid()}/{os.getgid()}")
print(f"👤 whoami: {os.popen('whoami').read().strip()}")
print(f"👤 id: {os.popen('id').read().strip()}")

hora_execucao = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
print(f"⏰ Horário de execução: {hora_execucao}")

print("👋 Fim do script")

# Mantém o container rodando (Ctrl + C para parar)
while True:
    pass

```
]Salve o arquivo pressionando Ctrl + X, depois Y para confirmar e Enter para sair.


### 3️⃣ Criar o Dockerfile
Agora, crie o Dockerfile que será usado para construir a imagem Docker. Para isso, crie um novo arquivo chamado Dockerfile:
```bash

FROM python:3.8-slim

WORKDIR /app

COPY exe10.py /app/

RUN chown appuser:appuser /app/exe10.py

USER appuser

CMD ["python", "exe10.py"]

```


### 4️⃣ Construir a Imagem Docker
Agora, que você tem o Dockerfile pronto, é hora de construir a imagem Docker. No terminal, dentro da pasta onde o Dockerfile está, execute:


```bash
docker build -t desf10-img .

```
Este comando irá construir a imagem com o nome desf10-img a partir do Dockerfile.


### 5️⃣ Rodar o Container
Após a construção da imagem, execute o container a partir dessa imagem. Use o seguinte comando:
```bash
docker run -d --name
app-nao-root desf10-img
```

### 6️⃣ Verificar o Usuário Dentro do Container
Agora, você pode verificar se o container está sendo executado com o usuário não-root. Para isso, use o comando whoami dentro do container:

```bash
docker exec -it app-nao-root whoami
```

### 7️⃣ Rodar o Script no Container
Você também pode rodar o script manualmente dentro do container para verificar se está funcionando corretamente. 

```bash
docker exec -it app-nao-root python /app/exe10.py

```
<img src="https://github.com/user-attachments/assets/7558a041-65d1-41d0-9459-503978bc779f" alt="Imagem">

Esse comando exibirá qualquer saída ou erro gerado pelo container.

### Para ver se não é  usuario root 
```bash
whoami
```


---





