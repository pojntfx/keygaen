package stories

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/keygean/pkg/components"
)

type SingleActionModalStory struct {
	Story

	modalOpen bool
}

func (c *SingleActionModalStory) Render() app.UI {
	return app.Div().Body(
		app.Button().
			Class("pf-c-button pf-m-primary").
			Type("button").
			Text("Open single action modal").
			OnClick(func(ctx app.Context, e app.Event) {
				c.modalOpen = !c.modalOpen
			}),
		app.If(
			c.modalOpen,
			c.WithRoot(
				&components.SingleActionModal{
					ID:          "success-modal-story",
					Icon:        "fas fa-check",
					Title:       "Key Successfully Generated!",
					Class:       "pf-m-success",
					Body:        "It has been added to the key list.",
					ActionLabel: "Continue to key list",

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
