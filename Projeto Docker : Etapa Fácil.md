 # 🚀 Desafios de Docker: Começando com Containers e Flask! 🐳

#### Este repositório contém uma série de desafios práticos para quem está aprendendo sobre Docker. Aqui, vamos explorar desde a criação de containers simples até a construção de aplicações em Python com Docker. Acompanhe o passo a passo de cada exercício abaixo! Vamos colocar a mão na massa e aprender como criar e gerenciar containers Docker de forma prática e divertida! 😄
---
## 🎯 Objetivo:
Executar um container NGINX que sirva uma página HTML com TailwindCSS (um framework moderno de CSS). Ao acessar localhost:8080, o navegador mostrará a landing page.

📘 Exercício	Descrição

🔗 01 - Container Nginx + Tailwind	Executa um container Nginx com site estático usando TailwindCSS.

[Ir para seção 1](#1-rodando-um-container-básico-com-nginx-e-tailwindcss-)

🔗 02 - Ubuntu Interativo	Executa um container Ubuntu de forma interativa, com scripts Bash.

[Ir para seção 2](#2-criando-e-rodando-um-container-interativo)


🔗 03 - Gerenciando Containers	Lista, para e remove containers usando comandos Docker.

3️⃣ Listando e removendo containers


🔗 04 - Flask com Dockerfile	Cria uma imagem para uma aplicação Flask simples.

[Ir para seção 4](#4-criando-um-dockerfile-com-flask-python)


---

### 📌 Você pode usar a imagem oficial do Nginx e montar um volume com sua landing page.

Passos:
a. Baixe a landing page do Tailwind:
```bash
curl -L https://github.com/tailwindlabs/tailwindcss-landing-page/archive/refs/heads/main.zip -o landing.zip
unzip landing.zip
cd tailwindcss-landing-page-main
```
b. Rode o container Nginx servindo essa pasta:
```bash
docker run --name nginx-tailwind -p 8080:80 -v $(pwd):/usr/share/nginx/html:ro -d nginx
```

---

### 1️⃣ Rodando um Container Básico com Nginx e TailwindCSS 🌐
Descrição
Neste desafio, vamos rodar um container Docker usando a imagem oficial do Nginx e configurar uma landing page com TailwindCSS. O objetivo é servir uma página estática através do Nginx, estilizada com o TailwindCSS.

### 🔹 Passo 1: Criando o Diretório do Projeto
Primeiro, crie um diretório onde armazenaremos os arquivos do projeto:
```bash
mkdir ~/docker-exercicios/01-nginx-tailwind
cd ~/docker-exercicios/01-nginx-tailwind
```
### 🔹 Passo 2:
Crie o arquivo index.html
```bash
nano index.html
```
```bash                                                                            
<!DOCTYPE html>
<html lang="pt-br">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Landing Page Criativa e Animada</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <style>
        /* Animações */
        @keyframes slideIn {
            0% {
                transform: translateX(-100%);
                opacity: 0;
            }
            100% {
                transform: translateX(0);
                opacity: 1;
            }
        }

        .slide-in {
            animation: slideIn 1s ease-out forwards;
        }

        @keyframes fadeIn {
            0% {
                opacity: 0;
            }

            100% {
                opacity: 1;
            }
        }

        .fade-in {
            animation: fadeIn 2s ease-out forwards;
        }
         /* Hover effect */
        .hover\:scale-105:hover {
            transform: scale(1.05);
        }
    </style>
    </head>

<body class="bg-gradient-to-r from-purple-500 to-blue-500 text-white font-sans">
    <header class="text-center py-20 slide-in">
        <h1 class="text-5xl font-bold mb-4 animate__animated animate__fadeIn">Bem-vi>
        <p class="text-xl mb-6 animate__animated animate__fadeIn animate__delay-1s">>
        <a href="#about" class="bg-white text-purple-500 py-3 px-6 rounded-full text>
    </header>

     <section id="about" class="py-20 bg-white text-gray-900 fade-in">
        <div class="max-w-4xl mx-auto text-center">
            <h2 class="text-4xl font-semibold mb-6">Sobre o Projeto</h2>
            <p class="text-xl mb-4">Este projeto foi feito para demonstrar como usar>
            <p class="text-lg">Com Tailwind, você pode estilizar rapidamente seu sit>
        </div>
          </div>
    </section>
  
    <section class="py-20">
        <div class="max-w-6xl mx-auto grid grid-cols-1 md:grid-cols-3 gap-12">
            <div class="text-center bg-white p-6 rounded-lg shadow-xl transform hove>
                <h3 class="text-2xl font-semibold mb-4">Design Responsivo</h3>
                <p class="text-gray-700">O layout se adapta a qualquer tamanho de te>
            </div>
            <div class="text-center bg-white p-6 rounded-lg shadow-xl transform hove>
                <h3 class="text-2xl font-semibold mb-4">Desempenho Otimizado</h3>
                <p class="text-gray-700">TailwindCSS é uma biblioteca leve que ofere>
            </div>
            <div class="text-center bg-white p-6 rounded-lg shadow-xl transform hove>
                <h3 class="text-2xl font-semibold mb-4">Personalização Fácil</h3>
                <p class="text-gray-700">Com Tailwind, você pode personalizar o desi>
            </div>
        </div>
    </section>

  
  <footer class="bg-gray-900 text-white py-10 text-center">
        <p>&copy; 2025 Meu Site Criativo. Todos os direitos reservados.</p>
    </footer>

</body>
</html>
```
---
### 🔹 Passo 3:
Rode o container com volume

```bash
docker run -d --name nginx-tailwind -p 8080:80 -v $(pwd)/index.html:/usr/share/nginx/html/index.html nginx
```
### 📌 Explicação:

-d: roda o container em segundo plano (modo detached)

-name: dá um nome para o container

-p 8080:80: mapeia a porta local 8080 para a porta 80 do container

-v: monta o arquivo local no caminho correto do NGINX

---

### 🔹Passo 4:
Acesse no navegador:
```bash
http://localhost:8080
```
<img src="https://github.com/user-attachments/assets/3daabff8-3564-4440-892c-a98706d159bf" alt="Imagem 1">
<img src="https://github.com/user-attachments/assets/a1161ad7-3170-48d0-b395-79a54391e268" alt="Imagem 2">
<img src="https://github.com/user-attachments/assets/b019fcc1-3cb2-4b39-9b41-7a1bfa79764f" alt="Imagem 3">

---

## 2️⃣ Criando e rodando um container interativo

🎯 Objetivo:
Rodar um container Ubuntu e usar seu terminal para instalar pacotes, visualizar logs etc.
### 🔹 Passo 1:
Execute um container interativo:
```bash 
docker run -it ubuntu bash
```
### 📌 Explicação:
-it: interativo + pseudo-terminal (permite usar o terminal dentro do container)

ubuntu: imagem oficial do Ubuntu

bash: executa o bash dentro do container

### 🔹 Passo 2:
Dentro do container, atualize o sistema e instale pacotes:
```bash
apt update
apt install -y curl
```
### 🔹Passo 3:
Testar os comandos:
```bash 
dmesg | tail     # Exibe os últimos logs do sistema
curl ifconfig.me # Exibe o IP do container
```
### 🔹 Passo 4 :Saia do container
Para sair do container use :
```bash
exit
```
---

### 3️⃣ Listando e removendo containers
a. Lista todos os containers (rodando e parados):
```bash
docker ps -a
```
b. Parar um container:
```bash
docker stop ubuntu-test
```
c. Remover um container:
```bash
docker rm ubuntu-test
```
---

 ### 4️⃣ Criando um Dockerfile com Flask (Python)
🎯 Objetivo:
Criar uma aplicação Flask simples que roda em um container. Ao acessar localhost:5000, ela responde com uma mensagem.

🔹 Passo 1:
 Crie a pasta do projeto:
```bash
mkdir 04-flask-app
cd 04-flask-app
```

### 🔹 Passo 2: Criar o arquivo principal app.py
Agora vamos criar o arquivo Python que representa a aplicação Flask:

```bash 
nano app.py
```
```bash
from flask import Flask

app = Flask(__name__)

@app.route('/')
def hello_world():
    return 'Olá, mundo do Flask no Docker!'

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
```

<img src="https://github.com/user-attachments/assets/40edba28-de38-4e4c-ae70-2b475a4ce742" alt="Image" />
---

 ### 🔹Passo 3: Crie o arquivo requirements.txt:
```bash
echo flask > requirements.txt
```
 ### 🔹Passo 4: Crie o Dockerfile:
```bash
nano Dockerfile
```

```bash
FROM python:3.8-slim
WORKDIR /app
COPY requirements.txt .
RUN pip install -r requirements.txt
COPY . .
EXPOSE 5000
CMD ["python", "app.py"]
```
### 📌 Explicação:
. FROM: imagem base do Python

. WORKDIR: diretório de trabalho no container

. COPY: copia os arquivos para o container

. RUN: instala dependências

. EXPOSE: expõe a porta 5000

. CMD: executa o script Python


 ### 🔹 Passo 5: Build da imagem:
```bash
docker build -t flask-app .
```

 ### 🔹 Passo 6: Execute o container:
```bash
docker run -d -p 5000:5000 flask-app
```

 ### 🔹 Passo 7 :Acesse no navegador:
```bash
http://localhost:5000
```
---

<img src="https://github.com/user-attachments/assets/4717b30d-2e45-4289-a779-4da5b0c93e47" alt="Imagem">


