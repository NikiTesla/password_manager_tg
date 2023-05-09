package environment

import (
	"database/sql"
	"fmt"
	"log"
)

type Environment struct {
	Config *Config
	DB     *sql.DB
}

func NewEnvironment(configFile string) (*Environment, error) {
	log.Println("Setting environment")

	cfg, err := NewConfig(configFile)
	if err != nil {
		return nil, fmt.Errorf("can't create environment because of config: %s", err.Error())
	}

	db, err := NewDataBase(cfg.DBConfig)
	if err != nil {
		return nil, fmt.Errorf("cant create environment bacause of database: %s", err.Error())
	}

	log.Printf("Host is %s\n", cfg.Host)
	log.Printf("Port is %d\n", cfg.Port)
	log.Printf("Database config is %+v", cfg.DBConfig)

	return &Environment{
		Config: cfg,
		DB:     db,
	}, nil
}
