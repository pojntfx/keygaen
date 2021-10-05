package stories

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/gridge/pkg/components"
)

type CreateKeyModalStory struct {
	Story
}

func (c *CreateKeyModalStory) Render() app.UI {
	return c.WithRoot(
		&components.CreateKeyModal{
			OnSubmit: func(fullName, email, _ string) {
				app.Window().Call("alert", fmt.Sprintf("Created key with full name %v, email %v and a password", fullName, email))
			},
		},
	)
}
