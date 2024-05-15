package infra

import "tech-challenge-fase-1/internal/adapter/driven/infra/config"

var (
	Host     = config.GetEnv("DB_HOST", "db")
	Port     = config.GetEnv("DB_PORT", "5432")
	User     = config.GetEnv("DB_USER", "tech-challenge-fase")
	Password = config.GetEnv("DB_PASSWORD", "tech-challenge-fase")
	DBname   = config.GetEnv("DB_NAME", "tech-challenge-fase")
)
