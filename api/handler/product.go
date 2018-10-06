package handler

import (
	"github.com/dsisconeto/arquitetura/api/db"
	"github.com/dsisconeto/arquitetura/api/helper"
	"github.com/dsisconeto/arquitetura/core/domain/product"
	"net/http"
)

func ProductCreate(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Name    string `json:"name"`
		CodeBar string `json:"code_bar"`
	}{}
	err := helper.RequestJson(r, data)
	if err != nil {
		helper.ResponseBadRequest(w)
		return
	}
	repo := factoryProductRepository()
	defer repo.DB.Close()
	prod, err := product.UseCaseCreate(data.Name, product.NewCodeBar(data.CodeBar), factoryProductRepository())
	if err != nil {
		helper.ResponseErrorJson(w, err)
		return
	}
	helper.ResponseJson(w, prod)
}

func ProductRegisterRouters(mux *http.ServeMux) {

	mux.HandleFunc("/products/create", ProductCreate)
}

func factoryProductRepository() *product.RepositoryMysql {
	return product.NewMysqlRepository(db.GetMysql())
}
