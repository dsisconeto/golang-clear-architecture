package product

type IRepositoryWrite interface {
	Store(product *Product) error
	Update(product *Product) error
}

type IRepositoryRead interface {
	FindOneById(id uint) *Product
	HasOneById(id uint) bool
	HasOneByCodeBar(codeBar *CodeBar) bool
	HasOneCodeBarAndDifferentId(codeBar *CodeBar, id uint) bool
}

type IRepository interface {
	IRepositoryRead
	IRepositoryWrite
}

type MockRepository struct {
	MockStore                       func(product *Product) error
	MockUpdate                        func(product *Product) error
	MockFindOneById                 func(id uint) *Product
	MockHasOneByCodeBar             func(codeBar *CodeBar) bool
	MockHasOneCodeBarAndDifferentId func(codeBar *CodeBar, id uint) bool
	MockHasOneById                  func(id uint) bool
}

func (m *MockRepository) Store(product *Product) error {
	return m.MockStore(product)
}

func (m *MockRepository) Update(product *Product) error {
	return m.MockUpdate(product)
}

func (m *MockRepository) FindOneById(id uint) *Product {
	return m.MockFindOneById(id)
}

func (m *MockRepository) HasOneById(id uint) bool {
	return m.MockHasOneById(id)
}

func (m *MockRepository) HasOneByCodeBar(codeBar *CodeBar) bool {
	return m.MockHasOneByCodeBar(codeBar)
}

func (m *MockRepository) HasOneCodeBarAndDifferentId(codeBar *CodeBar, id uint) bool {
	return m.MockHasOneCodeBarAndDifferentId(codeBar, id)
}
