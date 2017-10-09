package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("swaggerui", func() {
	Origin("*", func() {
		Methods("GET")
	})
	Files("/swaggerui/*filepath", "swaggerui/")
})

var _ = Resource("swagger", func() {
	Files("/swagger.json", "swagger/swagger.json")
})
