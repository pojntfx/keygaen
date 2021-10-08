package stories

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/gridge/pkg/components"
)

type PasswordModalStory struct {
	Story
}

func (c *PasswordModalStory) Render() app.UI {
	return c.WithRoot(
		&components.PasswordModal{
			OnSubmit: func(password string) {
				app.Window().Call("alert", "Successfully entered a password")
			},
			OnCancel: func() {},
		},
	)
}
