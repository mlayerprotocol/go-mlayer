package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type IpfsConfig struct {
	Host          string `mapstructure:"ipfs_url"`
	ProjectId     string `mapstructure:"ipfs_username"`
	ProjectSecret string `mapstructure:"ipfs_password"`
}

type EthChainConfig struct {
	Registry string `mapstructure:"bsc_registry"`
	ChainId  uint   `mapstructure:"bsc_chain_id"`
	RPCUrl   string `mapstructure:"bsc_rpc_url"`
}

type SqlConfig struct {
	DbDialect                  string         `mapstructure:"db_dialect"`
	DbHost                 string         `mapstructure:"db_host"`
	DbStoragePath                string         `mapstructure:"db_storage_dir"`
	DbPort					uint  `mapstructure:"db_port"`
	DbDatabase				string  `mapstructure:"db_database"`
	DbUser				string  `mapstructure:"db_user"`
	DbPassword 				string `mapstructure:"db_password"`
	DbSSLMode string `mapstructure:"db_sslmode"`
	DbTimezone string `mapstructure:"db_timezone"`
	DbMaxOpenConns int `mapstructure:"db_max_open_conns"`
	DbMaxIdleConns int `mapstructure:"db_max_idle_conns"`
	DbMaxConnLifetime int `mapstructure:"db_max_conn_lifetime_seconds"`
}

type MLChainAPI struct {
	url string  `mapstructure:"ml_api_url"`
}

type MainConfiguration struct {
	AddressPrefix		 	 string 		`mapstructure:"address_prefix"`
	NodePrivateKey           string         `mapstructure:"node_private_key"`
	NetworkPrivateKey        string         `mapstructure:"network_private_key"`      
	StakeContract            string         `mapstructure:"stake_contract"`
	ChainId                  uint           `mapstructure:"chain_id"`
	Token                    string         `mapstructure:"token_address"`
	EVMRPCUrl                string         `mapstructure:"evm_rpc_url"` // deprecated
	EVMRPCHttp               string         `mapstructure:"evm_rpc_http"`
	EVMRPCWss                string         `mapstructure:"evm_rpc_wss"`
	ProtocolVersion          string         `mapstructure:"protocol_version"`
	ChannelMessageBufferSize uint           `mapstructure:"channel_message_buffer_size"`
	Ipfs                     IpfsConfig     `mapstructure:"ipfs"`
	Bsc                      EthChainConfig `mapstructure:"bsc"`
	LogLevel                 string         `mapstructure:"log_level"`
	BootstrapPeers           []string       `mapstructure:"bootstrap_peers"`
	Listeners                []string       `mapstructure:"listeners"`
	RPCHost                  string         `mapstructure:"rpc_host"`
	WSAddress                string         `mapstructure:"ws_address"`
	RestAddress                string         `mapstructure:"rest_address"`
	RPCPort                  string         `mapstructure:"rpc_port"`
	RPCHttpPort              string         `mapstructure:"rpc_http_port"`
	Validator                bool           `mapstructure:"validator"`
	BootstrapNode            bool           `mapstructure:"bootstrap_node"`
	DataDir                  string         `mapstructure:"data_dir"`
	SQLDB                    SqlConfig     	`mapstructure:"sql"`
	MLBlockchainAPIUrl		 string			`mapstructure:"mlayer_api_url"`
	NetworkPublicKey      	 string  
	NetworkKeyAddress		 string 
}

var (
	Config MainConfiguration
)

func Init() *viper.Viper {
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvPrefix("ml")
	v.SetConfigName("config")        // name of config file (without extension)
	v.SetConfigType("toml")          // REQUIRED if the config file does not have the extension in the name
	v.AddConfigPath("/etc/mlayer/")  // path to look for the config file in
	v.AddConfigPath("$HOME/.mlayer") // call multiple times to add many search paths
	v.AddConfigPath(".")             // optionally look for config in the working directory

	err := v.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		// panic(fmt.Errorf("Config file: %o \n", err))
	}
	v.SetDefault("log_level", "info")
	v.SetDefault("channel_message_buffer_size", 128)
	v.SetDefault("db_max_open_conns",10)
	v.SetDefault("db_max_idle_conns",2)
	v.SetDefault("db_max_conn_lifetime_seconds",120)
	return v
}
func init() {
	c := LoadMainConfig()
	Config = *c
}

func LoadMainConfig() *MainConfiguration {
	v := Init()
	var c MainConfiguration
	if err := v.Unmarshal(&c); err != nil {
		fmt.Printf("Fatal: Couldn't read config: %s \n", err.Error())
	}
	c.NetworkPrivateKey = v.GetString("network_private_key") // needed to load from environment var
	if len(c.NetworkPrivateKey) == 0 {
		c.NetworkPrivateKey = v.GetString("network_private_key") // needed to load from environment var
	}
	if len(c.NodePrivateKey) == 0 {
		c.NodePrivateKey = v.GetString("node_private_key") // needed to load from environment var
	}

	c.DataDir = v.GetString("data_dir") // needed to load from environment var
	if len(c.DataDir) == 0 {
		c.DataDir = v.GetString("data_dir") // needed to load from environment var
	}
	return &c
}
