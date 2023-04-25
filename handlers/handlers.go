package handlers

import (
	// standard library
	"crypto/sha256"
	"crypto/subtle"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	// internal
	"github.com/HelloWorld/goProductAPI/entity"
	// third party
	"github.com/gorilla/mux"
)

// GetCompanyHandler is used to get a company
func GetCompanyHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// Read company ID
		id := mux.Vars(r)["id"]
		company, err := entity.GetCompany(id)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		responseData, err := json.Marshal(company)
		if err != nil {
			// Check if it is No company error or any other error
			if errors.Is(err, entity.ErrNoCompany) {
				// Write Header if no related company found.
				rw.WriteHeader(http.StatusNoContent)
			} else {
				rw.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		// Write body with found company
		rw.Header().Add("content-type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(responseData)
	}
}

// CreateCompanyHandler is used to create a new company.
func CreateCompanyHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// Read incoming JSON from request body
		data, err := ioutil.ReadAll(r.Body)
		// If no body is associated return with StatusBadRequest
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// Check if data is proper JSON (data validation)
		var company entity.Company
		err = json.Unmarshal(data, &company)
		if err != nil {
			rw.WriteHeader(http.StatusExpectationFailed)
			rw.Write([]byte("Invalid Data Format"))
			return
		}
		err = entity.AddCompany(company)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		// return after writing Body
		rw.WriteHeader(http.StatusCreated)
		rw.Write([]byte("Added New Company"))
	}
}

// DeleteCompanyHandler deletes the company with given ID.
func DeleteCompanyHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// Read company ID
		id := mux.Vars(r)["id"]
		err := entity.DeleteCompany(id)
		if err != nil {
			// Check if it is No company error or any other error
			if errors.Is(err, entity.ErrNoCompany) {
				// Write Header if no related company found.
				rw.WriteHeader(http.StatusNoContent)
			} else {
				rw.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		// Write Header with Accepted Status (done operation)
		rw.WriteHeader(http.StatusAccepted)
	}
}

// UpdateCompanyHandler updates the company with given ID.
func UpdateCompanyHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// Read company ID
		id := mux.Vars(r)["id"]
		err := entity.DeleteCompany(id)
		if err != nil {
			if errors.Is(err, entity.ErrNoCompany) {
				rw.WriteHeader(http.StatusNoContent)
			} else {
				rw.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		// Read incoming JSON from request body
		data, err := ioutil.ReadAll(r.Body)
		// If no body is associated return with StatusBadRequest
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// Check if data is proper JSON (data validation)
		var company entity.Company
		err = json.Unmarshal(data, &company)
		if err != nil {
			rw.WriteHeader(http.StatusExpectationFailed)
			rw.Write([]byte("Invalid Data Format"))
			return
		}
		// AddCompany with the requested body
		err = entity.AddCompany(company)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Write Header if no related company found.
		rw.WriteHeader(http.StatusAccepted)
	}
}

func AuthHandler(h http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if ok {
			username := sha256.Sum256([]byte(os.Getenv("USER_NAME")))
			password := sha256.Sum256([]byte(os.Getenv("USER_PASS")))
			userHash := sha256.Sum256([]byte(user))
			passHash := sha256.Sum256([]byte(pass))
			validUser := subtle.ConstantTimeCompare(userHash[:], username[:]) == 1
			validPass := subtle.ConstantTimeCompare(passHash[:], password[:]) == 1
			if validPass && validUser {
				h.ServeHTTP(rw, r)
				return
			}
		}
		http.Error(rw, "No/Invalid Credentials", http.StatusUnauthorized)
	}
}
