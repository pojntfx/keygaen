package stories

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/keygean/pkg/components"
)

type HomeStory struct {
	Story
}

func (c *HomeStory) Render() app.UI {
	return c.WithRoot(
		&components.Home{},
	)
}
