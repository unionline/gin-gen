package config

// Config ...
type Config struct {
	AppName string `toml:"app_name"`
	Version string `toml:"version"`
	Port    int    `toml:"port"`
	GinMode string `toml:"gin_mode"`
	
	MySQL MySQL `toml:"mysql"`
	Redis Redis `toml:"redis"`
}

// MySQL ...
type MySQL struct {
	Addr               string `toml:"addr"`
	Username           string `toml:"username"`
	Password           string `toml:"password"`
	DbName             string `toml:"db_name"`
	MaxIdleConnections int    `toml:"max_idle_connections"`
	MaxOpenConnections int    `toml:"max_open_connections"`
}

// Redis ...
type Redis struct {
	Addr     string `toml:"addr"`
	Password string `toml:"password"`
	Prefix   string `toml:"prefix"`
	DB       int    `toml:"db"`
	PoolSize int    `toml:"pool_size"`
}
