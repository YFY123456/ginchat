package models

import (
	"fmt"
	"ginchat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Identity      string
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP"` //`gorm:"default:2024-06-11 21:14:04.14"` //`gorm:"default:CURRENT_TIMESTAMP"`
	HeartbeatTime time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP"` //`gorm:"default:CURRENT_TIMESTAMP"`
	LoginOutTime  time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
func FindByID(id uint) UserBasic {
	user := UserBasic{}
	utils.DB.Where("id = ?", id).First(&user)
	return user
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	fmt.Println("数据数量:", len(data))
	result := utils.DB.Find(&data)
	if result.Error != nil {
		fmt.Println("查询错误:", result.Error)
		return nil
	}
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}
func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}
func DeleteUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}
func FindUserByNameAndPwd(name string, password string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ? and pass_word=?", name, password).First(&user)

	//token加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.MD5Encode(str)
	utils.DB.Model(&user).Where("id = ?", user.ID).Update("identity", temp)
	return user
}
func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ?", name).First(&user)
	return user
}
func FindUserByPhone(phone string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("Phone = ?", phone).First(&user)
}
func FindUserByEmail(email string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("email = ?", email).First(&user)
}
func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Updates(UserBasic{Name: user.Name, PassWord: user.PassWord, Phone: user.Phone, Email: user.Email})
}
