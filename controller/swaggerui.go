package controller

import (
	"github.com/goadesign/goa"
)

// SwaggeruiController implements the swaggerui resource.
type SwaggeruiController struct {
	*goa.Controller
}

// NewSwaggeruiController creates a swaggerui controller.
func NewSwaggeruiController(service *goa.Service) *SwaggeruiController {
	return &SwaggeruiController{Controller: service.NewController("SwaggeruiController")}
}
