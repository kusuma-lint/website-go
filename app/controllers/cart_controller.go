package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/kusuma-lint/website-go/app/models"
	"gorm.io/gorm"
)

func GetShoppingCartID(w http.ResponseWriter, r *http.Request) string {

	session, _ := store.Get(r, "shopping-cart-session")
	if session.Values["cart-id"] == nil {
		session.Values["cart-id"] = uuid.New().String()
		_ = session.Save(r, w)
	}
	return fmt.Sprintf("%v", session.Values["cart-id"])
}

func GetShoppingCart(db *gorm.DB, cartID string) (*models.Cart, error) {
	var cart models.Cart
	existCart, err := cart.GetCart(db, cartID)
	if err != nil {
		existCart, _ = cart.CreateCart(db, cartID)
	}

	fmt.Println(existCart)
	return existCart, nil
}

func (server *Server) GetCart(w http.ResponseWriter, r *http.Request) {

	var cart *models.Cart

	cartID := GetShoppingCartID(w, r)
	cart, _ = GetShoppingCart(server.DB, cartID)

	fmt.Println("cart id == ", cart.ID)
}

func (server *Server) AddItemToCart(w http.ResponseWriter, r *http.Request) {
	productID := r.FormValue("product_id")
	qty, _ := strconv.Atoi(r.FormValue("qty"))

	fmt.Println(productID)
	fmt.Println(qty)

	productModel := models.Product{}
	product, err := productModel.FindByID(server.DB, productID)
	if err != nil {
		http.Redirect(w, r, "/prodcuts/"+product.Slug, http.StatusSeeOther)
		return
	}

	if qty > product.Stock {
		http.Redirect(w, r, "/prodcuts/"+product.Slug, http.StatusSeeOther)
		return
	}

	var cart *models.Cart

	cartID := GetShoppingCartID(w, r)
	cart, _ = GetShoppingCart(server.DB, cartID)

	fmt.Println("cart id == ", cart.ID)

	http.Redirect(w, r, "/carts", http.StatusSeeOther)
}
