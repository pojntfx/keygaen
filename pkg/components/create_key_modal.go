package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type CreateKeyModal struct {
	app.Compo
}

func (c *CreateKeyModal) Render() app.UI {
	return app.P().Text("Create Key")
}
