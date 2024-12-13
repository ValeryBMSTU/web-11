package config

type Config struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`

	API     api     `yaml:"api"`
	Usecase usecase `yaml:"usecase"`
	DB      db      `yaml:"db"`
	JWT     jwt     `yaml:"jwt"`
}

type api struct {
	MinPasswordSize int `yaml:"min_password_size"`
	MaxPasswordSize int `yaml:"max_password_size"`
	MinUsernameSize int `yaml:"min_username_size"`
	MaxUsernameSize int `yaml:"max_username_size"`
}

type usecase struct {
	DefaultMessage string `yaml:"default_message"`
}

type db struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBname   string `yaml:"dbname"`
}

type jwt struct {
	Secret string `yaml:"secret"`
}
