 # 🚀 Desafios de Docker: Começando com Containers e Flask! 🐳

#### Este repositório contém uma série de desafios práticos para quem está aprendendo sobre Docker. Aqui, vamos explorar desde a criação de containers simples até a construção de aplicações em Python com Docker. Acompanhe o passo a passo de cada exercício abaixo! Vamos colocar a mão na massa e aprender como criar e gerenciar containers Docker de forma prática e divertida! 😄
---
## 🎯 Objetivo:
Executar um container NGINX que sirva uma página HTML com TailwindCSS (um framework moderno de CSS). Ao acessar localhost:8080, o navegador mostrará a landing page.


### 1️⃣ Rodando um Container Básico com Nginx e TailwindCSS 🌐
Descrição
Neste desafio, vamos rodar um container Docker usando a imagem oficial do Nginx e configurar uma landing page com TailwindCSS. O objetivo é servir uma página estática através do Nginx, estilizada com o TailwindCSS.

### 🔹 Passo 1: Criando o Diretório do Projeto
Primeiro, crie um diretório onde armazenaremos os arquivos do projeto:
```bash
mkdir ~/docker-exercicios/01-nginx-tailwind
cd ~/docker-exercicios/01-nginx-tailwind
```
---
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
###🔹 Passo 3:
Rode o container com volume
```bash
docker run -d --name nginx-tailwind -p 8080:80 -v $(pwd)/index.html:/usr/share/nginx/html/index.html nginx
```
📌 Explicação:

-d: roda o container em segundo plano (modo detached)
-name: dá um nome para o container
-p 8080:80: mapeia a porta local 8080 para a porta 80 do container
-v: monta o arquivo local no caminho correto do NGINX

### 🔹Passo 4:
Acesse no navegador:
```bash
http://localhost:8080
```
<img src="https://github.com/user-attachments/assets/3daabff8-3564-4440-892c-a98706d159bf" alt="Imagem 1">
<img src="https://github.com/user-attachments/assets/a1161ad7-3170-48d0-b395-79a54391e268" alt="Imagem 2">
<img src="https://github.com/user-attachments/assets/b019fcc1-3cb2-4b39-9b41-7a1bfa79764f" alt="Imagem 3">










### 🔹 Passo 5: Criar o arquivo principal app.py
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


