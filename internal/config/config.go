package config

type Config struct {
	Server    ServerConfig
	Database  DatabaseConfig
	Icanhazip IcanhazipConfig
}

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	DSN string
}

type IcanhazipConfig struct {
	Endpoint string
}
