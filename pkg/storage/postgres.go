package storage

import "database/sql"

type PostgresDB struct {
	DB *sql.DB
}

func (db *PostgresDB) CreateLoginPassword(loginPassword *LoginPassword) error {
	return nil
}
func (db *PostgresDB) GetLoginPassword(user_id int, serviceName string) (*LoginPassword, error) {
	return nil, nil
}
func (db *PostgresDB) DeleteLoginPassword(user_id int, serviceName string) error {
	return nil
}
