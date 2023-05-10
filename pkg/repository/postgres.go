package repository

import (
	"database/sql"
	"fmt"
)

type PostgresDB struct {
	DB *sql.DB
}

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

func (p *PostgresDB) checkIfExists(user_id int, serviceName string) bool {
	var exists bool
	row := p.DB.QueryRow("SELECT EXISTS(SELECT pass FROM passwords WHERE user_id = $1 AND name_of_service = $2)",
		user_id, serviceName)
	if err := row.Scan(&exists); err != nil {
		exists = true
	}

	return exists
}
