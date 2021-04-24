package model

import (
	"encoding/base64"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/tzr2020/gin_blog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

// 查询用户是否存在
func CheckUser(name string) int {
	var data User
	err = db.Select("id").Where("username=?", name).First(&data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	if data.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 新增用户
func CreateUser(data *User) int {
	// data.Password = ScryptPw(data.Password)
	err = db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询用户列表
func GetUSers(pageSize int, pageNum int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// 更新用户信息，密码除外
func EditUser(id int, data *User) int {
	var user User
	var m = make(map[string]interface{})
	m["username"] = data.Username
	m["role"] = data.Role
	err = db.Model(&user).Where("id=?", id).Updates(m).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id=?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 密码加密
func ScryptPw(password string) string {
	const keyLen = 10
	salt := make([]byte, 8)
	salt = []byte{34, 5, 67, 34, 90, 2, 67, 12}

	hashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, keyLen)
	if err != nil {
		log.Fatal(err)
	}
	finaPw := base64.StdEncoding.EncodeToString(hashPw)
	return finaPw
}

// 钩子函数
func (u *User) BeforeSave() {
	u.Password = ScryptPw(u.Password)
}

// 登录验证
func CheckLogin(username string, password string) int {
	var user User
	db.Where("username=?", username).First(&user)

	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 0 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCESS
}
