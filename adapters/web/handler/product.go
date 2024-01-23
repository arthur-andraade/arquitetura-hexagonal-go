package handler

import (
	"arquitetura-hexagonal/adapters/web/dto"
	"arquitetura-hexagonal/adapters/web/helpers"
	"arquitetura-hexagonal/application"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func MakeProductHandlers(r *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	r.Handle("/products/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET")
	r.Handle("/products", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST")
	r.Handle("/products/{id}/enable", n.With(
		negroni.Wrap(enableProduct(service)),
	)).Methods("PUT")
	r.Handle("/products/{id}/disable", n.With(
		negroni.Wrap(disableProduct(service)),
	)).Methods("PUT")
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]

		product, err := service.Get(id)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		serializeResponse(w, product)
	})
}

func createProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var dtoRequest dto.Product
		err := json.NewDecoder(r.Body).Decode(&dtoRequest)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(helpers.JsonError(err.Error()))
			return
		}

		product, err := service.Create(dtoRequest.Name, dtoRequest.Price)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(helpers.JsonError(err.Error()))
			return
		}

		serializeResponse(w, product)
	})
}

func enableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]

		product, err := service.Get(id)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		changeStatus(w, product, service.Enable)
	})
}

func disableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]

		product, err := service.Get(id)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		changeStatus(w, product, service.Disable)
	})
}

func changeStatus(
	response http.ResponseWriter,
	product application.ProductInterface,
	action func(application.ProductInterface) (application.ProductInterface, error),
) {

	product, err := action(product)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(helpers.JsonError(err.Error()))
	}

	serializeResponse(response, product)
}

func serializeResponse(
	response http.ResponseWriter,
	product application.ProductInterface,
) {

	dtoResponse := dto.NewProduct()
	dtoResponse.Rebind(product)

	err := json.NewEncoder(response).Encode(dtoResponse)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(helpers.JsonError(err.Error()))
		return
	}
}
