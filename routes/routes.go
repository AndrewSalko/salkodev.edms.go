package routes

import (
	"github.com/AndrewSalko/salkodev.edms.go/controller_departments"
	"github.com/AndrewSalko/salkodev.edms.go/controller_folders"
	"github.com/AndrewSalko/salkodev.edms.go/controller_orgs"
	"github.com/AndrewSalko/salkodev.edms.go/controller_users"
	"github.com/gin-gonic/gin"
)

func InitRoutes(routes *gin.Engine) {

	routes.POST("users/register", controller_users.Register)
	routes.POST("users/login", controller_users.Login)
	routes.GET("users/confirmregistration", controller_users.ConfirmRegistration)
	routes.GET("users/:uid", AuthMiddleware(), controller_users.GetUserByUID)
	routes.GET("users", AuthMiddleware(), controller_users.GetUsersPage)
	routes.POST("users/changepassword", AuthMiddleware(), controller_users.ChangePassword)
	routes.POST("users/refreshtoken", AuthMiddleware(), controller_users.RefreshToken)
	routes.POST("users/create", AuthMiddleware(), controller_users.CreateUser)
	routes.POST("users/modify", AuthMiddleware(), controller_users.ModifyUser)
	routes.DELETE("users/delete", AuthMiddleware(), controller_users.DeleteUser)

	routes.POST("users/groups/add", AuthMiddleware(), controller_users.AddToGroup)
	routes.POST("users/groups/remove", AuthMiddleware(), controller_users.RemoveFromGroup)

	routes.GET("orgs/:uid", AuthMiddleware(), controller_orgs.GetOrganizationByUID)
	routes.GET("orgs", AuthMiddleware(), controller_orgs.GetOrganizationsPage)
	routes.POST("orgs/create", AuthMiddleware(), controller_orgs.CreateOrganization)
	routes.POST("orgs/modify", AuthMiddleware(), controller_orgs.ModifyOrganization)
	routes.DELETE("orgs/delete", AuthMiddleware(), controller_orgs.DeleteOrganization)

	routes.GET("departments/:uid", AuthMiddleware(), controller_departments.GetDepartmentByUID)
	routes.GET("departments", AuthMiddleware(), controller_departments.GetDepartmentsPage)
	routes.POST("departments/create", AuthMiddleware(), controller_departments.CreateDepartment)
	routes.POST("departments/modify", AuthMiddleware(), controller_departments.ModifyDepartment)
	routes.DELETE("departments/delete", AuthMiddleware(), controller_departments.DeleteDepartment)

	routes.GET("folders/:uid", AuthMiddleware(), controller_folders.GetFolderByUID)
	routes.GET("folders", AuthMiddleware(), controller_folders.GetFoldersPage)
	routes.POST("folders/create", AuthMiddleware(), controller_folders.CreateFolder)
	routes.POST("folders/modify", AuthMiddleware(), controller_folders.ModifyFolder)
	routes.DELETE("folders/delete", AuthMiddleware(), controller_folders.DeleteFolder)

}
