// Package model 逻辑处理层
package model

import (
	"math/rand"
	"time"
)

// randGenerator 随机种子封装
var randGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))
