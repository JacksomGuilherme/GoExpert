package main

import (
	"client-server-api/database"
	"client-server-api/modelos"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"gorm.io/gorm"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/cotacao", BuscaCotacaoHandler)

	log.Fatalln(http.ListenAndServe(":8080", mux))
}

func BuscaCotacaoHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Millisecond*200)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		log.Println("Erro ao criar request: ", err)
		sendResponse(w, http.StatusBadRequest, modelos.Erro{Erro: err.Error()})
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Erro ao executar request: ", err)
		if errors.Is(err, context.DeadlineExceeded) {
			sendResponse(w, http.StatusRequestTimeout, modelos.Erro{Erro: err.Error()})
			return
		} else {
			sendResponse(w, http.StatusBadRequest, modelos.Erro{Erro: err.Error()})
			return
		}
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Erro ao ler corpo da resposta: ", err)
		sendResponse(w, http.StatusUnprocessableEntity, modelos.Erro{Erro: err.Error()})
		return
	}

	var usdBrl modelos.Usdbrl
	err = json.Unmarshal([]byte(body), &usdBrl)
	if err != nil {
		log.Println("Erro ao converter corpo de resposta para json: ", err)
		sendResponse(w, http.StatusUnprocessableEntity, modelos.Erro{Erro: err.Error()})
		return
	}

	db, err := database.Connect()
	if err != nil {
		log.Println("Erro ao conectar com banco de dados: ", err)
		sendResponse(w, http.StatusInternalServerError, modelos.Erro{Erro: err.Error()})
		return
	}

	err = insertCotacao(db, usdBrl.Cotacao)
	if err != nil {
		log.Println("Erro ao incluir cotação no banco: ", err)
		sendResponse(w, http.StatusRequestTimeout, modelos.Erro{Erro: err.Error()})
		return
	}

	cotacaoAtual, err := strconv.ParseFloat(usdBrl.Cotacao.Bid, 64)
	if err != nil {
		log.Println("Erro ao covnerter cotação: ", err)
		sendResponse(w, http.StatusRequestTimeout, modelos.Erro{Erro: err.Error()})
		return
	}
	sendResponse(w, http.StatusOK, cotacaoAtual)
}

func sendResponse(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

func insertCotacao(db *gorm.DB, cotacao modelos.Cotacao) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()

	err := db.WithContext(ctx).Create(&cotacao).Error
	if err != nil {
		return err
	}
	return nil
}
