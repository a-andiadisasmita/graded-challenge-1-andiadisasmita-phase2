package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"graded-challenge-1-andiadisasmita/models"

	"github.com/julienschmidt/httprouter"
)

// GetAllCustomers retrieves all customers from the database
func GetAllCustomers(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		rows, err := db.Query("SELECT id, name, email FROM customers")
		if err != nil {
			http.Error(w, "Failed to fetch customers", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var customers []models.Customer
		for rows.Next() {
			var c models.Customer
			if err := rows.Scan(&c.ID, &c.Name, &c.Email); err != nil {
				http.Error(w, "Failed to parse customer data", http.StatusInternalServerError)
				return
			}
			customers = append(customers, c)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

// GetCustomerByID retrieves a customer by ID
func GetCustomerByID(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := ps.ByName("id")
		var c models.Customer

		err := db.QueryRow("SELECT id, name, email, phone FROM customers WHERE id = $1", id).
			Scan(&c.ID, &c.Name, &c.Email, &c.Phone)
		if err == sql.ErrNoRows {
			http.Error(w, "Customer not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, "Failed to fetch customer", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(c)
	}
}

// CreateCustomer creates a new customer
func CreateCustomer(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var c models.Customer
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			http.Error(w, "Invalid input data", http.StatusBadRequest)
			return
		}

		query := `
			INSERT INTO customers (name, email, phone, created_at, updated_at)
			VALUES ($1, $2, $3, NOW(), NOW())
			RETURNING id, created_at, updated_at
		`
		err := db.QueryRow(query, c.Name, c.Email, c.Phone).
			Scan(&c.ID, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			http.Error(w, "Failed to create customer", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(c)
	}
}

// UpdateCustomer updates an existing customer's data
func UpdateCustomer(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := ps.ByName("id")
		var c models.Customer
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			http.Error(w, "Invalid input data", http.StatusBadRequest)
			return
		}

		query := `
			UPDATE customers
			SET name = $1, email = $2, phone = $3, updated_at = NOW()
			WHERE id = $4
		`
		res, err := db.Exec(query, c.Name, c.Email, c.Phone, id)
		if err != nil {
			http.Error(w, "Failed to update customer", http.StatusInternalServerError)
			return
		}

		rowsAffected, _ := res.RowsAffected()
		if rowsAffected == 0 {
			http.Error(w, "Customer not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Customer updated successfully"))
	}
}

// DeleteCustomer deletes a customer by ID
func DeleteCustomer(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := ps.ByName("id")

		res, err := db.Exec("DELETE FROM customers WHERE id = $1", id)
		if err != nil {
			http.Error(w, "Failed to delete customer", http.StatusInternalServerError)
			return
		}

		rowsAffected, _ := res.RowsAffected()
		if rowsAffected == 0 {
			http.Error(w, "Customer not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Customer deleted successfully"))
	}
}
