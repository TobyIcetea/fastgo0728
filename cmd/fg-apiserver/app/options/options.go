package options

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

// MySQLOptions defines options for mysql database.
type MySQLOptions struct {
	Addr                  string        `json:"addr,omitempty" mapstructure:"addr"`
	Username              string        `json:"username,omitempty" mapstructure:"username"`
	Password              string        `json:"-" mapstructure:"password"`
	Database              string        `json:"database" mapstructure:"database"`
	MaxIdleConnections    int           `json:"max-idle-connections" mapstructure:"max-idle-connections,omitempty"`
	MaxOpenConnections    int           `json:"max-open-connections" mapstructure:"max-open-connections"`
	MaxConnectionLifeTime time.Duration `json:"max-connection-life-time,omitempty" mapstructure:"max-connection-life-time"`
}

// NewMySQLOptions create a `zero` value instance
func NewMySQLOptions() *MySQLOptions {
	return &MySQLOptions{
		Addr:                  "127.0.0.1:3306",
		Username:              "toby",
		Password:              "toby(#)666",
		Database:              "toby",
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifeTime: time.Duration(10) * time.Second,
	}
}

type ServerOptions struct {
	MySQLOptions *MySQLOptions `json:"mysql" mapstructure:"mysql"`
}

// NewServerOptions 创建带有默认值的 ServerOptions 实例
func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		MySQLOptions: NewMySQLOptions(),
	}
}

// Validate 校验 ServerOptions 中的选项是否合法
func (o *ServerOptions) Validate() error {
	// 验证 MySQL 地址格式
	if o.MySQLOptions.Addr == "" {
		return fmt.Errorf("MySQL server address cannot be empty")
	}
	// 检查地址格式是否为 host:port
	host, portStr, err := net.SplitHostPort(o.MySQLOptions.Addr)
	if err != nil {
		return fmt.Errorf("Invalid MySQL address format '%s': '%w'", o.MySQLOptions.Addr, err)
	}

	// 验证端口是否为数字
	port, err := strconv.Atoi(portStr)
	if err != nil || port < 1 || port > 65535 {
		return fmt.Errorf("Invalid MySQL port: %s", portStr)
	}

	// 验证主机名是否为空
	if host == "" {
		return fmt.Errorf("MySQL hostname cannot be empty")
	}

	// 验证凭证和数据库名
	if o.MySQLOptions.Username == "" {
		return fmt.Errorf("MySQL username cannot be empty")
	}

	if o.MySQLOptions.Password == "" {
		return fmt.Errorf("MySQL password cannot be empty")
	}

	if o.MySQLOptions.Database == "" {
		return fmt.Errorf("MySQL database name cannot be empty")
	}

	// 验证连接池参数
	if o.MySQLOptions.MaxIdleConnections <= 0 {
		return fmt.Errorf("MySQL max idle connections must be greater than 0")
	}
	
	if o.MySQLOptions.MaxOpenConnections <= 0 {
		return fmt.Errorf("MySQL max open connections must be greater than 0")
	}

	if o.MySQLOptions.MaxIdleConnections > o.MySQLOptions.MaxOpenConnections {
		return fmt.Errorf("MySQL max idle connections cannot be greater than max open connections")
	}

	if o.MySQLOptions.MaxConnectionLifeTime <= 0 {
		return fmt.Errorf("MySQL max connection life time must be greater than 0")
	}

	return nil
}
