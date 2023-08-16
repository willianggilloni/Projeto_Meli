package rotas

import (
	"net/http"
	"testMeli/src/controllers"
)

var rotasProducts = []Rotas{
	{
		URI:    "/products",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarProducts,
	},
	{
		URI:    "/products",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarProducts,
	},
	{
		URI:    "/products/{productId}",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarProductsPorID,
	},
	{
		URI:    "/products/{productId}",
		Metodo: http.MethodPut,
		Funcao: controllers.AtualizarProducts,
	},
	{
		URI:    "/products/{productId}",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeletarProducts,
	},

	{
		URI:    "/products/carts",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarCarts,
	},
	{
		URI:    "/products/carts/{cartId}",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarCarts,
	},
}
