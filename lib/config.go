package lib

import (
	"github.com/spf13/viper"
	"os"
)

type AutoConfig struct {
	Server struct {
		Addr  string `toml:"addr"`
		Debug bool   `toml:"debug"`
	} `toml:"Server"`
	User []UserInfo `toml:"User"`
}

type UserInfo struct {
	Name     string   `toml:"name"`
	Password string   `toml:"password"`
	Auth     []string `toml:"auth,omitempty"`
	Dir      []string `toml:"dir"`
}

var Config AutoConfig

func init()  {
	fileInfo, err := os.Stat("config.toml")
	if err != nil || fileInfo.IsDir() {
		Log().Panic("未找到配置文件")
		return
	}
	config:=viper.New()
	config.AddConfigPath("./")
	config.SetConfigName("config")
	config.SetConfigType("toml")
	err = config.ReadInConfig()
	if err != nil {
		Log().Panic("加载配置文件失败: %s", err)
	}
	//直接反序列化为Struct
	if err :=config.Unmarshal(&Config);err !=nil{
		Log().Panic("序列化配置文件失败: %s", err)
	}
	Log().Info("配置文件：",Config)
}


