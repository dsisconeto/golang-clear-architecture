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
