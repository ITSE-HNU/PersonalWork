// Package app 应用封装层
package app

import (
	"fmt"
	"gitee.com/itse/personal-work/app/config"
	"gitee.com/itse/personal-work/app/entity"
	"gitee.com/itse/personal-work/app/service"
	"gitee.com/itse/personal-work/app/util"
	"github.com/google/wire"
	"gorm.io/gorm"
	"strconv"
)

// WireAppSet AppInjector for Wire DI
var WireAppSet = wire.NewSet(wire.Struct(new(App), "*"))

// App app 操作结构体
type App struct {
	DB           *gorm.DB
	Login        *service.LoginService
	PaperService *service.PaperService
}

// DbSourceInit 初始化数据库
func (a *App) DbSourceInit() error {
	db := a.DB.Model(&entity.Role{})

	if err := db.Create(&entity.Role{Name: "小学"}).Error; err != nil {
		return err
	}

	if err := db.Create(&entity.Role{Name: "初中"}).Error; err != nil {
		return err
	}

	if err := db.Create(&entity.Role{Name: "高中"}).Error; err != nil {
		return err
	}

	db = a.DB.Model(&entity.User{})

	if err := db.Create(&entity.User{
		Username: "张三1",
		Password: "123",
		RoleID:   1,
	}).Error; err != nil {
		return err
	}
	if err := db.Create(&entity.User{
		Username: "张三2",
		Password: "123",
		RoleID:   1,
	}).Error; err != nil {
		return err
	}
	if err := db.Create(&entity.User{
		Username: "张三3",
		Password: "123",
		RoleID:   1,
	}).Error; err != nil {
		return err
	}

	if err := db.Create(&entity.User{
		Username: "李四1",
		Password: "123",
		RoleID:   2,
	}).Error; err != nil {
		return err
	}
	if err := db.Create(&entity.User{
		Username: "李四2",
		Password: "123",
		RoleID:   2,
	}).Error; err != nil {
		return err
	}
	if err := db.Create(&entity.User{
		Username: "李四3",
		Password: "123",
		RoleID:   2,
	}).Error; err != nil {
		return err
	}

	if err := db.Create(&entity.User{
		Username: "王五1",
		Password: "123",
		RoleID:   3,
	}).Error; err != nil {
		return err
	}
	if err := db.Create(&entity.User{
		Username: "王五2",
		Password: "123",
		RoleID:   3,
	}).Error; err != nil {
		return err
	}
	if err := db.Create(&entity.User{
		Username: "王五3",
		Password: "123",
		RoleID:   3,
	}).Error; err != nil {
		return err
	}

	return nil
}

// ConfirmDB 查验数据库种是否有数据
func (a *App) ConfirmDB() error {
	db := a.DB.Model(&entity.Role{})
	roleRes := new([]entity.Role)

	if err := db.Find(roleRes).Error; err != nil {
		return err
	}
	roleLength := len(*roleRes)

	db = a.DB.Model(&entity.User{})
	userRes := new([]entity.User)

	if err := db.Find(userRes).Error; err != nil {
		return err
	}
	userLength := len(*userRes)

	if userLength == 0 && roleLength == 0 {
		err := a.DbSourceInit()
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}

// Run 运行
func Run() {
	config.InitConfig()
	a, _, err := BuildInjector()
	if err != nil {
		return
	}
	err = a.ConfirmDB()
	if err != nil {
		return
	}
	for {
		var currentUser = service.Current{
			Username: "",
			RoleID:   0,
			Role:     "",
			IsLogin:  false,
		}
		isFirst := true
		fmt.Println("请输入用户名、密码 (以空格隔开)")
		for !currentUser.IsLogin {
			if !isFirst {
				fmt.Println("请输入正确的用户名、密码")
			}
			current, err := a.Login.Login()
			if err != nil {
				fmt.Println("登录失败")
				isFirst = false
				continue
			}
			currentUser.Username = current.Username
			currentUser.Role = current.Role
			currentUser.RoleID = current.RoleID
			currentUser.IsLogin = true
			fmt.Println("登录成功！")
			break
		}

		isFailed := false
		isChanged := false
		isCheckFailed := false
	v1:
		for {
			if !isFailed && !isChanged && !isCheckFailed {
				fmt.Printf("准备生成%s数学题目，请输入生成题目数量(输入-1将退出当前用户，重新登录):\n", currentUser.Role)
				fmt.Println("题目数量的有效输入范围是 10 - 30")
			}
			if isFailed {
				fmt.Println("题目数量的有效输入范围是 10 - 30")
				fmt.Println("请重新输入")
			}
			if isChanged && !isFailed {
				fmt.Printf("准备生成%s数学题目，请输入生成题目数量\n", currentUser.Role)
			}
			if isCheckFailed {
				fmt.Println("请输入小学、初中和高中三个选项中的一个")
				fmt.Println("例: 切换为小学")
			}
			line := util.GetInput()
			id, mode, err := util.CheckMode(line)
			if err != nil {
				panic(err.Error())
			}
			if id == 4 {
				isCheckFailed = true
				isFailed = false
				continue
			}
			if id != -1 && id != 0 {
				currentUser.RoleID = id
				currentUser.Role = mode
				isChanged = true
				isFailed = false
				isCheckFailed = false
				fmt.Println("成功切换为" + mode)
				continue
			}
			count, err := strconv.Atoi(line)
			if count == -1 {
				isCheckFailed = false
				break v1
			}
			if count < 10 || count > 30 {
				isFailed = true
				isCheckFailed = false
				continue
			}
			err = a.PaperService.GeneratePaper(currentUser.Username, currentUser.RoleID, count)
			if err != nil {
				panic(err.Error())
			}
			isFailed = false
			isCheckFailed = false
		}
	}
}
