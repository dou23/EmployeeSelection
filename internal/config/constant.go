package config

// url api路径
const (
	Api      = "/api"      //一级路径
	Auth     = "/auth"     //认证
	Admin    = "/admin"    //管理员
	User     = "/user"     //用户
	List     = "/list"     //列表
	Del      = "/del"      //删除
	Edit     = "/edit"     //删除
	Register = "/register" // 用户登陆
	Login    = "/login"    // 用户登陆

	Selection = "/selection" //评选
	Content   = "/content"   //评选内容
	Create    = "/create"    //评选
)

// 请求头
const (
	AuthToken = "Auth-Token" // token
)
