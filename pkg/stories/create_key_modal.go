package stories

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/keygean/pkg/components"
)

type CreateKeyModalStory struct {
	Story

	modalOpen bool
}

func (c *CreateKeyModalStory) Render() app.UI {
	return app.Div().Body(
		app.Button().
			Class("pf-c-button pf-m-primary").
			Type("button").
			Text("Create key").
			OnClick(func(ctx app.Context, e app.Event) {
				c.modalOpen = !c.modalOpen
			}),
		app.If(
			c.modalOpen,
			c.WithRoot(
				&components.CreateKeyModal{
					OnSubmit: func(fullName, email, _ string) {
						app.Window().Call("alert", fmt.Sprintf("Created key with full name %v, email %v and a password", fullName, email))

						c.modalOpen = false
					},
					OnCancel: func(dirty bool, clear chan struct{}) {
						c.modalOpen = false

						c.Update()

						clear <- struct{}{}
					},
				},
			),
		),
	)
}
