package components

import (
	app "github.com/maxence-charriere/go-app/v9/pkg/app"
)

type ExportKeyModal struct {
	app.Compo

	OnSubmit func(armor bool)
	OnCancel func()

	skipArmor bool
}

func (c *ExportKeyModal) Render() app.UI {
	return &Modal{
		ID:    "export-key-modal",
		Title: "Export Key",
		Body: []app.UI{
			app.Form().
				Class("pf-c-form").
				ID("export-key-form").
				OnSubmit(func(ctx app.Context, e app.Event) {
					e.PreventDefault()

					c.OnSubmit(!c.skipArmor)

					c.clear()
				}).
				Body(
					app.Div().
						Class("pf-c-form__group").
						Aria("role", "group").
						Body(
							app.Div().
								Class("pf-c-form__group-control").
								Body(
									app.Div().
										Class("pf-c-check").
										Body(
											&Controlled{
												Component: app.Input().
													Class("pf-c-check__input").
													Type("checkbox").
													ID("armor-checkbox").
													OnInput(func(ctx app.Context, e app.Event) {
														c.skipArmor = !c.skipArmor
													}),
												Properties: map[string]interface{}{
													"checked": !c.skipArmor,
												},
											},
											app.Label().
												Class("pf-c-check__label").
												For("armor-checkbox").
												Body(
													app.I().
														Class("fas fa-shield-alt pf-u-mr-sm"),
													app.Text("Armor"),
												),
											app.Span().
												Class("pf-c-check__description").
												Text("To increase portability, ASCII armor the key."),
										),
								),
						),
				),
		},
		Footer: []app.UI{
			app.Button().
				Class("pf-c-button pf-m-primary").
				Type("submit").
				Form("export-key-form").
				Text("Export key"),
			app.Button().
				Class("pf-c-button pf-m-link").
				Type("button").
				Text("Cancel").
				OnClick(func(ctx app.Context, e app.Event) {
					c.clear()

					c.OnCancel()
				}),
		},
		OnClose: func() {
			c.clear()
			c.OnCancel()
		},
	}
}

func (c *ExportKeyModal) clear() {
	c.skipArmor = false
}
