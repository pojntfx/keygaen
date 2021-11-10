package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// Modal is a generic modal with a close handler
type Modal struct {
	app.Compo

	ID           string   // HTML ID of the modal; must be unique across the page
	Icon         string   // Class of the icon to use to the left of the title; may be empty
	Title        string   // Title of the modal
	Class        string   // Class to be applied to the modal's outmost component
	Body         []app.UI // Body of the modal
	Footer       []app.UI // Footer of the modal
	DisableFocus bool     // Disable auto-focusing the modal; useful if a child component, i.e. an input should be focused instead

	OnClose func() // Handler to call when closing/cancelling the modal

	removeEventListener func()
}

func (c *Modal) Render() app.UI {
	classes := "pf-c-modal-box pf-m-modal pf-m-sm"
	if c.Class != "" {
		classes += " " + c.Class
	}

	return app.Div().
		Class("pf-c-backdrop").
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
					&Autofocused{
						Disable: c.DisableFocus,
						Component: app.Div().
							Aria("role", "dialog").
							Aria("label", c.Title).
							Aria("labelledby", c.ID).
							Aria("modal", true).
							Class(classes).
							TabIndex(-1).
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
					},
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
