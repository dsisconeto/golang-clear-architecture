package common

type IDomainError interface {
	DomainMessage() string
	Error() string
}

type DomainError struct {
	Message string
}

type EntityNotFoundError struct {
	*DomainError
}

func NewDomainError(message string) *DomainError {
	return &DomainError{Message: message}
}

func FactoryDomainError(err IDomainError) IDomainError {
	switch err.(type) {
	case EntityNotFoundError:
		{
			return EntityNotFoundError{
				DomainError: NewDomainError("entidade nao encontrada"),
			}
		}

	}
	return nil
}

func (d *DomainError) Error() string {
	return d.Message
}

func (d *DomainError) DomainMessage() string {
	return d.Message
}
