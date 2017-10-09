package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var AdminAuth = APIKeySecurity("adminAuth", func() {
	Description("トークン(admin)")
	Header("X-Authorization")
})

var GeneralAuth = APIKeySecurity("generalAuth", func() {
	Description("トークン(general)")
	Header("X-Authorization")
})
