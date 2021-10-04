package stories

import (
	"net/url"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	activeTitleKey = "activeTitle"
)

type Index struct {
	app.Compo

	stories     map[string]app.UI
	activeTitle string
	sidebarOpen bool
}

func (c *Index) Render() app.UI {
	additionalSidebarClasses := "pf-m-collapsed"
	if c.sidebarOpen {
		additionalSidebarClasses = "pf-m-expanded"
	}

	if c.activeTitle == "" {
		t, err := url.QueryUnescape(app.Window().URL().Query().Get(activeTitleKey))
		if err != nil {
			panic(err)
		}

		c.activeTitle = t
	}

	return app.Div().
		Class("pf-c-page").
		Body(
			app.Header().
				Class("pf-c-page__header").
				Body(
					app.A().
						Class("pf-c-skip-to-content pf-c-button pf-m-primary").
						Href("#main").
						Text("Skip to content"),
					app.Header().
						Class("pf-c-masthead").
						Body(
							app.Span().
								Class("pf-c-masthead__toggle").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Aria("label", "Global navigation").
										OnClick(func(ctx app.Context, e app.Event) {
											c.sidebarOpen = !c.sidebarOpen
										}).
										Body(
											app.I().Class("fas fa-bars").Aria("hidden", true),
										),
								),
							app.Div().
								Class("pf-c-masthead__main").
								Body(
									app.A().
										Class("pf-c-masthead__brand").
										Href("/").
										Body(
											app.Img().
												Class("pf-c-brand").
												Src("/web/logo.png").
												Alt("Logo"),
										),
									app.Em().
										Class("pf-c-brand pf-u-ml-sm").
										Text("Stories"),
								),
						),
				),
			app.Div().
				Class("pf-c-page__sidebar "+additionalSidebarClasses).
				Aria("hidden", !c.sidebarOpen).
				Body(
					app.Div().
						Class("pf-c-page__sidebar-body").
						Body(
							app.Nav().
								Class("pf-c-nav").
								Aria("label", "Global").
								Body(
									app.Ul().
										Class("pf-c-nav__list").
										Body(
											app.Range(c.stories).Map(func(title string) app.UI {
												linkClasses := "pf-c-nav__link"
												if c.activeTitle == title {
													linkClasses += " pf-m-current"
												}

												return app.Li().
													Class("pf-c-nav__item").
													Body(
														app.A().
															Class(linkClasses).
															OnClick(func(ctx app.Context, e app.Event) {
																c.setActiveTitle(title, ctx)

																c.closeSidebarOnMobile()
															}).
															Text(title),
													)
											}),
										),
								),
						),
				),
			app.Main().
				ID("main").
				Class("pf-c-page__main").
				TabIndex(-1).
				Body(
					app.Section().
						Class("pf-c-page__main-section pf-m-limit-width pf-m-light pf-m-shadow-bottom").
						Body(
							app.Div().
								Class("pf-c-page__main-body").
								Body(
									app.Div().
										Class("pf-c-content").
										Body(
											app.H1().
												Text(c.activeTitle),
										),
								),
						),
					app.Section().
						Class("pf-c-page__main-section pf-m-limit-width pf-m-overflow-scroll").
						Body(
							app.Div().
								Class("pf-c-page__main-body").
								Body(
									c.stories[c.activeTitle],
								),
						),
				),
		)
}

func (c *Index) OnMount(ctx app.Context) {
	c.stories = map[string]app.UI{
		"Home":             &HomeStory{},
		"Create Key Modal": &CreateKeyModalStory{},
	}

	c.sidebarOpen = true
	c.closeSidebarOnMobile()
}

func (c *Index) OnAppUpdate(ctx app.Context) {
	if ctx.AppUpdateAvailable() {
		ctx.Reload()
	}
}

func (c *Index) setActiveTitle(title string, ctx app.Context) {
	c.activeTitle = title

	u := app.Window().URL()

	q := u.Query()
	q.Set(activeTitleKey, url.QueryEscape(c.activeTitle))

	u.RawQuery = q.Encode()
	ctx.NavigateTo(u)
}

func (c *Index) closeSidebarOnMobile() {
	if app.Window().Get("screen").Get("width").Int() < 1200 {
		c.sidebarOpen = false
	}
}
