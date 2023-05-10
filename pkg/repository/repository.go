package repository

// DataBase is interface to be implemented
// It is used as Bot database field
type DataBase interface {
	CreateLoginPassword(loginPassword *LoginPassword) error
	GetLoginPassword(user_id int, serviceName string) (string, string, error)
	DeleteLoginPassword(user_id int, serviceName string) error
}

// LoginPassword is analogue of row in table passwords
type LoginPassword struct {
	UserID      int    `json:"user_id"`
	ServiceName string `json:"service_name"`
	Login       string `json:"login"`
	Password    string `json:"password"`
}
