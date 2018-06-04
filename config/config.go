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
			Host:     "35.192.237.37",
			Username: "argyledb",
			Password: "Bro123=H$Argyle",
			Name:     "argyledb",
			Port:     "3306",
			Charset:  "utf8",
		},
	}
}
