package admin

import (
	"EmployeeSelection/internal/crypto"
	"EmployeeSelection/internal/database"
	"EmployeeSelection/internal/database/model"
	"EmployeeSelection/internal/jwt"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// / 后台管理员
type Admin struct {
	model.BaseModel `gorm:"embedded"`
	Account         string `json:"account" gorm:"primaryKey;uniqueIndex;comment:用户账号"`
	PasswordHash    string `json:"-" gorm:"comment:密码哈希"`
	Username        string `json:"username" gorm:"index;comment:用户名"`
	Avatar          string `json:"avatar" gorm:"comment:用户头像"`
	Salt            string `json:"-" gorm:"comment:加密盐"`
	Role            int    `json:"role" gorm:"comment:权限"`
}

func init() {
	err := database.GetDatabase().AutoMigrate(&Admin{})
	if err != nil {
		panic(err)
	}
}

func (u Admin) CheckPassword(password string) bool {
	return u.PasswordHash == u.getPasswordHash(password)
}

func (u Admin) getPasswordHash(password string) string {
	return crypto.GetSha256(crypto.GetSha256(password))
}

func (u *Admin) setPasswordHash(password string) {
	u.PasswordHash = u.getPasswordHash(password)
}

func (u Admin) getJwt(expiresin int) string {
	newJwt := jwt.NewJwt(u.Salt)
	info := map[string]interface{}{
		"id":      u.ID,
		"account": u.Account,
		"avatar":  u.Avatar,
	}
	return newJwt.Create(info, time.Second*time.Duration(expiresin))
}

func (u Admin) GetUserByJwt(jwtStr string) (user Admin, err error) {
	var segInfo map[string]interface{}
	newJwt := jwt.NewJwt("")
	segInfo, err = newJwt.Decode(jwtStr)
	if err != nil {
		return
	}
	jsUid := segInfo["id"].(json.Number)
	uid, _ := jsUid.Int64()
	user.ID = uid
	model.GetModel(&user) // user.Department, user.Position empty
	log.Println("---FoundUser--By--Jwt---user.Salt------", user.Salt)
	newJwt = jwt.NewJwt(user.Salt)
	_, err = newJwt.Parse(jwtStr)
	if err != nil {
		log.Println("--GetUserByJwt--Error:", err)
	}
	return
}

func (u Admin) GetJwtInfo() jwt.JwtInfo {
	expiresin := 3600 * 24 * 7 // 有效期 7 天
	return jwt.JwtInfo{
		Token:     u.getJwt(expiresin),
		Expiresin: expiresin,
	}
}

func (u *Admin) ResetSalt() {
	u.Salt = crypto.GetRandString(64)
}

// 注册
func (u *Admin) Register(password string) (Admin, error) {
	user := new(Admin)
	if u.Account != "" {
		user.Account = u.Account
		model.GetModel(user)
	}
	if user.ID > 0 {
		return Admin{}, fmt.Errorf("error: Register Fail. User exists")
	}
	user.Account = u.Account
	user.Avatar = u.Avatar
	user.Username = u.Username
	user.ResetSalt()
	if password == "" {
		return Admin{}, fmt.Errorf("error: Register Fail. User password can not be empty")
	}
	user.setPasswordHash(password)
	affected, err := model.CreateModel(user)
	log.Println("affected: ", affected)
	return *user, err
}
