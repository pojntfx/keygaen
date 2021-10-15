package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type KeyList struct {
	app.Compo

	Keys []GPGKey
}

func (c *KeyList) Render() app.UI {
	return app.Ul().
		Class("pf-c-data-list pf-m-grid-md").
		Aria("label", "Key list").
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
											Class("pf-c-data-list__cell").
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
																			Text(key.Email),
																		app.Text("- "+key.ID),
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
													Class("pf-l-stack pf-u-justify-content-center pf-u-mt-md").
													Body(
														app.Div().
															Class("pf-l-stack__item").
															Body(
																// TODO: Add dropdown items
																app.Button().
																	Class("pf-c-dropdown__toggle pf-m-plain").
																	Aria("expanded", "false").
																	Type("button").
																	Aria("label", "Actions").
																	Body(
																		app.I().
																			Class("fas fa-ellipsis-v").
																			Aria("hidden", true),
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
