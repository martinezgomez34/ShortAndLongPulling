package infraestructure

import (
	"database/sql"
	"errors"

	"github.com/martinezgomez34/abarrote/src/core"
	"github.com/martinezgomez34/abarrote/src/employee/domain"
)

type MySql struct {
	db *core.ConnMySQL
}

func NewSQL(db *core.ConnMySQL) *MySql {
	return &MySql{db: db}
}

func (r *MySql) Save(employee *domain.Employee) error {
	query := `INSERT INTO employees (firstname, lastname, age, numtel) VALUES (?, ?, ?, ?)`
	_, err := r.db.DB.Exec(query, employee.FirstName, employee.LastName, employee.Age, employee.NumTel)
	return err
}

func (r *MySql) GetAll() ([]*domain.Employee, error) {
	query := `SELECT id, firstname, lastname, age, numtel FROM employees`
	rows, err := r.db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []*domain.Employee
	for rows.Next() {
		var e domain.Employee
		err := rows.Scan(&e.ID, &e.FirstName, &e.LastName, &e.Age, &e.NumTel)
		if err != nil {
			return nil, err
		}
		employees = append(employees, &e)
	}
	return employees, nil
}


func (r *MySql) Update(employee *domain.Employee) error {
	query := `UPDATE employees SET firstname = ?, lastname = ?, age = ?, numtel = ? WHERE id = ?`
	_, err := r.db.DB.Exec(query, employee.FirstName, employee.LastName, employee.Age, employee.NumTel, employee.ID)
	return err
}

func (r *MySql) Delete(id int32) error {
	query := `DELETE FROM employees WHERE id = ?`
	_, err := r.db.DB.Exec(query, id)
	return err
}


func (r *MySql) GetByID(id int32) (*domain.Employee, error) {
	query := `SELECT id, firstname, lastname, age, numtel FROM employees WHERE id = ?`
	row := r.db.DB.QueryRow(query, id)

	var employee domain.Employee
	err := row.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.Age, &employee.NumTel)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &employee, nil
}
