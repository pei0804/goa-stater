//go:generate goagen bootstrap -d github.com/pei0804/goa-stater/design

package main

import (
	"flag"

	"github.com/deadcheat/goacors"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/pei0804/goa-stater/normal/app"
	"github.com/pei0804/goa-stater/normal/controller"
	"github.com/pei0804/goa-stater/normal/design"
	"github.com/pei0804/goa-stater/normal/mymiddleware"
)

// Server 実行に必要な値を保持している
type Server struct {
	service *goa.Service
}

// NewServer Server構造体を作成する
func NewServer(s *goa.Service) *Server {
	return &Server{
		service: s,
	}
}

func (s *Server) mountController() {
	// Mount "example" controller
	example := controller.NewExampleController(s.service)
	app.MountExampleController(s.service, example)
	// Mount "swagger" controller
	swagger := controller.NewSwaggerController(s.service)
	app.MountSwaggerController(s.service, swagger)
	// Mount "swaggerui" controller
	swaggerui := controller.NewSwaggeruiController(s.service)
	app.MountSwaggeruiController(s.service, swaggerui)
}

func (s *Server) mountMiddleware(noSecure bool, env string) {
	s.service.Use(middleware.RequestID())
	s.service.Use(middleware.LogRequest(true))
	s.service.Use(middleware.ErrorHandler(s.service, true))
	s.service.Use(middleware.Recover())
	s.service.Use(goacors.WithConfig(s.service, design.CorsConfig[env]))

	if noSecure {
		app.UseAdminAuthMiddleware(s.service, mymiddleware.NewTestModeMiddleware())
		app.UseGeneralAuthMiddleware(s.service, mymiddleware.NewTestModeMiddleware())
	} else {
		app.UseAdminAuthMiddleware(s.service, mymiddleware.NewAdminUserAuthMiddleware())
		app.UseGeneralAuthMiddleware(s.service, mymiddleware.NewGeneralUserAuthMiddleware())
	}
}

func main() {
	service := goa.New("pei0804/goa-stater")
	var (
		port     = flag.String("port", ":8080", "addr to bind")
		env      = flag.String("env", "production", "実行環境 (production, staging, develop)")
		noSecure = flag.Bool("noSecure", false, "テストモードで実行。trueにすると、常にユーザーID: 1 参加しているインターンID: 1になります。")
	)
	flag.Parse()
	s := NewServer(service)
	s.mountMiddleware(*noSecure, *env)
	s.mountController()

	// Start service
	if err := service.ListenAndServe(*port); err != nil {
		service.LogError("startup", "err", err)
	}
}
