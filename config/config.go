package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Username string
	Password string
	Name     string
	Port     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Host:     "ideofuzion.mysql.database.azure.com",
			Username: "ideofuzion@ideofuzion",
			Password: "bro123=H$",
			Name:     "todoapp",
			Port:     "3306",
			Charset:  "utf8",
		},
	}
}
