package product

type RepositoryMock struct {
	MockStore                       func(product *Product) error
	MockUpdate                      func(product *Product) error
	MockFindOneById                 func(id uint) *Product
	MockHasOneByCodeBar             func(codeBar *CodeBar) bool
	MockHasOneCodeBarAndDifferentId func(codeBar *CodeBar, id uint) bool
	MockHasOneById                  func(id uint) bool
}

func (m *RepositoryMock) Store(product *Product) error {
	return m.MockStore(product)
}

func (m *RepositoryMock) Update(product *Product) error {
	return m.MockUpdate(product)
}

func (m *RepositoryMock) FindOneById(id uint) *Product {
	return m.MockFindOneById(id)
}

func (m *RepositoryMock) HasOneById(id uint) bool {
	return m.MockHasOneById(id)
}

func (m *RepositoryMock) HasOneByCodeBar(codeBar *CodeBar) bool {
	return m.MockHasOneByCodeBar(codeBar)
}

func (m *RepositoryMock) HasOneCodeBarAndDifferentId(codeBar *CodeBar, id uint) bool {
	return m.MockHasOneCodeBarAndDifferentId(codeBar, id)
}
