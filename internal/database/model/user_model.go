package model

import (
	"EmployeeSelection/internal/crypto"
	"EmployeeSelection/internal/database"
	"EmployeeSelection/internal/jwt"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// 普通用户
type User struct {
	BaseModel    `gorm:"embedded"`
	Account      string `json:"account" gorm:"primaryKey;uniqueIndex;comment:用户账号"`
	PasswordHash string `json:"password_hash" gorm:"-;comment:密码哈希"`
	Password     string `json:"password" gorm:"comment:密码"`
	Username     string `json:"username" gorm:"index;comment:用户名"`
	Avatar       string `json:"avatar" gorm:"comment:用户头像"`
	Salt         string `json:"salt" gorm:"-;comment:加密盐"`
	State        int    `json:"state" gorm:"comment:员工状态0正常，1停用"`
}

func init() {
	err := database.GetDatabase().AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
}

func (u User) CheckPassword(password string) bool {
	return u.PasswordHash == u.getPasswordHash(password)
}

func (u User) getPasswordHash(password string) string {
	return crypto.GetSha256(crypto.GetSha256(password))
}

func (u *User) setPasswordHash(password string) {
	u.PasswordHash = u.getPasswordHash(password)
}

func (u *User) setPassword(password string) {
	u.Password = password
}

func (u *User) CheckPwd(password string) bool {
	return u.Password == password
}

func (u User) getJwt(expiresin int) string {
	newJwt := jwt.NewJwt(u.Salt)
	info := map[string]interface{}{
		"id":      u.ID,
		"account": u.Account,
		"avatar":  u.Avatar,
	}
	return newJwt.Create(info, time.Second*time.Duration(expiresin))
}

func (u User) GetUserByJwt(jwtStr string) (user User, err error) {
	var segInfo map[string]interface{}
	newJwt := jwt.NewJwt("")
	segInfo, err = newJwt.Decode(jwtStr)
	if err != nil {
		return
	}
	jsUid := segInfo["id"].(json.Number)
	uid, _ := jsUid.Int64()
	user.ID = uid
	GetModel(&user) // user.Department, user.Position empty
	log.Println("---FoundUser--By--Jwt---user.Salt------", user.Salt)
	newJwt = jwt.NewJwt(user.Salt)
	_, err = newJwt.Parse(jwtStr)
	if err != nil {
		log.Println("--GetUserByJwt--Error:", err)
	}
	return
}

func (u User) GetJwtInfo() jwt.JwtInfo {
	expiresin := 3600 * 24 * 7 // 有效期 7 天
	return jwt.JwtInfo{
		Token:     u.getJwt(expiresin),
		Expiresin: expiresin,
	}
}

func (u *User) ResetSalt() {
	u.Salt = crypto.GetRandString(64)
}

// 注册
func (u *User) Register(password string) (User, error) {
	user := new(User)
	if u.Account != "" {
		user.Account = u.Account
		GetModel(user)
	}
	if user.ID > 0 {
		return User{}, fmt.Errorf("error: Register Fail. User exists")
	}
	user.Account = u.Account
	user.Avatar = u.Avatar
	user.Username = u.Username
	user.ResetSalt()
	if password == "" {
		return User{}, fmt.Errorf("error: Register Fail. User password can not be empty")
	}
	user.setPasswordHash(password)
	user.setPassword(password)
	affected, err := CreateModel(user)
	log.Println("affected: ", affected)
	return *user, err
}
