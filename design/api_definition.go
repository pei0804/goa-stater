package design

import (
	"os"

	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
	. "github.com/pei0804/goa-stater/design/constant"
	_ "github.com/pei0804/goa-stater/design/models"
	_ "github.com/pei0804/goa-stater/design/resource"
)

var _ = API("pei0804/goa-stater", func() {
	Title("pei0804/goa-stater")
	Description("pei0804/goa-stater")
	Contact(func() {
		Name("pei")
		Email("peeeei0804@gmail.com")
		URL("https://github.com/pei0804/goa-stater/issues")
	})
	License(func() {
		Name("MIT")
		URL("")
	})
	Docs(func() {
		Description("wiki")
		URL("")
	})
	Host(func() string {
		switch os.Getenv("Op") {
		case "develop":
			return "localhost:8080"
		case "staging":
			return "staging.com"
		case "production":
			return "production.com"
		}
		return "localhost:8080"
	}())
	Scheme(func() string {
		switch os.Getenv("Op") {
		case "develop":
			return "http"
		case "staging":
			return "https"
		case "production":
			return "https"
		}
		return "http"
	}())
	BasePath("/")
	Trait(AdminUserTrait, func() {
		Security(AdminAuth)
		Response(Unauthorized, ErrorMedia)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
	Trait(GeneralUserTrait, func() {
		Security(GeneralAuth)
		Response(Unauthorized, ErrorMedia)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
})
