package stories

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/gridge/pkg/components"
)

type ModalStory struct {
	Story

	modalOpen bool
}

func (c *ModalStory) Render() app.UI {
	c.EnableShallowReflection()

	return c.WithRoot(
		app.Div().Body(
			app.Button().
				Class("pf-c-button pf-m-primary").
				Type("button").
				Text("Open Modal").
				OnClick(func(ctx app.Context, e app.Event) {
					c.modalOpen = !c.modalOpen
				}),
			app.If(
				c.modalOpen,
				&components.Modal{
					ID:    "success-modal-story",
					Icon:  "fas fa-check",
					Title: "Key successfully generated!",
					Class: "pf-m-success",
					Body:  app.Text("It has been added to the key list."),
					Footer: app.Button().
						Aria("disabled", "false").
						Class("pf-c-button pf-m-primary").
						Type("button").
						Text("Continue to key list").
						OnClick(func(ctx app.Context, e app.Event) {
							c.modalOpen = false
						}),

					OnClose: func() {
						c.modalOpen = false

						c.Update()
					},
				},
			),
		),
	)
}
