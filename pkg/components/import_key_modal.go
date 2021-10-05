package components

import (
	app "github.com/maxence-charriere/go-app/v9/pkg/app"
)

type ImportKeyModal struct {
	app.Compo

	OnSubmit func(key string)
	OnCancel func()

	key string
}

func (c *ImportKeyModal) Render() app.UI {
	return app.Form().
		Class("pf-c-form").
		OnSubmit(func(ctx app.Context, e app.Event) {
			e.PreventDefault()

			c.OnSubmit(c.key)

			c.clear()
		}).
		Body(
			app.Div().
				Class("pf-c-form__group").
				Body(
					app.Div().
						Class("pf-c-file-upload").
						Body(
							app.Div().
								Class("pf-c-file-upload__file-select").
								Body(
									app.Div().
										Class("pf-c-input-group").
										Body(
											app.Input().
												Class("pf-c-form-control").
												ID("select-file-input").
												Aria("label", "Drag and drop a key or select one").
												ReadOnly(true).
												Placeholder("Drag and drop a key or select one").
												Aria("describedby", "select-file-button"),
											app.Button().
												Class("pf-c-button pf-m-control").
												Type("button").
												ID("select-file-button").
												Text("Select Key"),
											app.Button().
												Class("pf-c-button pf-m-control").
												Type("button").
												Disabled(c.key == "").
												Text("Clear").
												OnClick(func(ctx app.Context, e app.Event) {
													c.clear()
												}),
										),
								),
							app.Div().
								Class("pf-c-file-upload__file-details").
								Body(
									app.Textarea().
										Class("pf-c-form-control pf-m-resize-vertical").
										ID("enter-key-input").
										Aria("label", "Paste the key's contents here").
										Placeholder("Or paste the key's contents here").
										Required(true).
										OnInput(func(ctx app.Context, e app.Event) {
											c.key = ctx.JSSrc().Get("value").String()
										}).
										Text(c.key),
								),
						),
				),
			app.Div().
				Class("pf-c-form__group pf-m-action").
				Body(
					app.Div().
						Class("pf-c-form__group-control").
						Body(
							app.Div().
								Class("pf-c-form__actions").
								Body(
									app.Button().
										Class("pf-c-button pf-m-primary").
										Type("submit").
										Text("Import Key"),
									app.Button().
										Class("pf-c-button pf-m-link").
										Type("button").
										Text("Cancel").
										OnClick(func(ctx app.Context, e app.Event) {
											c.clear()

											c.OnCancel()
										}),
								),
						),
				),
		)
}

func (c *ImportKeyModal) clear() {
	c.key = ""
}
