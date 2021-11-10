package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

// TextOutputModalTab is a tab which can be displayed in a text output modal
type TextOutputModalTab struct {
	Language string // Language of the text to be displayed (i.e. text/plain, French etc.)
	Title    string // Title of the text to be displayed
	Body     string // Text to be displayed
}

// TextOutputModal is a modal to display multiple text snippets
type TextOutputModal struct {
	app.Compo

	Title string               // Title of the modal
	Tabs  []TextOutputModalTab // Tabs to be displayed

	OnClose func() // Handler to call when closing/cancelling the modal

	selectedIndex int
}

func (c *TextOutputModal) Render() app.UI {
	return &Modal{
		ID:    "text-output-modal",
		Title: c.Title,
		Body: []app.UI{
			app.Div().
				Class("pf-c-code-editor pf-m-read-only").
				Body(
					app.Div().
						Class("pf-c-code-editor__header").
						Body(
							app.Div().
								Class("pf-c-code-editor__controls").
								Body(
									app.Div().
										Class("pf-c-tabs pf-m-box pf-m-color-scheme--light-300").
										Body(
											app.Ul().
												Class("pf-c-tabs__list").
												Body(
													app.Range(c.Tabs).Slice(func(i int) app.UI {
														classes := "pf-c-tabs__item"
														if c.selectedIndex == i {
															classes += " pf-m-current"
														}

														return app.Li().
															Class(classes).
															OnClick(func(ctx app.Context, e app.Event) {
																c.selectedIndex = i
															}).
															Body(
																app.Button().
																	Class("pf-c-tabs__link").
																	Body(
																		app.Span().
																			Class("pf-c-tabs__item-text").
																			Text(c.Tabs[i].Title),
																	),
															)
													}),
												),
										),
								),
							app.Div().
								Class("pf-c-code-editor__tab").
								Body(
									app.Span().
										Class("pf-c-code-editor__tab-icon").
										Body(
											app.I().
												Class("fas fa-code"),
										),
									app.Span().
										Class("pf-c-code-editor__tab-text").
										Text(c.Tabs[c.selectedIndex].Language),
								),
						),
					app.Div().
						Class("pf-c-code-editor__main").
						Body(
							app.Textarea().
								Rows(25).
								Style("width", "100%").
								Style("resize", "vertical").
								Style("border", "0").
								Class("pf-c-form-control").
								ReadOnly(true).
								Text(c.Tabs[c.selectedIndex].Body),
						),
				),
		},
		Footer: []app.UI{
			app.Button().
				Class("pf-c-button pf-m-primary").
				OnClick(func(ctx app.Context, e app.Event) {
					c.OnClose()
				}).
				Text("OK"),
		},
		OnClose: func() {
			c.OnClose()
		},
	}
}
