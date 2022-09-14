package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

type IpfsConfig struct {
	Host     string `mapstructure:"ipfs_url"`
	User     string `mapstructure:"ipfs_username"`
	Password string `mapstructure:"ipfs_password"`
}

type EthChainConfig struct {
	Registry string `mapstructure:"bsc_registry"`
	ChainId  uint   `mapstructure:"bsc_chain_id"`
	RPCUrl   string `mapstructure:"bsc_rpc_url"`
}

type Configuration struct {
	PrivateKey               string         `mapstructure:"private_key"`
	StakeContract            string         `mapstructure:"stake_contract"`
	ChainId                  uint           `mapstructure:"chain_id"`
	Token                    string         `mapstructure:"token_address"`
	RPCUrl                   string         `mapstructure:"rpc_url"`
	Network                  string         `mapstructure:"network"`
	ChannelMessageBufferSize uint           `mapstructure:"channel_message_buffer_size"`
	Ipfs                     IpfsConfig     `mapstructure:"ipfs"`
	Bsc                      EthChainConfig `mapstructure:"bsc"`
	LogLevel                 string         `mapstructure:"log_level"`
}

var (
	Config Configuration
)

func Init() *viper.Viper {
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvPrefix("icm")
	v.SetConfigName("config")         // name of config file (without extension)
	v.SetConfigType("toml")           // REQUIRED if the config file does not have the extension in the name
	v.AddConfigPath("/etc/splanch/")  // path to look for the config file in
	v.AddConfigPath("$HOME/.splanch") // call multiple times to add many search paths
	v.AddConfigPath(".")              // optionally look for config in the working directory

	err := v.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		panic(fmt.Errorf("Fatal: config file: %w \n", err))
	}
	v.SetDefault("log_level", "info")
	v.SetDefault("channel_message_buffer_size", 128)
	return v
}
func init() {
	c := LoadConfig()
	Config = *c
}

func LoadConfig() *Configuration {
	v := Init()
	var c Configuration
	if err := v.Unmarshal(&c); err != nil {
		fmt.Printf("Fatal: Couldn't read config: %w \n", err)
	}
	c.PrivateKey = v.GetString("private_key") // needed to load from environment var
	return &c
}
