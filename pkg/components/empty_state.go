package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

// EmptyState is the initial placeholder of the key list
type EmptyState struct {
	app.Compo

	OnCreateKey func() // OnCreateKey is the handler to call to create a key
	OnImportKey func() // OnCreateKey is the handler to call to import a key
}

func (c *EmptyState) Render() app.UI {
	return app.Div().
		Class("pf-c-empty-state").
		Body(
			app.Div().
				Class("pf-c-empty-state__content").
				Body(
					app.I().
						Class("fas fa-folder-open pf-c-empty-state__icon").
						Aria("hidden", true),
					app.H1().
						Class("pf-c-title pf-m-lg").
						Text("No keys yet"),
					app.Div().
						Class("pf-c-empty-state__body").
						Text("To get started, please create or import a key."),
					app.Button().
						Class("pf-c-button pf-m-primary").
						Type("button").
						Text("Create key").
						OnClick(func(ctx app.Context, e app.Event) {
							c.OnCreateKey()
						}),
					app.Div().
						Class("pf-c-empty-state__secondary").
						Body(
							app.Button().
								Class("pf-c-button pf-m-link").
								Type("button").
								Text("Or import key").
								OnClick(func(ctx app.Context, e app.Event) {
									c.OnImportKey()
								}),
						),
				),
		)
}
