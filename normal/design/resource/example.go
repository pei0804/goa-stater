package resource

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
	. "github.com/pei0804/goa-stater/normal/design/constant"
	"github.com/pei0804/goa-stater/normal/design/media"
)

var _ = Resource("example", func() {
	DefaultMedia(media.MessageType)
	BasePath("/examples")
	Action("main", func() {
		Description("main example")
		Routing(
			GET("/main"),
		)
		Response(OK)
		UseTrait(GeneralUserTrait)
	})
	Action("sub", func() {
		Description("sub example")
		Routing(
			POST("/sub"),
		)
		Payload(func() {
			Attribute("message")
			Required("message")
		})
		Response(OK)
		UseTrait(GeneralUserTrait)
	})
	Action("id", func() {
		Description("ID")
		Routing(
			GET(":id"),
		)
		Params(func() {
			Param("id")
			Required("id")
		})
		Response(OK)
		UseTrait(GeneralUserTrait)
	})
})
