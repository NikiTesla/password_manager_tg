package repository

import (
	"database/sql"
	"fmt"
)

type PostgresDB struct {
	DB *sql.DB
}

// CreateLoginPassword checks if record does exist. If yes - just update username and password in database
// else creates new record with user id (id of telegram chat), name of service, username and password
func (p *PostgresDB) CreateLoginPassword(loginPassword *LoginPassword) error {
	if !p.checkIfExists(loginPassword.UserID, loginPassword.ServiceName) {
		_, err := p.DB.Exec("INSERT INTO passwords(user_id, name_of_service, username, pass) VALUES($1, $2, $3, $4)",
			loginPassword.UserID, loginPassword.ServiceName, loginPassword.Login, loginPassword.Password)
		if err != nil {
			return err
		}
	} else {
		_, err := p.DB.Exec("UPDATE passwords SET pass = $1, username = $2 WHERE user_id = $3",
			loginPassword.Password, loginPassword.Login, loginPassword.UserID)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetLoginPassword checks if record does exist. If not - returnng empty strings and error
// else reads row from database with user_id and serviceName, returns username and password
func (p *PostgresDB) GetLoginPassword(user_id int, serviceName string) (string, string, error) {
	if !p.checkIfExists(user_id, serviceName) {
		return "", "", fmt.Errorf("no data for the service")
	}

	var username, password string
	row := p.DB.QueryRow("SELECT username, pass FROM passwords WHERE user_id = $1 AND name_of_service = $2",
		user_id, serviceName)
	if err := row.Scan(&username, &password); err != nil {
		return "", "", err
	}

	return username, password, nil
}

// DeleteLoginPassword check if record exists. If not - returns error
// else deleting row from table with user id and service name
func (p *PostgresDB) DeleteLoginPassword(user_id int, serviceName string) error {
	if !p.checkIfExists(user_id, serviceName) {
		return fmt.Errorf("no data for the service")
	}

	_, err := p.DB.Exec("DELETE FROM passwords WHERE user_id = $1 AND name_of_service = $2",
		user_id, serviceName)
	if err != nil {
		return err
	}

	return nil
}

// checkIfExists makes select exists(..) and returns result
func (p *PostgresDB) checkIfExists(user_id int, serviceName string) bool {
	var exists bool
	row := p.DB.QueryRow("SELECT EXISTS(SELECT pass FROM passwords WHERE user_id = $1 AND name_of_service = $2)",
		user_id, serviceName)
	if err := row.Scan(&exists); err != nil {
		exists = true
	}

	return exists
}
