package options

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/TobyIcetea/fastgo/internal/apiserver"
	genericoptions "github.com/TobyIcetea/fastgo/pkg/options"
)

type ServerOptions struct {
	MySQLOptions *genericoptions.MySQLOptions `json:"mysql" mapstructure:"mysql"`
	Addr         string                       `json:"addr" mapstructure:"addr"`
	// JWTKey 定义 JWT 密钥
	JWTKey string `json:"jwt-key" mapstructure:"jwt-key"`
	// Expiration 定义 JWT Token 的过期时间
	Expiration time.Duration `json:"expiration" mapstructure:"expiration"`
}

// NewServerOptions 创建带有默认值的 ServerOptions 实例
func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		MySQLOptions: genericoptions.NewMySQLOptions(),
		Addr:         "0.0.0.0:6666",
		Expiration:   time.Hour * 2,
	}
}

// Validate 校验 ServerOptions 中的选项是否合法
func (o *ServerOptions) Validate() error {
	// 验证 mysql 配置是否正确
	if err := o.MySQLOptions.Validate(); err != nil {
		return err
	}

	// 验证服务器地址
	if o.Addr == "" {
		return fmt.Errorf("server address cannot be empty")
	}

	// 检查地址格式是否为 host:port
	_, portStr, err := net.SplitHostPort(o.Addr)
	if err != nil {
		return fmt.Errorf("invalid server address format '%s': '%w'", o.Addr, err)
	}

	// 验证端口是否为数字且在有效范围内
	port, err := strconv.Atoi(portStr)
	if err != nil || port < 1 || port > 65535 {
		return fmt.Errorf("invalid server port: %s", portStr)
	}

	// 校验 JWTKey 长度
	if len(o.JWTKey) < 6 {
		return fmt.Errorf("JWTKey must be at least 6 characters long")
	}

	return nil
}

// Config 基于 ServerOptions 构建 apiserver.Config
func (o *ServerOptions) Config() (*apiserver.Config, error) {
	return &apiserver.Config{
		MySQLOptions: o.MySQLOptions,
		Addr:         o.Addr,
		JWTKey:       o.JWTKey,
		Expiration:   o.Expiration,
	}, nil
}
