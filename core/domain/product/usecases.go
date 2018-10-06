package product

import "github.com/dsisconeto/arquitetura/core/common"

func UseCaseCreate(name string, codeBar *CodeBar, repository IRepository) (*Product, error) {
	var product *Product
	var err error

	if repository.HasOneByCodeBar(codeBar) {
		err := FactoryDomainError(DuplicateCodeBarError{})
		return nil, err
	}
	product, err = NewProduct(0, name, codeBar)
	if err != nil {
		return nil, err
	}

	err = repository.Store(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func UseCaseEdit(id uint, name string, codeBar *CodeBar, repository IRepository) (*Product, error) {
	var err error

	if !repository.HasOneById(id) {
		return nil, common.FactoryDomainError(common.EntityNotFoundError{})
	}

	if repository.HasOneCodeBarAndDifferentId(codeBar, id) {
		err := FactoryDomainError(DuplicateCodeBarError{})
		return nil, err
	}

	product, err := NewProduct(id, name, codeBar)
	if err != nil {
		return nil, err
	}

	err = repository.Update(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
