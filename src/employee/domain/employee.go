package domain

type Employee struct {
	ID int32
	FirstName string
	LastName string
	Age int16
	NumTel int32
}

func NewEmployee(firstname string, lastname string, age int16, numtel int32) *Employee{
	return &Employee{
		FirstName:  firstname,
		LastName: lastname,
		Age: age,
		NumTel: numtel,

	}
}