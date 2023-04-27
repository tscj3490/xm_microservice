package models

import "database/sql"

type CompanyType string

const (
	Corporations       CompanyType = "Corporations"
	NonProfit          CompanyType = "NonProfit"
	Cooperative        CompanyType = "Cooperative"
	SoleProprietorship CompanyType = "Sole Proprietorship"
)

//Company defines a structure for a company
type Company struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Employees   int32       `json:"employees"`
	Registered  bool        `json:"registered"`
	Type        CompanyType `json:"type"`
}

func (c *Company) Create(db *sql.DB) error {
	query := "INSERT INTO company (id, name, description, employees, registered, type) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	err := db.QueryRow(query, c.ID, c.Name, c.Description, c.Employees, c.Registered, c.Type).Scan(&c.ID)
	if err != nil {
		return err
	}
	return nil
}
