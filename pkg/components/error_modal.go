package components

import (
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type ErrorModal struct {
	app.Compo

	ID          string
	Icon        string
	Title       string
	Class       string
	Body        string
	Error       error
	ActionLabel string

	OnClose  func()
	OnAction func()
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
