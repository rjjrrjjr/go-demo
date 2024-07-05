package config

// CustomerConfig 自定义配置
type CustomerConfig struct {
	FdbConf             FireDBConfig `yaml:"fdb_conf"`
	GuestCryptoGrayRate float32      `yaml:"guest_crypto_gray_rate"` // 嘉宾数据'加密写'灰度百分比例
}

// FireDBConfig fireDB配置项
type FireDBConfig struct {
	// WriteMode 写模式 0：明文写；1：脱敏写
	WriteMode int32 `yaml:"write_mode"`
	// QueryMode 读模式 0：明文查；1：脱敏查；2：兼容查
	QueryMode int32 `yaml:"query_mode"`
	// EnableLog 是否记录日志
	EnableLog bool   `yaml:"enable_log"`
	AesKey    string `yaml:"aes_key"`
}
