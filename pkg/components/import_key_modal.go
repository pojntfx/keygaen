package components

import (
	app "github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	selectKeyInput = "select-key-input"
)

type ImportKeyModal struct {
	app.Compo

	OnSubmit func(key string)
	OnCancel func()

	key string
}

func (c *ImportKeyModal) Render() app.UI {
	return &Modal{
		ID:    "import-key-modal",
		Title: "Import Key",
		Body: []app.UI{
			app.Form().
				Class("pf-c-form").
				ID("import-key-form").
				OnSubmit(func(ctx app.Context, e app.Event) {
					e.PreventDefault()

					c.OnSubmit(c.key)

					c.clear()
				}).
				Body(
					app.Div().
						Class("pf-c-form__group").
						Body(
							&FileUpload{
								ID:                    selectKeyInput,
								FileSelectionLabel:    "Drag and drop a key or select one",
								ClearLabel:            "Clear",
								TextEntryLabel:        "Or paste the key's contents here",
								TextEntryBlockedLabel: c.key,
								FileContents:          []byte(c.key),

								OnChange: func(fileContents []byte) {
									c.key = string(fileContents)
								},
								OnClear: func() {
									c.key = ""
								},
							},
						),
				),
		},
		Footer: []app.UI{
			app.Button().
				Class("pf-c-button pf-m-primary").
				Type("submit").
				Form("import-key-form").
				Text("Import Key"),
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
			c.OnCancel()
		},
	}
}

func (c *ImportKeyModal) clear() {
	// Clear input value
	app.Window().GetElementByID(selectKeyInput).Set("value", app.Null())

	// Clear key
	c.key = ""
}
