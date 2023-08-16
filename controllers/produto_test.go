package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testMeli/src/modelos"
	"testing"
)

func TestCriarProducts(t *testing.T) {
	payload := []byte(`{"title": "Test Product", "price": 10.99, "quantity": 100, "stock": 100}`)

	req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CriarProducts)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	var responseProduct modelos.Produto
	err = json.Unmarshal(rr.Body.Bytes(), &responseProduct)
	if err != nil {
		t.Errorf("Error unmarshaling response body: %v", err)
	}

	// Add assertions here to validate the responseProduct and other expectations
}

// Similar testing functions can be created for other controller functions
func TestBuscarCarts(t *testing.T) {
	req, err := http.NewRequest("GET", "/carts", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(BuscarCarts)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var responseCart []modelos.Cart
	err = json.Unmarshal(rr.Body.Bytes(), &responseCart)
	if err != nil {
		t.Errorf("Error unmarshaling response body: %v", err)
	}

	// Add assertions here to validate the responseCart and other expectations
}
