package stories

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/gridge/pkg/components"
)

type HomeStory struct {
	app.Compo
}

func (c *HomeStory) Render() app.UI {
	return &components.Home{}
}
