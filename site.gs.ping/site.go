
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `
        <!DOCTYPE html>
        <html lang="pt-br">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>GS PING - Monitoramento de Precisão</title>
            <link href="https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;700&display=swap" rel="stylesheet">
            <link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.4/css/all.min.css" rel="stylesheet">
            <style>
                body {
                    font-family: 'Poppins', sans-serif;
                    margin: 0;
                    padding: 0;
                    background-color: #1a1a1a;
                    color: white;
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    height: 100vh;
                    text-align: center;
                }
                .container {
                    max-width: 600px;
                    padding: 30px;
                    border: 2px solid #00b894;
                    border-radius: 10px;
                    background-color: #333;
                }
                .icon {
                    font-size: 80px;
                    color: #00b894;
                    margin-bottom: 20px;
                }
                .title {
                    font-size: 36px;
                    font-weight: 700;
                    margin-bottom: 15px;
                    color: #00b894;
                }
                .description {
                    font-size: 18px;
                    color: #ccc;
                    margin-bottom: 30px;
                }
                .footer {
                    margin-top: 20px;
                    font-size: 14px;
                    color: #888;
                }
                .footer p {
                    margin: 0;
                }
                .footer a {
                    color: #00b894;
                    text-decoration: none;
                }
            </style>
        </head>
        <body>
            <div class="container">
                <div class="icon">
                    <i class="fas fa-bullseye"></i> <!-- Ícone de alvo -->
                </div>
                <div class="title">
                    GS PING - A Precisão no Monitoramento
                </div>
                <div class="description">
                    Com GS PING, seu sistema é monitorado em tempo real, garantindo performance e estabilidade 24/7.
                </div>
                <div class="footer">
                    <p>© 2025 GS PING | <a href="https://www.linkedin.com/in/marjorie-pedroso-talberg-89112a35a" target="_blank">LinkedIn</a></p>
                </div>
            </div>
        </body>
        </html>
    `)
}

func main() {
    http.HandleFunc("/", handler) // Mapeia a URL "/"
    fmt.Println("Servidor rodando na porta 8080...")
    err := http.ListenAndServe(":8080", nil) // Inicia o servidor na porta 8080
    if err != nil {
        fmt.Println("Erro ao iniciar o servidor:", err)
    }
}
