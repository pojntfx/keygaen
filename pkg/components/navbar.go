package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

// Navbar is the primary navigation menu
type Navbar struct {
	app.Compo
}

func (c *Navbar) Render() app.UI {
	return app.Header().
		Class("pf-c-page__header").
		Body(
			app.Div().
				Class("pf-c-page__header-brand").
				Body(
					app.A().
						Class("pf-c-page__header-brand-link").
						Href("#").
						Body(
							app.Img().
								Class("pf-c-brand", "pf-u-p-xs").
								Src("/web/logo-dark.png").
								Alt("Logo"),
						),
				),
			app.Div().
				Class("pf-c-page__header-tools").
				Body(
					app.Div().
						Class("pf-c-page__header-tools-group").
						Body(
							app.Div().
								Class("pf-c-page__header-tools-item").
								Body(
									app.A().
										Href("https://github.com/pojntfx/keygaen").
										Target("_blank").
										Class("pf-c-button pf-m-plain").
										Aria("label", "Help").
										Body(
											app.I().
												Class("pf-icon fas fa-question-circle").
												Aria("hidden", true),
										),
								),
						),
				),
		)
}
