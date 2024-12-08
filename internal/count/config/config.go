package config

type Config struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port_count"`

	API     api     `yaml:"api"`
	Usecase usecase `yaml:"usecase"`
	DB      db      `yaml:"db"`
}

type api struct {
	MaxMessageSize int `yaml:"max_message_size"`
}

type usecase struct {
	DefaultMessageCount int `yaml:"default_message_count"`
}

type db struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	User   string `yaml:"user"`
	DBname string `yaml:"dbname"`
}
