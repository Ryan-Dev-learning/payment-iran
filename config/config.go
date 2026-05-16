package config

import "github.com/BurntSushi/toml"

type Config struct {
	App       AppConfig       `toml:"app"`
	Server    ServerConfig    `toml:"server"`
	NovinoPay NovinoPayConfig `toml:"novinopay"`
}

type AppConfig struct {
	Name string `toml:"name"`
	Env  string `toml:"env"`
}

type ServerConfig struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

type NovinoPayConfig struct {
	CallbackUrl          string `toml:"callback_url"`
	MerchantID           string `toml:"merchant_id"`
	InitTransactionUrl   string `toml:"init_transaction_url"`
	VerifyTransactionUrl string `toml:"verify_transaction_url"`
}

func Load(path string) (Config, error) {
	var cfg Config
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}
