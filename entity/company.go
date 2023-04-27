package entity

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

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

// ErrNoCompany is used if no company found
var ErrNoCompany = errors.New("no company found")

func GetDataPath() string {
	// get the absolute path to the project's root directory
	projectDir, err := filepath.Abs(filepath.Dir("."))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// construct the path to the data.json file
	dataFilePath := filepath.Join(projectDir, "data", "data.json")

	return dataFilePath
}

// GetCompany takes id as input and returns the corresponding company, else it returns ErrNoCompany error.
func GetCompany(id string) (Company, error) {
	// Read JSON file
	data, err := ioutil.ReadFile(GetDataPath())
	if err != nil {
		return Company{}, err
	}
	// read companies
	var companies []Company
	err = json.Unmarshal(data, &companies)

	if err != nil {
		return Company{}, err
	}
	// iterate through company array
	for i := 0; i < len(companies); i++ {
		// if we find one company with the given ID
		if companies[i].ID == id {
			// return company
			return companies[i], nil
		}
	}
	return Company{}, ErrNoCompany
}

// DeleteCompany takes id as input and deletes the corresponding company, else it returns ErrNoCompany error.
func DeleteCompany(id string) error {
	// Read JSON file
	data, err := ioutil.ReadFile(GetDataPath())
	if err != nil {
		return err
	}
	// read companies
	var companies []Company
	err = json.Unmarshal(data, &companies)
	if err != nil {
		return err
	}
	// iterate through company array
	for i := 0; i < len(companies); i++ {
		// if we find one company with the given ID
		if companies[i].ID == id {
			companies = removeElement(companies, i)
			// Write Updated JSON file
			updatedData, err := json.Marshal(companies)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile("./data/data.json", updatedData, os.ModePerm)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return ErrNoCompany
}

// AddCompany adds an input company to the company list in JSON document.
func AddCompany(company Company) error {
	// Load existing companies and append the data to company list
	var companies []Company
	fmt.Println("111--", GetDataPath())
	data, err := ioutil.ReadFile(GetDataPath())
	fmt.Println("333--", data, err)
	if err != nil {
		return err
	}
	// Load our JSON file to memory using array of companies
	err = json.Unmarshal(data, &companies)
	if err != nil {
		return err
	}
	// Add new Company to our list
	companies = append(companies, company)

	// Write Updated JSON file
	updatedData, err := json.Marshal(companies)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./data/data.json", updatedData, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// removeElement is used to remove element from company array at given index
func removeElement(arr []Company, index int) []Company {
	ret := make([]Company, 0)
	ret = append(ret, arr[:index]...)
	return append(ret, arr[index+1:]...)
}
