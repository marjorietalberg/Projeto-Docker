# Projeto Docker e Containers 🚀

Este projeto é uma coletânea prática de exercícios com Docker, divididos em três níveis de dificuldade: Fácil, Médio e Difícil. Ele aborda desde os conceitos mais básicos, como criação e execução de containers, uso de imagens oficiais (como Nginx e Ubuntu), construção de imagens com Dockerfile e persistência de dados com volumes, até tópicos mais avançados, como redes Docker, Docker Compose, multi-stage builds, segurança de containers (usuário não-root) e análise de vulnerabilidades com a ferramenta Trivy. É ideal para quem deseja aprender Docker na prática, explorando casos reais de desenvolvimento e boas práticas de containerização.

<img src="https://github.com/user-attachments/assets/ce760706-ab2f-4e2a-88c5-26b4fdbb988b" alt="Docker Project Banner">

---

## Menu de Navegação
- 📌 Etapas do Projeto

- [Etapa Fácil](./Projeto%20Docker%20:%20Etapa%20Fácil.md): Introdução ao Docker e comandos básicos.
- [Etapa Médio](./Projeto%20Docker%20:%20Etapa%20Médio.md): Containers com Flask, MySQL e redes.
- [Etapa Difícil](./Projeto%20Docker%20:%20Etapa%20Difícil.md): Multi-stage builds e integração de serviços.

- 📌 [Links Úteis](#-links-úteis)
- 📌 [Conclusão](#-conclusão)

---

## 🛠️ Ferramentas Usadas

Durante o desenvolvimento, utilizei algumas ferramentas incríveis. Abaixo estão as ferramentas e tecnologias que usei, com seus respectivos ícones.

| Ferramenta   | Ícone                   | Descrição |
|--------------|-------------------------|-----------|
| **Docker**   | ![Docker](https://img.icons8.com/?size=100&id=kUIdznZvxAud&format=png&color=000000) | Ferramenta principal para a criação e gerenciamento de containers. |
| **Python**   | ![Python](https://img.icons8.com/?size=100&id=12584&format=png&color=000000) | Linguagem de programação usada para criar aplicações como Flask. |
| **Node.js**  | ![Node.js](https://img.icons8.com/?size=100&id=FQlr_bFSqEdG&format=png&color=000000) | Plataforma JavaScript usada para criar aplicações back-end. |
| **Go (Golang)** | ![Go](https://img.icons8.com/?size=100&id=10710&format=png&color=000000) | Linguagem de programação usada para otimizar uma aplicação com multi-stage builds. |
| **MySQL**    | ![MySQL](https://img.icons8.com/?size=100&id=11572&format=png&color=000000) | Banco de dados relacional usado no projeto. |
| **PostgreSQL** | ![PostgreSQL](https://img.icons8.com/?size=100&id=25010&format=png&color=000000) | Banco de dados utilizado para um dos containers no projeto. |
| **TailwindCSS** | ![TailwindCSS](https://img.icons8.com/?size=100&id=qOFWMoaAQIdR&format=png&color=000000) | Framework CSS para criar interfaces de usuário modernas e responsivas. |

---

# Etapas do Projeto :

## 1️⃣ Parte 1: Containerização Básica

### 🚀 Rodando um Container Básico
Na primeira parte, criei um container usando a imagem oficial do **Nginx** e configurei uma página estática usando **TailwindCSS**. O objetivo foi entender o processo básico de criação e execução de containers.

#### 🖥️ Criando e Rodando um Container Interativo
Aqui, iniciei um container com a imagem **Ubuntu** e interagi com o terminal. Testei comandos básicos e instalei pacotes, o que me ajudou a aprender como navegar e gerenciar containers interativos.

#### 📋 Listando e Removendo Containers
Aprendi como listar containers em execução, parar e remover containers. Essas operações são essenciais para um gerenciamento eficiente.


---

## 2️⃣ Parte 2: Aplicações e Banco de Dados

#### 🐍 Criando um Dockerfile para uma Aplicação Python
Aqui, criei um **Dockerfile** para rodar uma aplicação **Flask** dentro de um container. A aplicação exibe uma mensagem simples no navegador, e o Dockerfile prepara o ambiente necessário para rodá-la.

#### 💾 Criando e Utilizando Volumes para Persistência de Dados
Para garantir que os dados não se perdessem quando o container fosse reiniciado, configurei um **volume Docker** para o banco de dados **MySQL**. Isso foi essencial para garantir a persistência dos dados.

#### 🚀 Criando e Rodando um Container Multi-stage
Essa parte do projeto envolveu a criação de uma aplicação **Go** (Golang) com **multi-stage builds**. Isso ajudou a otimizar o tamanho das imagens Docker, o que é uma prática essencial em ambientes de produção.

---

## 3️⃣ Parte 3: Otimização e Segurança

#### 🌐 Construindo uma Rede Docker para Comunicação entre Containers
Para criar ambientes mais complexos, construí uma rede Docker personalizada, permitindo que containers **Node.js** e **MongoDB** se comunicassem. Isso foi importante para entender a comunicação entre containers em microservices.

#### 📦 Criando um Compose File para Rodar uma Aplicação com Banco de Dados
Aqui, usei o **Docker Compose** para orquestrar múltiplos containers de forma simples. Configurei uma aplicação com **PostgreSQL**, facilitando a criação de ambientes complexos com várias dependências.

#### 🌐 Criando uma Imagem Personalizada com um Servidor Web e Arquivos Estáticos
Criei uma imagem personalizada com o **Nginx**, configurado para servir uma página estática HTML/CSS. Isso me ajudou a entender como personalizar e otimizar imagens para produção.

#### 🔒 Evitar Execução como Root
A segurança foi uma parte essencial do projeto. Criamos um Dockerfile para rodar uma aplicação com um usuário **não-root**, garantindo um ambiente mais seguro.

#### 🛡️ Analisando Imagem Docker com Trivy
Para fechar o projeto, utilizei a ferramenta **Trivy** para escanear imagens Docker em busca de vulnerabilidades. Isso me ajudou a garantir que as imagens que estou utilizando são seguras e não possuem falhas críticas.

---

## 🔗 Links Úteis

Durante o desenvolvimento, aprendi muito sobre as ferramentas que usei. Aqui estão alguns links para quem quiser explorar mais sobre elas:

- [Docker](https://www.docker.com/)
- [Python](https://www.python.org/)
- [Node.js](https://nodejs.org/en/)
- [Go (Golang)](https://golang.org/)
- [MySQL](https://www.mysql.com/)
- [PostgreSQL](https://www.postgresql.org/)
- [TailwindCSS](https://tailwindcss.com/)

---

## 💬 Conclusão

Esse projeto foi uma excelente oportunidade para aprender sobre Docker, containers e práticas de segurança em ambientes de produção. A flexibilidade e poder do Docker permitiram que eu experimentasse com diferentes tecnologias e ferramentas, criando um ambiente de desenvolvimento isolado e eficiente.

---

**Espero que você tenha gostado de conhecer o projeto!** 😄 Fique à vontade para deixar sugestões ou perguntas nos comentários! 👇
