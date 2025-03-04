package domain

type IProduct interface {
	Save(product *Product) error
	GetAll() ([]*Product, error)
	GetByID(id int32) (*Product, error)
	Update(product *Product) error
	Delete(id int32) error
}
