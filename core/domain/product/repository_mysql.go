package product

import (
	"github.com/jinzhu/gorm"
)

type RepositoryMysql struct {
	DB *gorm.DB
}

type productMysql struct {
	ID      uint `gorm:"primary_key" `
	Name    string
	CodeBar string
}

func NewMysqlRepository(DB *gorm.DB) *RepositoryMysql {
	return &RepositoryMysql{DB: DB}
}

func (productMysql) TableName() string {
	return "products"
}

func (repo *RepositoryMysql) Mapper(p *productMysql) *Product {
	p2, _ := NewProduct(p.ID, p.Name, NewCodeBar(p.CodeBar))
	return p2
}

func (repo *RepositoryMysql) MapperReverser(p *Product) *productMysql {
	return &productMysql{p.ID, p.Name, p.CodeBar.Code}
}

func (repo *RepositoryMysql) Store(product *Product) error {
	prod := repo.MapperReverser(product)
	repo.DB.Create(&prod)
	return nil
}

func (repo *RepositoryMysql) Update(product *Product) error {
	prod := repo.MapperReverser(product)
	repo.DB.Update(&prod)
	return nil
}

func (repo *RepositoryMysql) FindOneById(id uint) *Product {
	var prod *productMysql
	repo.DB.Where("id = ?", id).First(&prod)
	return repo.Mapper(prod)
}

func (repo *RepositoryMysql) HasOneById(id uint) bool {
	var count int64
	repo.DB.Where("id = ?", id).Count(&count)
	return count > 0
}

func (repo *RepositoryMysql) HasOneByCodeBar(codeBar *CodeBar) bool {
	var count int64
	repo.DB.Where("code_bar = ?", codeBar.Code).Count(count)
	return count > 0
}

func (repo *RepositoryMysql) HasOneCodeBarAndDifferentId(codeBar *CodeBar, id uint) bool {
	var count int64
	repo.DB.Where("code_bar = ? AND id <> ?", codeBar.Code, id).Count(count)
	return count > 0
}
