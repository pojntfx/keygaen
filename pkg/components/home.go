package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

var demoKeys = []GPGKey{
	{
		ID:       "039292",
		FullName: "Isegard Example",
		Email:    "isegard@example.com",
		Private:  true,
	},
	{
		ID:       "838431",
		FullName: "Fred Example",
		Email:    "fred@example.com",
		Private:  true,
		Public:   true,
	},
	{
		ID:       "123456",
		FullName: "Alice Example",
		Email:    "alice@example.com",
		Public:   true,
	},
	{
		ID:       "319312",
		FullName: "Bob Example",
		Email:    "bob@example.com",
		Public:   true,
	},
}

type Home struct {
	app.Compo
}

func (c *Home) Render() app.UI {
	return app.Div().
		Class("pf-c-page").
		Body(
			app.A().
				Class("pf-c-skip-to-content pf-c-button pf-m-primary").
				Href("#gridge-main").
				Body(
					app.Text("Skip to content"),
				),
			&Navbar{},
			app.Main().
				Class("pf-c-page__main").
				ID("gridge-main").
				TabIndex(-1).
				Body(
					app.Section().
						Class("pf-c-page__main-section pf-m-light pf-m-no-padding pf-u-px-sm-on-xl").
						Body(
							&Toolbar{
								OnCreateKey: func() {
									app.Window().Call("alert", "Created key")
								},
								OnImportKey: func() {
									app.Window().Call("alert", "Imported key")
								},

								OnEncryptAndSign: func() {
									app.Window().Call("alert", "Encrypted and signed")
								},
								OnDecryptAndVerify: func() {
									app.Window().Call("alert", "Decrypted and verified")
								},
							},
						),
					app.Section().
						Class("pf-c-page__main-section pf-m-fill").
						Body(
							&KeyList{
								Keys: demoKeys,
							},
						),
				),
		)
}

func (c *Home) OnAppUpdate(ctx app.Context) {
	if ctx.AppUpdateAvailable() {
		ctx.Reload()
	}
}
