package controller

import (
	"github.com/goadesign/goa"
	"github.com/pei0804/goa-stater/normal/app"
)

// ExampleController implements the example resource.
type ExampleController struct {
	*goa.Controller
}

// NewExampleController creates a example controller.
func NewExampleController(service *goa.Service) *ExampleController {
	return &ExampleController{Controller: service.NewController("ExampleController")}
}

// ID runs the id action.
func (c *ExampleController) ID(ctx *app.IDExampleContext) error {
	// ExampleController_ID: start_implement

	// Put your logic here

	// ExampleController_ID: end_implement
	res := &app.Messagetype{}
	return ctx.OK(res)
}

// Main runs the main action.
func (c *ExampleController) Main(ctx *app.MainExampleContext) error {
	// ExampleController_Main: start_implement

	// Put your logic here

	// ExampleController_Main: end_implement
	res := &app.Messagetype{}
	return ctx.OK(res)
}

// Sub runs the sub action.
func (c *ExampleController) Sub(ctx *app.SubExampleContext) error {
	// ExampleController_Sub: start_implement

	// Put your logic here

	// ExampleController_Sub: end_implement
	res := &app.Messagetype{}
	return ctx.OK(res)
}
