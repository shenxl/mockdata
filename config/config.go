package config

import (
	"io/ioutil"
	"os"

	"github.com/naoina/toml"
)

type AppConfig struct {
	Mysql struct {
		Addr    string
		ShowLog bool
	}

	Log struct {
		LogFile  string
		LogLevel string
	}
}

var App AppConfig

func Initialize(file string) error {

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	if err := toml.Unmarshal(buf, &App); err != nil {
		return err
	}

	if len(App.Log.LogLevel) == 0 {
		App.Log.LogLevel = "INFO"
	}

	// 初始化默认值

	return nil
}
