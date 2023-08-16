package modelos

import (
	"errors"
	"time"
)

type Produto struct {
	ID       uint64   "json:id.omitempty"
	Title    string   "json:title,omitempty"
	Price    *float64 "json:price,omitempty"
	Quantity *uint64  "json:quantity,omitempty"
}

type Cart struct {
	ID       uint64    "json:id.omitempty"
	Quantity *uint64   "json:quantity,omitempty"
	CreatAt  time.Time "json:creatAt,omitempty"
}

func (cart *Cart) PrepararCart(etapa string) error {
	if erro := cart.validarCart(etapa); erro != nil {
		return erro
	}
	return nil
}

func (produto *Produto) Preparar(etapa string) error {
	if erro := produto.validar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (produto *Produto) validar(etapa string) error {
	if produto.Title == "" {
		return errors.New("O titulo é obrigatorio e nao pode estar em branco")
	}
	if produto.Price == nil {
		return errors.New("There was an error when trying to insert the product.The price is mandatory and cannot be blank")
	}
	if produto.Quantity == nil {
		return errors.New("There was an error when trying to insert the product.The quantity is mandatory and cannot be blank.")
	}
	if etapa == "atualizacao" {
		if produto.ID == 0 {
			return errors.New("Product with ID {productID} was not found.")
		}
	}
	if etapa == "cadastro" {
		return errors.New("The id is generated automatically")
	}
	return nil
}

func (cart *Cart) validarCart(etapa string) error {
	if cart.ID == 0 {
		return errors.New("Shopping cart with ID {cartID} was not found.")
	}
	if cart.Quantity == nil || *cart.Quantity == 0 || *cart.Quantity < 0 {
		return errors.New("One of the cart products does not have sufficient stock.")
	}
	return nil

}
