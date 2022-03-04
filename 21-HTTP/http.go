package main

import (
	"log"
	"net/http"
)

func home (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func usuarios (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuários!"))
}


func main() {
	// HTTP é um protocolo de comunicação - Base da comunicação WEB
	// Cliente (Faz a requisição - Request)   -   Servidor (processa requisição e envia resposta - Response)

	// Request - Response

	// Rotas 
	    // URI - Identificador do recurso 
			//Método- GET, POST, PUT, DELETE
			// GET - BUSCAR DADOS DO USUÁRIO
			//POST - CADASTRAR DADOS
			// PUT - ATUALIZAÇÕES DE DADOS
			// DELETE - DELETER DADOS

			http.HandleFunc("/home", home)

			http.HandleFunc("/usuarios", usuarios)

			log.Fatal(http.ListenAndServe(":5000", nil))




}