# 🚀 Guia Completo: Trabalhando com Docker e Segurança</h1>

## 📌 Menu 

- [🚀 Projeto: Site Estático com Docker + NGINX + Material Kit (Creative Tim)](#projeto-site-estatico-com-docker-nginx-material-kit-creative-tim)

- [🔐 Evitar Execução como Root em Containers](#evitar-execucao-como-root)

- [🚀Comparando Imagens Docker](#comparando-imagens-docker)


## 🚀 Projeto 9: Site Estático com Docker + NGINX + Material Kit (Creative Tim)
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

# Projeto 10 :🔐 Evitar Execução como Root em Containers

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
```
<img src="https://github.com/user-attachments/assets/ceb2b6a1-1c23-4195-9a1c-a3e3581f66cc" alt="Image" />

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


---

# Projeto 12: 🚀 CORRIGINDO VULNERABILIDADES DO DOCKERFILE 🐳 

### 🔹 1. Criar pasta do projeto
```bash
mkdir exe12
```
Entrar na pasta 
```bash
cd exe12
```

🔹 2. Criar o arquivo app.py
```bash
nano app.py
```
```bash
from flask import Flask
import requests

app = Flask(__name__)

@app.route("/")
def home():
    return "Olá, Seja bem vindo ! "

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)
```
🔹 3. Criar o requirements.txt
```bash
flask==1.1.1
requests==2.22.0

```
🔹 4. Criar o Dockerfile 
```bash
nano Dockerfile

```
```bash

FROM python:3.9-slim

RUN apt-get update && apt-get upgrade -y && apt-get dist-upgrade -y && apt-get clean

RUN apt-get install -y --no-install-recommends \
    gcc \
    build-essential \
    libpq-dev \
    perl-base \
    zlib1g \
    && rm -rf /var/lib/apt/lists/*

RUN pip install --no-cache-dir --upgrade setuptools==70.0.0

RUN groupadd -r grupoTeste && useradd -r -g grupoTeste userTeste

WORKDIR /app

COPY requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt

COPY . .

RUN chown -R userTeste:grupoTeste /app

USER userTeste

EXPOSE 5000

CMD ["python", "app.py"]
```

### 🔹 4. Verifique se os arquivos foram criados corretamente
```bash
ls
```
Suba o container
```bash
docker build -t exe12-image .
docker run -d --name exe12-container -p 5000:5000 exe12-image
```
<img src="https://github.com/user-attachments/assets/11f59f36-b879-469e-b7d6-e3740fb11991" alt="Image">

### Acesse pelo navegador
```bash 
http://localhost:5000
```
<img src="https://github.com/user-attachments/assets/b756a252-8c1e-427e-ad6d-3e4574682e62" alt="Image">

### Para ver se não é não é o root
```bash

docker exec -it exe12-container sh
```
```bash
whoami
```
---

## 📊 Comparando Imagens Docker
Compare a imagem usada anteriormente com a nova que foi desenvolvida
a. Gere o relatório da imagem anterior
```bash
trivy image --severity HIGH,CRITICAL --format json python:3.9 > resultado.json
```
```bash
jq -r '
  "| Pacote | Versão atual | Correção disponível | Severidade | Ação sugerida |",
  "|--------|--------------|---------------------|------------|---------------|",
  (.Results[] | select(.Vulnerabilities != null) | .Vulnerabilities[] |
    "| \(.PkgName) | \(.InstalledVersion) | \(.FixedVersion // "❌ Não") | \(.Severity) | \(
      if .FixedVersion then
        "Atualizar para a versão \(.FixedVersion)"
      else
        "Considerar substituição ou monitorar atualizações futuras"
      end
    ) |"
  )
' resultado.json | uniq | column -t -s '|' > relatorio1.md
```
b. Gere o relatório da nova imagem
```bash
trivy image --severity HIGH,CRITICAL --format json exe12-image > resultado.json
```

```bash
jq -r 'jq -r '
  "| Pacote | Versão atual | Correção disponível | Severidade | Ação sugerida |",
  "|--------|--------------|---------------------|------------|---------------|",
  (.Results[] | select(.Vulnerabilities != null) | .Vulnerabilities[] |
    "| \(.PkgName) | \(.InstalledVersion) | \(.FixedVersion // "❌ Não") | \(.Severity) | \(
      if .FixedVersion then
        "Atualizar para a versão \(.FixedVersion)"
      else
        "Considerar substituição ou monitorar atualizações futuras"
      end
    ) |"
  )
' resultado2.json | uniq | column -t -s '|' > relatorio2.md
  "| Pacote | Versão atual | Correção disponível | Severidade | Ação sugerida |",
  "|--------|--------------|---------------------|------------|---------------|",
  (.Results[] | select(.Vulnerabilities != null) | .Vulnerabilities[] |
    "| \(.PkgName) | \(.InstalledVersion) | \(.FixedVersion // "❌ Não") | \(.Severity) | \(
      if .FixedVersion then
        "Atualizar para a versão \(.FixedVersion)"
      else
        "Considerar substituição ou monitorar atualizações futuras"
      end
    ) |"
  )
' resultado2.json | uniq | column -t -s '|' > relatorioant.md
```
 ### Veja o relatório da imagem anterior
 ```bash
cat relatorio1.md
```
```bash
echo $(( $(wc -l < relatorio.md) - 2 ))
```
```bash

echo $(( $(wc -l < relatorioant.md) - 2 ))
```

<img src="https://github.com/user-attachments/assets/e700fdcf-1416-46af-8ad8-cdfce27e4ccf" alt="Image">

---
### ✅ Conclusão
Neste exercício, foi possível identificar e corrigir diversas vulnerabilidades presentes em um Dockerfile com más práticas. A imagem original utilizava uma base genérica (python:3.9), o que aumentava consideravelmente o tamanho da imagem e trazia dependências desnecessárias e potencialmente inseguras. Além disso, a execução da aplicação ocorria com privilégios de root, o que representa um sério risco de segurança.

Por meio da aplicação de boas práticas, como a adoção da imagem python:3.9-slim, a utilização de builds em múltiplas etapas (multi-stage), a execução da aplicação com um usuário não-root e o uso do pip install com --no-cache-dir, foi possível gerar uma imagem muito mais enxuta, segura e alinhada com os padrões modernos de DevSecOps.

A comparação entre as versões demonstra a importância de otimizar e proteger as imagens Docker desde a base, reforçando que segurança não deve ser um passo adicional, mas um componente essencial desde o início do ciclo de desenvolvimento.
