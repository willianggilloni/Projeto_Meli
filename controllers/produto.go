package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"testMeli/src/banco"
	"testMeli/src/modelos"
	"testMeli/src/repositorios"
	"testMeli/src/respostas"
)

type Produtos struct {
	db *sql.DB
}

func NovoRepositorioDeProdutos(db *sql.DB) *Produtos {
	return &Produtos{db}
}

func CriarProducts(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var produto modelos.Produto
	if erro = json.Unmarshal(corpoRequest, &produto); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()
	repositorio := repositorios.NovoRepositorioDeProdutos(db)
	produto.ID, erro = repositorio.Criar(produto)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, produto)
}

func BuscarProducts(w http.ResponseWriter, r *http.Request) {
	title := strings.ToLower(r.URL.Query().Get("title"))
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeProdutos(db)
	produtos, erro := repositorio.Buscar(title)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, produtos)
}

func BuscarProductsPorID(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	produtoID, erro := strconv.ParseUint(parametros["productId"], 10, 64)
	//if erro != nil {
	//	respostas.Erro(w, http.StatusBadRequest, erro)
	//	return
	//}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeProdutos(db)
	produto, erro := repositorio.BuscarPorID(produtoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, produto)
}
func AtualizarProducts(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)
	produtoID, erro := strconv.ParseUint(parametro["productId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	var produto modelos.Produto
	if erro = json.Unmarshal(corpoRequisicao, &produto); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = produto.Preparar("edicao"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeProdutos(db)
	erro = repositorio.Atualizar(produtoID, produto)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func DeletarProducts(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	produtoID, erro := strconv.ParseUint(parametros["productId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeProdutos(db)
	if erro = repositorio.Deletar(produtoID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func CriarCarts(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var cart modelos.Cart
	if erro = json.Unmarshal(corpoRequest, &cart); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()
	repositorio := repositorios.NovoRepositorioDeCarts(db)
	cart.ID, erro = repositorio.CriarCart(cart)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, cart)
}

func BuscarCarts(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	cartID, erro := strconv.ParseUint(parametros["cartId"], 10, 64)
	//if erro != nil {
	//	respostas.Erro(w, http.StatusBadRequest, erro)
	//	return
	//}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeCarts(db)
	cart, erro := repositorio.BuscarCartPorID(cartID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, cart)
}
