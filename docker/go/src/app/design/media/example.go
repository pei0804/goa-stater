package media

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var MessageType = MediaType("application/vnd.messageType+json", func() {
	Description("example")
	Attributes(func() {
		Attribute("id", Integer, "id", func() {
			Default(0)
			Example(1)
		})
		Attribute("message", String, "message", func() {
			Default("")
			Example("ok")
		})
	})
	Required(
		"id",
		"message",
	)
	View("default", func() {
		Attribute("id")
		Attribute("message")
		Required(
			"id",
			"message",
		)
	})
})
