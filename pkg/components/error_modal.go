package components

import (
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// ErrorModal is a modal to display an error
type ErrorModal struct {
	app.Compo

	ID          string // HTML ID of the modal; must be unique across the page
	Icon        string // Class of the icon to use to the left of the title; may be empty
	Title       string // Title of the modal
	Class       string // Class to be applied to the modal's outmost component
	Body        string // Body text of the modal
	Error       error  // The error display (must not be nil)
	ActionLabel string // Text to display on the modal's primary action

	OnClose  func() // Handler to call when closing/cancelling the modal
	OnAction func() // Handler to call when triggering the modal's primary action
}

func (c *ErrorModal) Render() app.UI {
	return &Modal{
		ID:    c.ID,
		Icon:  c.Icon,
		Title: c.Title,
		Class: c.Class,
		Body: []app.UI{
			app.Text(c.Body),
			app.Div().
				Class("pf-c-code-editor pf-m-read-only pf-u-mt-lg").
				Body(
					app.Div().
						Class("pf-c-code-editor__main").
						Body(
							app.Textarea().
								Rows(len(strings.Split(c.Error.Error(), "\n"))).
								Style("width", "100%").
								Style("resize", "vertical").
								Style("border", "0").
								Class("pf-c-form-control").
								ReadOnly(true).
								Text(c.Error.Error()),
						),
				),
		},
		Footer: []app.UI{
			app.Button().
				Aria("disabled", "false").
				Class("pf-c-button pf-m-primary").
				Type("button").
				Text(c.ActionLabel).
				OnClick(func(ctx app.Context, e app.Event) {
					c.OnAction()
				}),
		},

		OnClose: func() {
			c.OnClose()
		},
	}
}
