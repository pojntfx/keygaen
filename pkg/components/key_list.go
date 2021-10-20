package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	keyListID = "key-list"
)

type KeyList struct {
	app.Compo

	Keys []GPGKey

	OnExport func(keyID string)
	OnDelete func(keyID string)

	expandedKeyID string

	removeEventListeners []func()
}

func (c *KeyList) Render() app.UI {
	return app.Ul().
		Class("pf-c-data-list pf-m-grid-md").
		Aria("label", "Key list").
		ID(keyListID).
		Body(
			app.Range(c.Keys).Slice(func(i int) app.UI {
				key := c.Keys[i]

				return app.Li().
					Class("pf-c-data-list__item").
					Body(
						app.Div().
							Class("pf-c-data-list__item-row").
							Body(
								app.Div().
									Class("pf-c-data-list__item-content").
									Body(
										app.Div().
											Class("pf-c-data-list__cell pf-u-pb-0").
											Body(
												app.Div().
													Class("pf-l-flex pf-m-column").
													Body(
														app.Div().
															Body(
																app.P().
																	Text(key.FullName),
																app.Small().
																	Body(
																		app.A().
																			Href("mailto:"+key.Email).
																			OnClick(func(ctx app.Context, e app.Event) {
																				// Prevent go-app from taking over this link
																				e.Call("stopImmediatePropagation")
																			}).
																			Text(key.Email),
																		app.Code().Text(" - "+key.Label),
																	),
															),
														app.If(
															key.Private || key.Public,
															app.Div().
																Class("pf-c-label-group").
																Body(
																	app.Div().
																		Class("pf-c-label-group__main").
																		Body(
																			app.Ul().
																				Class("pf-c-label-group__list").
																				Aria("role", "list").
																				Aria("label", "Key attributes").
																				Body(
																					app.If(
																						key.Public,
																						app.Li().
																							Class("pf-c-label-group__list-item").
																							Body(
																								app.Span().
																									Class("pf-c-label pf-m-green").
																									Body(
																										app.Span().
																											Class("pf-c-label__content").
																											Body(
																												app.Span().
																													Class("pf-c-label__icon").
																													Body(
																														app.I().
																															Class("fa-fw fas fa-globe").
																															Aria("hidden", true),
																													),
																												app.Text("Public"),
																											),
																									),
																							),
																					),
																					app.If(
																						key.Private,
																						app.Li().
																							Class("pf-c-label-group__list-item").
																							Body(
																								app.Span().
																									Class("pf-c-label pf-m-blue").
																									Body(
																										app.Span().
																											Class("pf-c-label__content").
																											Body(
																												app.Span().
																													Class("pf-c-label__icon").
																													Body(
																														app.I().
																															Class("fa-fw fas fa-user-lock").
																															Aria("hidden", true),
																													),
																												app.Text("Private"),
																											),
																									),
																							),
																					),
																				),
																		),
																),
														),
													),
											),
										app.Div().
											Class("pf-c-data-list__item-action").
											Body(
												app.Div().
													Class("pf-l-stack pf-u-justify-content-center pf-u-mt-md-on-md").
													Body(
														app.Div().
															Class("pf-l-stack__item").
															Body(
																app.Div().
																	Class(func() string {
																		dropdownClasses := "pf-c-dropdown"
																		if c.expandedKeyID == key.ID {
																			dropdownClasses += " pf-m-expanded"
																		}

																		return dropdownClasses
																	}()).
																	Body(
																		app.Button().
																			Class("pf-c-dropdown__toggle pf-m-plain").
																			ID("expand-key-actions-button-"+key.ID).
																			Aria("expanded", true).
																			Type("button").
																			Aria("label", "Actions").
																			OnClick(func(ctx app.Context, e app.Event) {
																				if c.expandedKeyID == key.ID {
																					c.closeKeyActions()

																					return
																				}

																				c.expandedKeyID = key.ID
																			}).
																			Body(
																				app.I().
																					Class("fas fa-ellipsis-v").
																					Aria("hidden", true),
																			),
																		app.If(
																			c.expandedKeyID == key.ID,
																			app.Ul().
																				Class("pf-c-dropdown__menu pf-m-align-right-on-md").
																				Aria("labelledby", "expand-key-actions-button-"+key.ID).
																				Body(
																					app.Li().
																						Body(
																							app.Button().
																								Class("pf-c-dropdown__menu-item pf-m-icon").
																								Type("button").
																								OnClick(func(ctx app.Context, e app.Event) {
																									c.closeKeyActions()

																									c.OnExport(key.ID)
																								}).
																								Body(
																									app.Span().
																										Class("pf-c-dropdown__menu-item-icon").
																										Body(
																											app.I().
																												Class("fas fa-file-export").
																												Aria("hidden", true),
																										),
																									app.Text("Export"),
																								),
																						),
																					app.Li().
																						Body(
																							app.Button().
																								Class("pf-c-dropdown__menu-item pf-m-icon").
																								Type("button").
																								OnClick(func(ctx app.Context, e app.Event) {
																									c.closeKeyActions()

																									c.OnDelete(key.ID)
																								}).
																								Body(
																									app.Span().
																										Class("pf-c-dropdown__menu-item-icon").
																										Body(
																											app.I().
																												Class("fas fa-trash").
																												Aria("hidden", true),
																										),
																									app.Text("Delete"),
																								),
																						),
																				),
																		),
																	),
															),
													),
											),
									),
							),
					)
			}),
		)
}

func (c *KeyList) closeKeyActions() {
	c.expandedKeyID = ""
}

func (c *KeyList) OnMount(ctx app.Context) {
	c.removeEventListeners = []func(){
		app.Window().AddEventListener("keyup", func(ctx app.Context, e app.Event) {
			if e.Get("key").String() == "Escape" {
				c.closeKeyActions()

				c.Update()
			}
		}),
		app.Window().AddEventListener("click", func(ctx app.Context, e app.Event) {
			// Close if we clicked outside the dropdown menu
			if c.expandedKeyID != "" {
				if dropdown := app.Window().Get("document").Call("querySelector", "#"+keyListID+" .pf-c-dropdown__menu"); !dropdown.IsNull() && !dropdown.Call("contains", e.Get("target")).Bool() {
					c.closeKeyActions()

					c.Update()
				}
			}
		}),
	}
}

func (c *KeyList) OnDismount() {
	if c.removeEventListeners != nil {
		for _, listener := range c.removeEventListeners {
			listener()
		}
	}
}
