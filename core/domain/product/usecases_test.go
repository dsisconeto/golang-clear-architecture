package product

import (
	"errors"
	"testing"
)

func TestUseCaseCreate(t *testing.T) {

	type tableTest struct {
		number    int
		arguments func() (name string, codeBar *CodeBar, repository IRepository)
		expected  func(product *Product, err error) error
	}
	tests := []tableTest{
		{
			number: 0,
			arguments: func() (name string, codeBar *CodeBar, repository IRepository) {
				repository = &RepositoryMock{
					MockStore: func(product *Product) error {
						return nil
					},
					MockHasOneByCodeBar: func(codeBar *CodeBar) bool {
						return true
					},
				}
				return "Product A", NewCodeBar("123"), repository
			},
			expected: func(product *Product, err error) error {
				if product != nil {
					return errors.New("produto tem que ser nulo")
				}
				if err == nil {
					return errors.New("o error e igual a nil, mas deveria ser DuplicateCodeBarError")
				}
				switch err.(type) {
				case DuplicateCodeBarError:
					return nil
				}
				return errors.New("error diferente de DuplicateCodeBarError")
			},
		},
	}
	for _, test := range tests {
		name, codeBar, repository := test.arguments()
		p, e := UseCaseCreate(name, codeBar, repository)
		if err := test.expected(p, e); err != nil {
			t.Errorf("UseCreate test number %d, %s", test.number, err)
		}
	}
}

func TestUseCaseEdit(t *testing.T) {

	type tableTest struct {
		number    int
		arguments func() (id uint, name string, codeBar *CodeBar, repository IRepository)
		expected  func(product *Product, err error) error
	}

	tests := []tableTest{
		{
			number: 0,
			arguments: func() (id uint, name string, codeBar *CodeBar, repository IRepository) {
				repository = &RepositoryMock{
					MockUpdate: func(product *Product) error {
						return nil
					},
					MockHasOneCodeBarAndDifferentId: func(codeBar *CodeBar, id uint) bool {
						return true
					},
				}
				return 1, "Product A", NewCodeBar("123"), repository
			},
			expected: func(product *Product, err error) error {
				if product != nil {
					return errors.New("produto tem que ser nulo")
				}
				if err == nil {
					return errors.New("o error e igual a nil, mas deveria ser DuplicateCodeBarError")
				}
				switch err.(type) {
				case DuplicateCodeBarError:
					return nil
				}
				return errors.New("error diferente de DuplicateCodeBarError")
			},
		},
	}

	for _, test := range tests {
		id, name, codeBar, repository := test.arguments()
		p, e := UseCaseEdit(id, name, codeBar, repository)
		if err := test.expected(p, e); err != nil {
			t.Errorf("UseEdit test number %d, %s", test.number, err)
		}
	}
}
