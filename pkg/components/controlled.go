package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

// Controlled sets DOM properties of the encapsulated component after it is mounted
type Controlled struct {
	app.Compo

	Component  app.UI                 // The component to be focused
	Properties map[string]interface{} // Map of properties to set
}

func (c *Controlled) Render() app.UI {
	return c.Component
}

func (c *Controlled) OnUpdate(ctx app.Context) {
	ctx.Defer(func(_ app.Context) {
		for key, value := range c.Properties {
			c.JSValue().Set(key, value)
		}
	})
}
