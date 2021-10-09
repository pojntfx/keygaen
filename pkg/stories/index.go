package stories

import (
	"fmt"
	"net/url"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type SelfReferencingComponent interface {
	app.UI

	WithRoot(app.UI) app.UI

	EnableShallowReflection()
	SetOnRoot(onRoot func(root app.UI))
}

type Story struct {
	app.Compo

	root              app.UI
	shallowReflection bool

	onRoot func(root app.UI)
}

func (c *Story) WithRoot(root app.UI) app.UI {
	if c.shallowReflection || c.root == nil {
		c.root = root
	}

	if c.onRoot != nil {
		c.onRoot(c.root)
	}

	return c.root
}

func (c *Story) EnableShallowReflection() {
	c.shallowReflection = true
}

func (c *Story) SetOnRoot(onRoot func(root app.UI)) {
	c.onRoot = onRoot

	if c.root != nil {
		c.onRoot(c.root)
	}
}

const (
	activeTitleKey    = "activeTitle"
	standaloneKey     = "standalone"
	sidebarBreakpoint = 1200
)

type Index struct {
	app.Compo

	stories   map[string]SelfReferencingComponent
	storyCode string

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

		if t == "" {
			c.activeTitle = "Home"
		} else {
			c.activeTitle = t
		}
	}

	if c.stories == nil {
		c.stories = map[string]SelfReferencingComponent{
			"Home":                 &HomeStory{},
			"Create Key Modal":     &CreateKeyModalStory{},
			"Import Key Modal":     &ImportKeyModalStory{},
			"Encrypt/Sign Modal":   &EncryptAndSignModalStory{},
			"Decrypt/Verify Modal": &DecryptAndVerifyModalStory{},
			"File Upload":          &FileUploadStory{},
			"Password Modal":       &PasswordModalStory{},
			"Success Modal":        &SuccessModalStory{},
			"Modal":                &ModalStory{},
		}

		c.updateCodeQueries()
	}

	if app.Window().URL().Query().Has(standaloneKey) {
		return c.stories[c.activeTitle]
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
						Class("pf-c-page__main-section pf-u-p-0 pf-m-light pf-m-shadow-bottom").
						Body(
							app.Div().
								Class("pf-c-page__main-body").
								Body(
									app.Div().
										Class("pf-c-content pf-u-display-flex pf-u-justify-content-space-between").
										Body(
											app.H1().
												Class("pf-u-mb-0").
												Body(
													app.Text(c.activeTitle),
												),
											app.A().
												Class("pf-c-button pf-m-plain pf-u-ml-sm").
												Aria("label", "Fullscreen").
												OnClick(func(ctx app.Context, e app.Event) {
													u := app.Window().URL()

													q := u.Query()

													q.Set(standaloneKey, "true")

													u.RawQuery = q.Encode()

													app.Window().Call("open", u.String(), "_blank")
												}).
												Body(
													app.I().
														Class("fas fa-expand-arrows-alt").
														Aria("hidden", true),
												),
										),
								),
						),
					app.Section().
						Class("pf-c-page__main-section pf-u-p-0 pf-m-overflow-scroll").
						Body(
							app.Div().
								Class("pf-c-page__main-body").
								Body(c.stories[c.activeTitle]),
						),
					app.Section().
						Class("pf-c-page__main-section pf-u-p-0 pf-m-no-fill pf-m-light pf-m-shadow-top").
						Body(
							app.Div().
								Class("pf-c-page__main-body").
								Body(
									app.Div().
										Class("pf-c-code-block").
										Body(
											app.Div().
												Class("pf-c-code-block__header").
												Body(
													app.Div().
														Class("pf-c-code-block__actions").
														Body(
															app.Div().
																Class("pf-c-code-block__actions-item").
																Body(
																	app.Button().
																		Class("pf-c-button pf-m-plain").
																		Type("button").
																		Aria("label", "Reload props").
																		OnClick(func(ctx app.Context, e app.Event) {
																			c.Update()
																		}).
																		Body(
																			app.I().
																				Class("fas fa-sync").
																				Aria("hidden", true),
																		),
																),
														),
												),
											app.Div().
												Class("pf-c-code-block__content").
												Body(
													app.Pre().
														Class("pf-c-code-block__pre").
														Text(c.storyCode),
												),
										),
								),
						),
				),
		)
}

func (c *Index) OnMount(ctx app.Context) {
	c.sidebarOpen = true
	c.closeSidebarOnMobile()
}

func (c *Index) OnResize(ctx app.Context) {
	if c.sidebarOpen {
		c.closeSidebarOnMobile()
	} else {
		c.openSidebarOnDesktop()
	}
}

func (c *Index) OnAppUpdate(ctx app.Context) {
	if ctx.AppUpdateAvailable() {
		ctx.Reload()
	}
}

func (c *Index) OnNav(ctx app.Context) {
	if c.stories == nil {
		return
	}

	c.updateCodeQueries()
}

func (c *Index) updateCodeQueries() {
	for titleCandidate, story := range c.stories {
		if c.activeTitle == titleCandidate {
			story.SetOnRoot(func(root app.UI) {
				c.storyCode = fmt.Sprintf("%#v", root)
			})
		}
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
	if app.Window().Get("innerWidth").Int() < sidebarBreakpoint {
		c.sidebarOpen = false
	}
}

func (c *Index) openSidebarOnDesktop() {
	if app.Window().Get("innerWidth").Int() >= sidebarBreakpoint {
		c.sidebarOpen = true
	}
}
