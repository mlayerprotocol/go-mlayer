package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

type IpfsConfig struct {
	Host string `mapstructure:"ipfs_host"`
}

type EthChainConfig struct {
	Registry string `mapstructure:"bsc_registry"`
	ChainId  uint   `mapstructure:"bsc_chain_id"`
	RPCUrl   string `mapstructure:"bsc_rpc_url"`
}

type Configuration struct {
	NodePrivateKey           string         `mapstructure:"node_private_key"`
	EvmPrivateKey            string         `mapstructure:"evm_private_key"`
	StakeContract            string         `mapstructure:"stake_contract"`
	ChainId                  uint           `mapstructure:"chain_id"`
	Token                    string         `mapstructure:"token_address"`
	EVMRPCUrl                string         `mapstructure:"evm_rpc_url"`
	EVMRPCWss                string         `mapstructure:"evm_rpc_wss"`
	Network                  string         `mapstructure:"network"`
	ChannelMessageBufferSize uint           `mapstructure:"channel_message_buffer_size"`
	Ipfs                     IpfsConfig     `mapstructure:"ipfs"`
	Bsc                      EthChainConfig `mapstructure:"bsc"`
	LogLevel                 string         `mapstructure:"log_level"`
	BootstrapPeers           []string       `mapstructure:"bootstrap_peers"`
	Listeners                []string       `mapstructure:"listeners"`
	RPCHost                  string         `mapstructure:"rpc_host"`
	RPCPort                  string         `mapstructure:"rpc_port"`
	Validator                bool           `mapstructure:"validator"`
	BootstrapNode            bool           `mapstructure:"bootstrap_node"`
}

var (
	Config Configuration
)

func Init() *viper.Viper {
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvPrefix("icm")
	v.SetConfigName("config")     // name of config file (without extension)
	v.SetConfigType("toml")       // REQUIRED if the config file does not have the extension in the name
	v.AddConfigPath("/etc/icm/")  // path to look for the config file in
	v.AddConfigPath("$HOME/.icm") // call multiple times to add many search paths
	v.AddConfigPath(".")          // optionally look for config in the working directory

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
	c.EvmPrivateKey = v.GetString("private_key") // needed to load from environment var
	if len(c.EvmPrivateKey) == 0 {
		c.EvmPrivateKey = v.GetString("evm_private_key") // needed to load from environment var
	}

	if len(c.NodePrivateKey) == 0 {
		c.NodePrivateKey = v.GetString("node_private_key") // needed to load from environment var
	}
	return &c
}
