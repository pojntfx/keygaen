package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Modal struct {
	app.Compo

	ID     string
	Icon   string
	Title  string
	Class  string
	Body   []app.UI
	Footer []app.UI

	OnClose func()

	removeEventListener func()
}

func (c *Modal) Render() app.UI {
	classes := "pf-c-modal-box pf-m-modal pf-m-sm"
	if c.Class != "" {
		classes += " " + c.Class
	}

	return app.Div().
		Class("pf-c-backdrop pf-x-c-backdrop--nested").
		Body(
			app.Div().
				Class("pf-l-bullseye").
				OnClick(func(ctx app.Context, e app.Event) {
					// Close if we clicked outside the modal
					if e.Get("currentTarget").Call("isSameNode", e.Get("target")).Bool() {
						c.OnClose()
					}
				}).
				Body(
					app.Div().
						Aria("role", "dialog").
						Aria("label", c.Title).
						Aria("labelledby", c.ID).
						Aria("modal", true).
						Class(classes).
						Body(
							app.Button().
								Aria("disabled", "false").
								Aria("label", "Close").
								Class("pf-c-button pf-m-plain").
								Type("button").
								OnClick(func(ctx app.Context, e app.Event) {
									c.OnClose()
								}).
								Body(
									app.I().
										Class("fas fa-times").
										Aria("hidden", true),
								),
							app.Header().
								Class("pf-c-modal-box__header").
								Body(
									app.H1().
										ID(c.ID).
										Class("pf-c-modal-box__title pf-m-icon").
										Body(
											app.If(
												c.Icon != "",
												app.Span().
													Class("pf-c-modal-box__title-icon").
													Body(
														app.I().
															Class(c.Icon),
													),
											),
											app.Span().
												Class("pf-u-screen-reader").
												Text(c.Title),
											app.Span().
												Class("pf-c-modal-box__title-text").
												Text(c.Title),
										),
								),
							app.Div().
								Class("pf-c-modal-box__body").
								Body(c.Body...),
							app.If(
								c.Footer != nil,
								app.Footer().
									Class("pf-c-modal-box__footer").
									Body(c.Footer...),
							),
						),
				),
		)
}

func (c *Modal) OnMount(ctx app.Context) {
	c.removeEventListener = app.Window().AddEventListener("keyup", func(ctx app.Context, e app.Event) {
		if e.Get("key").String() == "Escape" {
			c.OnClose()
		}
	})
}

func (c *Modal) OnDismount() {
	if c.removeEventListener != nil {
		c.removeEventListener()
	}
}
