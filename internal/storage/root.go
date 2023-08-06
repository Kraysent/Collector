package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config struct {
	DBName      string `yaml:"db"`
	User        string `yaml:"user"`
	PasswordEnv string `yaml:"password_env"`
	Password    string `yaml:"-"`
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
}

type Storage struct {
	Config Config
	DB     *sql.DB
}

func (s *Storage) Connect() error {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		s.Config.Host, s.Config.Port, s.Config.User,
		s.Config.Password, s.Config.DBName,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	s.DB = db

	return nil
}

func (s *Storage) Disconnect() error {
	return s.DB.Close()
}
