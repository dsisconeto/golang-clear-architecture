package product

import (
	"github.com/dsisconeto/arquitetura/core/common"
)

type DuplicateCodeBarError struct {
	*common.DomainError
}

func FactoryDomainError(err common.IDomainError) common.IDomainError {
	switch err.(type) {
	case DuplicateCodeBarError:
		{
			return DuplicateCodeBarError{
				DomainError: common.NewDomainError("não pode existir dois productos com mesmo codigo de barras"),
			}
		}

	}
	return nil
}
