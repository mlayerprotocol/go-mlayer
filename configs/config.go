package configs

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)



type IpfsConfig struct {
	Host          string `toml:"ipfs_url"`
	ProjectId     string `toml:"ipfs_username"`
	ProjectSecret string `toml:"ipfs_password"`
}

type EthChainConfig struct {
	Registry string  `toml:"bsc_registry"`
	ChainId  int `toml:"bsc_chain_id"`
	RPCUrl   string  `toml:"bsc_rpc_url"`
}

type SqlConfig struct {
	DbDialect         string `toml:"db_dialect"`
	DbHost            string `toml:"db_host"`
	DbStoragePath     string `toml:"db_storage_dir"`
	DbPort            uint   `toml:"db_port"`
	DbDatabase        string `toml:"db_database"`
	DbUser            string `toml:"db_user"`
	DbPassword        string `toml:"db_password"`
	DbSSLMode         string `toml:"db_sslmode"`
	DbTimezone        string `toml:"db_timezone"`
	DbMaxOpenConns    int    `toml:"db_max_open_conns"`
	DbMaxIdleConns    int    `toml:"db_max_idle_conns"`
	DbMaxConnLifetime int    `toml:"db_max_conn_lifetime_seconds"`
	
}

type MLChainAPI struct {
	url string `toml:"ml_api_url"`
}

type ChainId string

func (n *ChainId) Bytes() []byte {
	s := string(*n)
	number, err := strconv.Atoi(string(s))
	if err == nil {
		return big.NewInt(int64(number)).FillBytes(make([]byte, 32))
	}
	return []byte(s)
}
func (n *ChainId) Equals(s string) bool {
	v := string(*n)
	return v == s
}

type MainConfiguration struct {
	AddressPrefix            string         `toml:"network_address_prefix"`
	StakeContract            string         `toml:"stake_contract"`
	ChainId                  ChainId        `toml:"chain_id"`
	Token                    string         `toml:"token_address"`
	EVMRPCUrl                string         `toml:"evm_rpc_url"` // deprecated
	EVMRPCHttp               string         `toml:"evm_rpc_http"`
	EVMRPCWss                string         `toml:"evm_rpc_wss"`
	ProtocolVersion          string         `toml:"protocol_version"`
	ChannelMessageBufferSize uint           `toml:"channel_message_buffer_size"`
	Ipfs                     IpfsConfig     `toml:"ipfs"`
	Bsc                      EthChainConfig `toml:"bsc"`
	LogLevel                 string         `toml:"log_level"`
	BootstrapPeers           []string       `toml:"bootstrap_peers"`
	ListenerAdresses         []string       `toml:"listener_addresses"`
	RPCHost                  string         `toml:"rpc_host"`
	WSAddress                string         `toml:"ws_address"`
	RestAddress              string         `toml:"rest_address"`
	RPCPort                  string         `toml:"rpc_port"`
	RPCHttpPort              string         `toml:"rpc_http_port"`
	Validator                bool           `toml:"validator"`
	BootstrapNode            bool           `toml:"bootstrap_node"`
	DataDir                  string         `toml:"data_dir"`
	SQLDB                    SqlConfig      `toml:"sql"`
	MLBlockchainAPIUrl       string         `toml:"mlayer_api_url"`
	PrivateKey               string         `toml:"private_key"`
	PublicKey        string
	OperatorAddress          string
	ValidatorPublicKey       string
	PrivateKeyBytes  []byte 
	PublicKeyBytes []byte 
	PrivateKeySECP  []byte 
	PublicKeySECP []byte 
	Context *context.Context
}

var (
	Config MainConfiguration
)
var possiblePaths = []string{
	"./config",
	"/etc/mlayer/config",
	"$HOME/.mlayer/config",
}

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
	 // panic( err)
	}
	v.SetDefault("log_level", "info")
	v.SetDefault("channel_message_buffer_size", 128)
	v.SetDefault("db_max_open_conns", 10)
	v.SetDefault("db_max_idle_conns", 2)
	v.SetDefault("db_max_conn_lifetime_seconds", 120)
	return v
}
func init() {
	c := LoadMainConfig()
	
	Config = *c
}
func LoadConfig() (*MainConfiguration, error) {
	var config MainConfiguration

	// Try loading the configuration file from the possible paths
	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			if _, err := toml.DecodeFile(path, &config); err != nil {
				return nil, fmt.Errorf("failed to decode config file %s: %w", path, err)
			}
			log.Printf("Loaded configuration from: %s",path)
			// Override with environment variables
			// kong.Parse(&config)
			
			return &config, nil
		}
	}
	
	return nil, fmt.Errorf("no valid configuration file found in paths: %v", possiblePaths)
}
func LoadMainConfig() *MainConfiguration {
	v := Init()
	parsed, _ := LoadConfig()
	m, err := json.Marshal(parsed)
	d := make(map[string]interface{})
	json.Unmarshal(m, &d)
	v.MergeConfigMap(d)
	if err != nil {         // Handle errors reading the config file
		panic( err)
	   }
	var c MainConfiguration
	if err := v.Unmarshal(&c); err != nil {
		fmt.Printf("Fatal: Couldn't read config: %s \n", err.Error())
	}

	
	c.PrivateKey = v.GetString("private_key") // needed to load from environment var
	if len(v.GetString("private_key")) > 0 {
		c.PrivateKey = v.GetString("private_key") // needed to load from environment var
	}

	

	
	if len(v.GetString("data_dir")) == 0 {
		c.DataDir = v.GetString("data_dir") // needed to load from environment var
	}

	
	return &c
}
