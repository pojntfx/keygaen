package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	selectKeyInput = "select-key-input"
)

// ImportKeyModal is a modal to import keys with
type ImportKeyModal struct {
	app.Compo

	OnSubmit func(key []byte)                      // Handler to call to import the key
	OnCancel func(dirty bool, clear chan struct{}) // Handler to call when closing/cancelling the modal

	key []byte

	dirty bool
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
								ID:                         selectKeyInput,
								FileSelectionLabel:         "Drag and drop a key or select one",
								ClearLabel:                 "Clear",
								TextEntryInputPlaceholder:  "Or paste the key's contents here",
								TextEntryInputBlockedLabel: "File has been selected.",
								FileContents:               []byte(c.key),

								OnChange: func(fileContents []byte) {
									c.key = fileContents

									c.dirty = true
								},
								OnClear: func() {
									c.key = []byte{}
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
				Text("Import key"),
			app.Button().
				Class("pf-c-button pf-m-link").
				Type("button").
				Text("Cancel").
				OnClick(func(ctx app.Context, e app.Event) {
					handleCancel(c.clear, c.dirty, c.OnCancel)
				}),
		},
		OnClose: func() {
			handleCancel(c.clear, c.dirty, c.OnCancel)
		},
	}
}

func (c *ImportKeyModal) clear() {
	// Clear input value
	app.Window().GetElementByID(selectKeyInput).Set("value", app.Null())

	// Clear key
	c.key = []byte{}
	c.dirty = false
}
