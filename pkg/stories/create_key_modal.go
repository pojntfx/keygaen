package stories

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/gridge/pkg/components"
)

type CreateKeyModalStory struct {
	Story
}

func (c *CreateKeyModalStory) Render() app.UI {
	return c.WithRoot(
		&components.CreateKeyModal{},
	)
}
