package telegram

import (
	"fmt"

	"github.com/NikiTesla/vk_telegram/pkg/repository"
)

// mock database for commands testing
type MockDB struct {
	records []*repository.LoginPassword
}

func (m *MockDB) CreateLoginPassword(loginPassword *repository.LoginPassword) error {
	for _, v := range m.records {
		if loginPassword.ServiceName == v.ServiceName && loginPassword.UserID == v.UserID {
			v.Login = loginPassword.Login
			v.Password = loginPassword.Password

			return nil
		}
	}
	m.records = append(m.records, loginPassword)

	return nil
}

func (m *MockDB) GetLoginPassword(user_id int, serviceName string) (string, string, error) {
	for _, v := range m.records {
		if user_id == v.UserID && serviceName == v.ServiceName {
			return v.Login, v.Password, nil
		}
	}

	return "", "", fmt.Errorf("no data for the service")
}

func (m *MockDB) DeleteLoginPassword(user_id int, serviceName string) error {
	for i, v := range m.records {
		if user_id == v.UserID && serviceName == v.ServiceName {
			m.records = append(m.records[:i], m.records[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("no data for the service")
}
