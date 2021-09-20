package config

import "fmt"

// GORMType gorm 配置信息
type GORMType struct {
	Debug             bool `yaml:"Debug"`
	MaxLifetime       int  `yaml:"MaxLifetime"`
	MaxOpenConns      int  `yaml:"MaxOpenConns"`
	MaxIdleConns      int  `yaml:"MaxIdleConns"`
	EnableAutoMigrate bool `yaml:"EnableAutoMigrate"`
}

// DBType 数据库配置定义
type DBType struct {
	DBName string `yaml:"DBName"`
}

// DSN 得到数据库连接
func (d *DBType) DSN() string {
	return fmt.Sprintf(
		"%s", d.DBName,
	)
}

// CType 配置文件类型定义
type CType struct {
	GORM GORMType `yaml:"GORM"`
	DB   DBType   `yaml:"DB"`
}
