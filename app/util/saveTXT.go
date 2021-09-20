package util

import (
	"fmt"
	"gitee.com/itse/personal-work/app/schema"
	"os"
	"strconv"
)

// SaveTxt 保存 Paper 为 txt
func SaveTxt(username string, params schema.Paper) error {
	_, err := PathExists("paperResult")
	if err != nil {
		return err
	}
	_, err = PathExists("paperResult/" + username)
	if err != nil {
		return err
	}
	//以追加模式打开文件，当文件不存在时生成文件
	txt, err := os.OpenFile("./paperResult/"+username+"/"+params.Name+".txt", os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	topic := params.Topic
	var str []byte
	for _, item := range topic {
		// 要追加的字符串
		str = []byte(strconv.Itoa(item.ID) + ". " + item.Title + "\n" + "\n")

		// 写入文件
		n, err := txt.Write(str)
		// 当 n != len(b) 时，返回非零错误
		if err == nil && n != len(str) {
			println(`错误代码：`, n)
			return err
		}
	}
	//return txt.Close()
	defer func(txt *os.File) {
		err := txt.Close()
		if err != nil {
			return
		}
	}(txt)
	return nil
}

// PathExists 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		// 创建文件夹
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			return true, nil
		}
	}
	return false, err
}
