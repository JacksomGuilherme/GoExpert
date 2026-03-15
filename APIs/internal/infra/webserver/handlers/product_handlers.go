package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JacksomGuilherme/GoExpert/APIs/internal/dto"
	"github.com/JacksomGuilherme/GoExpert/APIs/internal/entity"
	"github.com/JacksomGuilherme/GoExpert/APIs/internal/infra/database"
	entityPkg "github.com/JacksomGuilherme/GoExpert/APIs/pkg/entity"
	"github.com/go-chi/chi"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// Create product godoc
// @Summary      Create product
// @Description  Create product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        request 	body 		dto.CreateProductInput 	true 	"product request"
// @Success      201
// @Failure      500  		{object}  	dto.Error
// @Router       /products [post]
// @Security ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Get Product godoc
// @Summary      Get product
// @Description  Get product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param		 id			path		string	false	"product ID" Format(uuid)
// @Success      200		{object}		entity.Product
// @Failure      401  		{object}  	dto.Error
// @Failure      404  		{object}  	dto.Error
// @Failure      500  		{object}  	dto.Error
// @Router       /products/{id} [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		error := dto.Error{Message: "id is required"}
		json.NewEncoder(w).Encode(error)
		return
	}

	product, err := h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// List Products godoc
// @Summary      List products
// @Description  List products
// @Tags         products
// @Accept       json
// @Produce      json
// @Param		 page		query		string	false	"page number"
// @Param		 limit		query		string	false	"limit"
// @Success      200		{array}		entity.Product
// @Failure      401  		{object}  	dto.Error
// @Failure      404  		{object}  	dto.Error
// @Failure      500  		{object}  	dto.Error
// @Router       /products [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// Update product godoc
// @Summary      Update product
// @Description  Update product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param		 id			path		string	false	"product ID" Format(uuid)
// @Param        request 	body 		dto.CreateProductInput 	true 	"product request"
// @Success      201
// @Failure      401  		{object}  	dto.Error
// @Failure      500  		{object}  	dto.Error
// @Router       /products/{id} [put]
// @Security ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		error := dto.Error{Message: "id is required"}
		json.NewEncoder(w).Encode(error)
		return
	}

	var product entity.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	_, err = h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete product godoc
// @Summary      Delete product
// @Description  Delete product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param		 id			path		string	false	"product ID" Format(uuid)
// @Success      200
// @Failure      401  		{object}  	dto.Error
// @Failure      500  		{object}  	dto.Error
// @Router       /products/{id} [delete]
// @Security ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		error := dto.Error{Message: "id is required"}
		json.NewEncoder(w).Encode(error)
		return
	}

	_, err := h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := dto.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusOK)
}
