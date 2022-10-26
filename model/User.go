package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

// 数据库操作

// GetUsername 获取用户名
func GetUsername(id int) string {
	var user User
	db.Select("username").Where("id = ?", id).First(&user)
	return user.Username
}

// CheckUser 查询用户是否存在
func CheckUser(name string) int {
	var user User
	db.Select("id").Where("username = ? ", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_UESD
	}
	return errmsg.SUCCSE
}

// BeforeSave 钩子函数
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	return
}

// ScryptPw 密码加密
func ScryptPw(passwd string) string {
	const KeyLen = 10
	salt := []byte{13, 24, 3, 47, 98, 56, 43, 2}
	hashPw, err := scrypt.Key([]byte(passwd), salt, 2<<10, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(hashPw)
	return fpw
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	// 密码加密，函数调用方式
	//data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// GetUsers 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	// 偏移量，置为-1表示取消分页功能
	offset := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offset = -1
	}
	err = db.Limit(pageSize).Offset(offset).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// EditUser 编辑用户
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	// 独立处理，不允许编辑密码
	err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
