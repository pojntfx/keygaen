package stories

import (
	"errors"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/keygean/pkg/components"
)

type ErrorModalStory struct {
	Story

	modalOpen bool
}

func (c *ErrorModalStory) Render() app.UI {
	return app.Div().Body(
		app.Button().
			Class("pf-c-button pf-m-primary").
			Type("button").
			Text("Open error modal").
			OnClick(func(ctx app.Context, e app.Event) {
				c.modalOpen = !c.modalOpen
			}),
		app.If(
			c.modalOpen,
			c.WithRoot(
				&components.ErrorModal{
					ID:          "error-modal-story",
					Icon:        "fas fa-times",
					Title:       "An Error Occurred",
					Class:       "pf-m-danger",
					Body:        "The following details may be of help:",
					Error:       errors.New(`panic: syscall/js: call of Value.Set on null`),
					ActionLabel: "Close",

					OnClose: func() {
						c.modalOpen = false

						c.Update()
					},
					OnAction: func() {
						c.modalOpen = false

						c.Update()
					},
				},
			),
		),
	)
}
