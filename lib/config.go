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
	User []struct {
		Name     string   `toml:"name"`
		Password string   `toml:"password"`
		Auth     []string `toml:"auth,omitempty"`
		Dir      []string `toml:"dir"`
	} `toml:"User"`
}

var Config *viper.Viper

func init()  {
	fileInfo, err := os.Stat("config.toml")
	if err != nil || fileInfo.IsDir() {
		Log().Panic("未找到配置文件")
		return
	}
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./")
	err = viper.ReadInConfig()
	if err != nil {
		Log().Panic("加载配置文件失败: %s", err)
	}
	Config=viper.GetViper()
	//直接反序列化为Struct
	var conf AutoConfig
	if err :=Config.Unmarshal(&conf);err !=nil{
		Log().Panic("序列化配置文件失败: %s", err)
	}
	Log().Info("配置文件：",conf)
}


