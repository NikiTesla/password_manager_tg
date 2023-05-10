package telegram

import (
	"fmt"

	"github.com/NikiTesla/vk_telegram/pkg/repository"
)

type MockDB struct {
	records []*repository.LoginPassword
}

func (m *MockDB) CreateLoginPassword(loginPassword *repository.LoginPassword) error {
	m.records = append(m.records, loginPassword)

	return nil
}

func (m *MockDB) GetLoginPassword(user_id int, serviceName string) (string, string, error) {
	for _, v := range m.records {
		if user_id == v.UserID && serviceName == v.ServiceName {
			return v.Login, v.Password, nil
		}
	}

	return "", "", nil
}
func (m *MockDB) DeleteLoginPassword(user_id int, serviceName string) error {
	for i, v := range m.records {
		if user_id == v.UserID && serviceName == v.ServiceName {
			m.records = append(m.records[:i], m.records[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("there is no such record in database")
}
