package domain

type IEmployee interface{
	Save(employee *Employee) error
	GetAll() ([]*Employee, error) 
	Update(employee *Employee) error
	Delete(id int32) error
}