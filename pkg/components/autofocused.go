package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// Autofocused calls `focus` on the encapsulated component after it is mounted
type Autofocused struct {
	app.Compo

	Disable   bool   // Disable the focus
	Component app.UI // The component to be focused
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
