package data

type IRepository interface {
	GetAll() ([]interface{}, error)
	GetById(id int) (interface{}, error)
	Create(entity interface{}) error
	Update(entity interface{}) error
	Delete(id int) error
}
