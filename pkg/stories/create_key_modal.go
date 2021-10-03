package stories

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/gridge/pkg/components"
)

type CreateKeyModalStory struct {
	app.Compo
}

func (c *CreateKeyModalStory) Render() app.UI {
	return &components.CreateKeyModal{}
}
