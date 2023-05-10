package storage

type Storage interface {
	CreateLoginPassword(loginPassword *LoginPassword) error
	GetLoginPassword(user_id int, serviceName string) (*LoginPassword, error)
	DeleteLoginPassword(user_id int, serviceName string) error
}

type LoginPassword struct {
	UserID      int    `json:"user_id"`
	ServiceName string `json:"service_name"`
	Login       string `json:"login"`
	Password    string `json:"password"`
}
