package model

type ParametersDb struct {
	DbHost     string `yaml:"database_host"`
	DbPort     string `yaml:"database_port"`
	DbName     string `yaml:"database_name"`
	DbUser     string `yaml:"database_user"`
	DbPassword string `yaml:"database_password"`
}
