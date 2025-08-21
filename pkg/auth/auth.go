package auth

import "golang.org/x/crypto/bcrypt"

// Encrypt 使用 bcrypt 加密纯文本
func Encrypt(source string) (string, error) {
	// GenerateFromPassword 用于安全地生成密码哈希，其中的 cost 表示哈希计算的“工作因子”
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)

	return string(hashedBytes), err
}

// Compare 比较密文和明文是否相同
func Compare(hashedPassword, password string) error {
	// CompareHashAndPassword 用于比较哈希后的密码和明文密码是否匹配
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
