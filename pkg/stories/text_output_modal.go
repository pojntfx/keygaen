package stories

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/keygaen/pkg/components"
)

type TextOutputModalStory struct {
	Story

	modalOpen bool
}

func (c *TextOutputModalStory) Render() app.UI {
	return app.Div().Body(
		app.Button().
			Class("pf-c-button pf-m-primary").
			Type("button").
			Text("Open text output modal").
			OnClick(func(ctx app.Context, e app.Event) {
				c.modalOpen = !c.modalOpen
			}),
		app.If(
			c.modalOpen,
			c.WithRoot(
				&components.TextOutputModal{
					Title: "Example Text Output Modal",
					Tabs: []components.TextOutputModalTab{
						{
							Language: "text/plain",
							Title:    "Cypher",
							Body:     "asdfasdfasdfadsfasdf",
						},
						{
							Language: "text/plain",
							Title:    "Signature",
							Body:     "uas-02rioj23jd",
						},
					},
					OnClose: func() {
						c.modalOpen = false

						c.Update()
					},
				},
			),
		),
	)
}
