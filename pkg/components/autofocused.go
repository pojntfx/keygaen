package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Autofocused struct {
	app.Compo

	Disable   bool
	Component app.UI
}

func (c *Autofocused) Render() app.UI {
	return c.Component
}

func (c *Autofocused) OnMount(ctx app.Context) {
	if !c.Disable {
		ctx.Defer(func(_ app.Context) {
			c.JSValue().Call("focus")
		})
	}
}
