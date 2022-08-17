// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	syspermdept "ark-zero-admin/app/core/cmd/api/internal/handler/sys/perm/dept"
	syspermjob "ark-zero-admin/app/core/cmd/api/internal/handler/sys/perm/job"
	syspermmenu "ark-zero-admin/app/core/cmd/api/internal/handler/sys/perm/menu"
	syspermrole "ark-zero-admin/app/core/cmd/api/internal/handler/sys/perm/role"
	user "ark-zero-admin/app/core/cmd/api/internal/handler/user"
	"ark-zero-admin/app/core/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/login/captcha",
				Handler: user.GetLoginCaptchaHandler(serverCtx),
			},
		},
		rest.WithPrefix("/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/info",
				Handler: user.GetUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/permmenu",
				Handler: user.GetUserPermMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/profile/info",
				Handler: user.GetUserProfileInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/profile/update",
				Handler: user.UpdateUserProfileHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/password/update",
				Handler: user.UpdateUserPasswordHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/logout",
				Handler: user.LogoutHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.PermMenuAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: syspermmenu.GetSysPermMenuListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/add",
					Handler: syspermmenu.AddSysPermMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: syspermmenu.DeleteSysPermMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: syspermmenu.UpdateSysPermMenuHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/sys/perm/menu"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.PermMenuAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: syspermrole.GetSysRoleListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/add",
					Handler: syspermrole.AddSysRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: syspermrole.DeleteSysRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: syspermrole.UpdateSysRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/perm/update",
					Handler: syspermrole.UpdateSysRolePermMenuHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/sys/role"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.PermMenuAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: syspermdept.GetSysDeptListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/add",
					Handler: syspermdept.AddSysDeptHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: syspermdept.DeleteSysDeptHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: syspermdept.UpdateSysDeptHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/transfer",
					Handler: syspermdept.TransferSysDeptHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/sys/dept"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.PermMenuAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: syspermjob.GetSysJobListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/add",
					Handler: syspermjob.AddSysJobHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: syspermjob.DeleteSysJobHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: syspermjob.UpdateSysJobHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/sys/job"),
	)
}
