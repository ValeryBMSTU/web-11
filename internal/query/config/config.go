package config

type Config struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port_query"`

	API     api     `yaml:"api"`
	Usecase usecase `yaml:"usecase"`
	DB      db      `yaml:"db"`
}

type api struct {
	MaxMessageSize int `yaml:"max_message_size"`
}

type usecase struct {
	DefaultMessageQuery string `yaml:"default_message_query"`
}

type db struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	User   string `yaml:"user"`
	DBname string `yaml:"dbname"`
}
