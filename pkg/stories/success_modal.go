package stories

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/gridge/pkg/components"
)

type SuccessModalStory struct {
	Story

	modalOpen bool
}

func (c *SuccessModalStory) Render() app.UI {
	c.EnableShallowReflection()

	return c.WithRoot(
		app.Div().Body(
			app.Button().
				Class("pf-c-button pf-m-primary").
				Type("button").
				Text("Open Success Modal").
				OnClick(func(ctx app.Context, e app.Event) {
					c.modalOpen = !c.modalOpen
				}),
			app.If(
				c.modalOpen,
				&components.SuccessModal{
					ID:          "success-modal-story",
					Icon:        "fas fa-check",
					Title:       "Key successfully generated!",
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
