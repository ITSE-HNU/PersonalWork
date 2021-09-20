package app

import (
	"fmt"
	"gitee.com/itse/personal-work/app/config"
	"gitee.com/itse/personal-work/app/entity"
	"gitee.com/itse/personal-work/app/service"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var WireAppSet = wire.NewSet(wire.Struct(new(App), "*"))

type App struct {
	DB    *gorm.DB
	Login *service.LoginService
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

//ConfirmDB 查验数据库种是否有数据
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
	var currentUser = service.Current{
		Username: "",
		RoleID:   0,
		Role:     "",
		IsLogin:  false,
	}
	for !currentUser.IsLogin {
		current, err := a.Login.Login()
		if err != nil {
			fmt.Println("登录失败")
			return
		}
		currentUser.Username = current.Username
		currentUser.Role = current.Role
		currentUser.RoleID = current.RoleID
		current.IsLogin = true
		break
	}
	err = service.GeneratePaper(currentUser.RoleID)
	if err != nil {
		return
	}
}
