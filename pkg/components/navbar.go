package components

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

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
								Class("pf-c-brand").
								Src("/gridge/web/logo.png").
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
										Href("https://github.com/pojntfx/gridge").
										Target("_blank").
										Class("pf-c-button pf-m-plain").
										Aria("label", "Help").
										Body(
											app.I().
												Class("pf-icon pf-icon-help").
												Aria("hidden", true),
										),
								),
						),
				),
		)
}
