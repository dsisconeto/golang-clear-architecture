package product

import "github.com/dsisconeto/arquitetura/core/common"

type Product struct {
	ID      uint
	Name    string
	CodeBar *CodeBar
}

func NewProduct(id uint, name string, codeBar *CodeBar) (*Product, error) {

	var err error
	if len(name) == 0 {
		err = common.NewDomainError("nome do produto e requerido")
	}

	if len(codeBar.Code) == 0 {
		err = common.NewDomainError("o codigo de barras e requerido")
	}

	if err != nil {
		return nil, err
	}

	return &Product{ID: id, Name: name, CodeBar: codeBar}, nil
}
