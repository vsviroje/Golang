package config

// GeneralConfig struct
type GeneralConfig struct {
	ApplicationPort string     `mapstructure:"app_port" json:"app_port"`
	MySqlDBConfig   *SQLConfig `mapstructure:"mysqldb" json:"mysqldb"`
}

// MySQL DB Config struct
type SQLConfig struct {
	Enabled         bool   `mapstructure:"enabled" json:"enabled"`
	Driver          string `mapstructure:"driver" json:"driver"`
	DatabaseName    string `mapstructure:"db_name" json:"db_name"`
	DatabaseHost    string `mapstructure:"host" json:"host"`
	TimeOut         int    `mapstructure:"timeout" json:"timeout"`
	DialTimeOut     int64  `mapstructure:"dial_timeout" json:"dial_timeout"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns" json:"max_idle_conns"`
	MaxOpenConnsint int    `mapstructure:"max_open_conns_int" json:"max_open_conns_int"`
	Username        string `mapstructure:"username" json:"username"`
	Password        string `mapstructure:"password" json:"password"`
	SSLMode         string `mapstructure:"ssl_mode" json:"ssl_mode"`
	URI             string `mapstructure:"uri" json:"uri"`
}
