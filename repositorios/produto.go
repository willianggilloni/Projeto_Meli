package repositorios

import (
	"database/sql"
	"testMeli/src/modelos"
)

type Produtos struct {
	db *sql.DB
}

type Carts struct {
	db *sql.DB
}

func NovoRepositorioDeProdutos(db *sql.DB) *Produtos {
	return &Produtos{db}
}

func NovoRepositorioDeCarts(db *sql.DB) *Carts {
	return &Carts{db}
}

func (repositorio Produtos) Criar(produtos modelos.Produto) (uint64, error) {
	statement, erro := repositorio.db.Prepare("insert into tb_testMeli(Title,Price,Quantity) values (?, ?, ?)")

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(produtos.Title, produtos.Price, produtos.Quantity)
	if erro != nil {
		return 0, erro
	}
	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

func (repositorio Produtos) Buscar(title string) ([]modelos.Produto, error) {
	var query = "select id, Title, Price, Quantity from tb_testMeli"
	var parametro = []any{}
	if title != "" {
		query = query + " where Title like ?"
		parametro = append(parametro, "%"+title+"%")

	}
	linhas, erro := repositorio.db.Query(query, parametro...)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var produtos []modelos.Produto

	for linhas.Next() {
		var produto modelos.Produto

		if erro := linhas.Scan(&produto.ID, &produto.Title, &produto.Price, &produto.Quantity); erro != nil {
			return nil, erro
		}
		produtos = append(produtos, produto)
	}
	return produtos, nil
}

func (repositorio Produtos) BuscarPorID(ID uint64) (modelos.Produto, error) {
	linhas, erro := repositorio.db.Query(
		"select id, Title, Price, Quantity from tb_testMeli where id = ?",
		ID,
	)
	if erro != nil {
		return modelos.Produto{}, erro
		defer linhas.Close()
	}
	var produto modelos.Produto

	if linhas.Next() {
		if erro = linhas.Scan(&produto.ID, &produto.Title, &produto.Price, &produto.Quantity); erro != nil {
			return modelos.Produto{}, erro
		}
	}

	return produto, nil
}

func (repositorio Produtos) Atualizar(ID uint64, produto modelos.Produto) error {
	statement, erro := repositorio.db.Prepare("update tb_testMeli set Title = ?, Price = ?, Quantity = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(produto.Title, produto.Price, produto.Quantity, ID); erro != nil {
		return erro
	}
	return nil
}

func (repositorio Produtos) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from tb_testMeli where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}
	return nil
}

func (repositorio Carts) CriarCart(carts modelos.Cart) (uint64, error) {
	statement, erro := repositorio.db.Prepare("insert into tb_Carts(Quantity,CreatAt) values (?,?)")

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(carts.Quantity, carts.CreatAt)
	if erro != nil {
		return 0, erro
	}
	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

func (repositorio Carts) BuscarCartPorID(ID uint64) (modelos.Cart, error) {
	linhas, erro := repositorio.db.Query(
		"select id, Quantity,CreatAt from tb_Carts where id = ?",
		ID,
	)
	if erro != nil {
		return modelos.Cart{}, erro
		defer linhas.Close()
	}
	var cart modelos.Cart

	if linhas.Next() {
		if erro = linhas.Scan(&cart.ID, &cart.Quantity, &cart.CreatAt); erro != nil {
			return modelos.Cart{}, erro
		}
	}

	return cart, nil
}
