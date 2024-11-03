package service

import (
	"errors"
	"go-service-api/global"
	"go-service-api/model"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

// Create 创建用户
// 用户名和密码不能为空, 且用户名不能重复, 密码需要加密
func (u *UserService) Create(user *model.User) error {
	// 判断用户名和密码是否为空
	if user.Username == "" || user.Password == "" {
		return errors.New("用户名或密码不能为空")
	}
	// 生成UID
	user.UID = uuid.New().String()
	// 哈希密码
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hasedPassword)
	if err := global.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// Login 登录
func (u *UserService) Verify(user *model.User) (*model.User, error) {
	// 判断用户名和密码是否为空
	if user.Username == "" || user.Password == "" {
		return nil, errors.New("用户名或密码不能为空")
	}
	var newUser model.User
	// 查询用户
	if err := global.DB.Where("username = ?", user.Username).First(&newUser).Error; err != nil {
		return nil, err
	}
	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(user.Password)); err != nil {
		return nil, errors.New("密码错误")
	}
	return &newUser, nil
}

func (u *UserService) Info(user *model.User) error {
	if err := global.DB.Where("username = ?", user.Username).First(&user).Error; err != nil {
		return err
	}
	return nil
}

// Update 更新用户
func (u *UserService) Update(user *model.User) error {
	user.Dept = nil
	user.Role = nil
	if err := global.DB.Debug().Save(&user).Error; err != nil {
		return err
	}
	return nil
}

// Delete 删除用户
func (u *UserService) Delete(id uint) error {
	if err := global.DB.Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
