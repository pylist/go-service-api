package config

var (
	Server *ServerConfig
	DB     *DBConfig
)

type ServerConfig struct {
	Mode string
	Host string
	Port int
}

type DBConfig struct {
	Types    string
	Name     string
	Username string
	Password string
	Host     string
	Charset  string
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	MaxIdleConns string
	// SetMaxOpenConns 设置打开数据库连接的最大数量
	MaxOpenConns string
	// SetConnMaxLifetime 设置了连接可复用的最大时间
	ConnMaxLifetime string
}
