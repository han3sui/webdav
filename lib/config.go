package lib

import (
	"os"

	"github.com/spf13/viper"
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

// UserMap 使用用户名作为key，生成map，用于后续查询
var UserMap = make(map[string]UserInfo)

// Config 配置文件实体
var Config AutoConfig

func init() {
	fileInfo, err := os.Stat("config.toml")
	if err != nil || fileInfo.IsDir() {
		Log().Panic("未找到配置文件")
		return
	}
	config := viper.New()
	config.AddConfigPath("./")
	config.SetConfigName("config")
	config.SetConfigType("toml")
	err = config.ReadInConfig()
	if err != nil {
		Log().Panic("加载配置文件失败: %s", err)
	}
	//直接反序列化为Struct
	if err := config.Unmarshal(&Config); err != nil {
		Log().Panic("序列化配置文件失败: %s", err)
	}
	Log().Info("配置文件：%v", Config)
	for _, v := range Config.User {
		_, ok := UserMap[v.Name]
		if ok {
			Log().Warning("存在重复的用户：%v", v.Name)
		} else {
			UserMap[v.Name] = v
		}
	}
}
