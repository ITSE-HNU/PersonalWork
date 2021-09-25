// Package config 配置文件生成
package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

var (
	// C 配置文件
	C    = new(CType)
	once sync.Once
)

// InitConfig 获得配置
func InitConfig() {
	once.Do(func() {

		path := "application.yml"

		file, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err.Error())
		}

		err = yaml.Unmarshal(file, &C)
		if err != nil {
			panic(err.Error())
		}
	})
}
