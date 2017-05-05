package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("", func() {
	Title("")
	Description("")
	Contact(func() {
		Name("pei")
		Email("satak47cpc@gmail.com")
		URL("https://github.com/tikasan/goa-simple-sample/issues")
	})
	License(func() {
		Name("MIT")
		URL("")
	})
	Docs(func() {
		Description("wiki")
		URL("")
	})
	Host("localhost:8080")
	Scheme("http")
	Scheme("https")
	BasePath("/api/v1")

	Origin("http://localhost:8080/swagger", func() {
		Expose("")
		Methods("GET", "POST", "PUT", "DELETE")
		MaxAge(600)
		Credentials()
	})
})
