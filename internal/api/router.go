package api

import (
	"EmployeeSelection/internal/api/auth"
	"EmployeeSelection/internal/api/employee_management"
	"EmployeeSelection/internal/api/selection"
	selectionContent "EmployeeSelection/internal/api/selection_content"
	"EmployeeSelection/internal/config"
	"EmployeeSelection/internal/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartWebService() {
	engine := gin.New()
	engine.MaxMultipartMemory = 8 << 20 //8M
	engine.Use(middleware.HandlerCORS())

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	engine.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	engine.Use(gin.Recovery())
	apiGroup := engine.Group(config.Api)

	authGroup := apiGroup.Group(config.Auth)
	adminGroup := authGroup.Group(config.Admin)
	adminGroup.POST(config.Register, auth.AdminRegister)
	adminGroup.POST(config.Login, auth.AdminLogin)

	userGroup := authGroup.Group(config.User)

	userGroup.POST(config.Login, auth.UserLogin)
	userGroup.Use(auth.CheckAdminByJWT())
	userGroup.POST(config.Register, auth.UserRegister)

	manageUser := apiGroup.Group(config.User)
	manageUser.Use(auth.CheckAdminByJWT())
	manageUser.POST(config.List, employeeManagement.GetUserList)
	manageUser.POST(config.Del, employeeManagement.UserDel)
	manageUser.POST(config.Edit, employeeManagement.UserInfoEdit)

	selectionGroup := manageUser.Group(config.Selection)
	selectionGroup.POST(config.List, selection.SelectionList)
	selectionGroup.POST(config.Create, selection.SelectionCreate)
	selectionGroup.POST(config.Del, selection.SelectionDel)
	selectionGroup.POST(config.Edit, selection.SelectionEdit)
	selectionContentGroup := selectionGroup.Group(config.Content)
	selectionContentGroup.POST(config.List, selectionContent.SelectionContentList)
	selectionContentGroup.POST(config.Create, selectionContent.SelectionContentCreate)
	selectionContentGroup.POST(config.Del, selectionContent.SelectionContentDel)
	selectionContentGroup.POST(config.Edit, selectionContent.SelectionContentEdit)

	port := fmt.Sprintf(":%v", config.GetWebServicePort())
	engine.Run(port)
}
