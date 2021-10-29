package components

import (
	"errors"
	"log"

	"github.com/ProtonMail/go-crypto/openpgp"
	"github.com/ProtonMail/gopenpgp/armor"
	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Home struct {
	app.Compo

	createKeyModalOpen                bool
	importKeyModal                    bool
	encryptAndSignModalOpen           bool
	decryptAndVerifyModalOpen         bool
	keySuccessfullyGeneratedModalOpen bool

	keyImportPasswordModalOpen       bool
	keySuccessfullyImportedModalOpen bool

	publicKeyID             string
	privateKeyID            string
	createDetachedSignature bool

	encryptAndSignPasswordModalOpen bool
	encryptAndSignDownloadModalOpen bool

	confirmCloseModalOpen bool
	confirmModalClose     func()

	viewCypherAndSignatureModalOpen bool

	decryptAndVerifyPasswordModalOpen bool
	decryptAndVerifyDownloadModalOpen bool

	viewPlaintextModalOpen bool

	selectedKeyID             string
	deleteKeyConfirmModalOpen bool

	exportKeyModalOpen bool
	viewKeyModalOpen   bool

	viewPrivateKey bool

	err       error
	onRecover func()

	keys []GPGKey

	keyPasswordChan chan string

	wrongPassword bool

	keyDuplicateModalOpen bool
}

func (c *Home) Render() app.UI {
	if c.keys == nil {
		c.keys = []GPGKey{}
	}

	if c.keyPasswordChan == nil {
		c.keyPasswordChan = make(chan string)
	}

	privateKey := GPGKey{}
	privateKeyExport := []byte{}
	privateKeyExportArmored := ""
	for _, candidate := range c.keys {
		if candidate.ID == c.privateKeyID {
			privateKey = candidate

			parsedKey, err := crypto.NewKeyFromArmored(privateKey.Content)
			if err != nil {
				c.panic(err, func() {})

				break
			}

			privateKeyExport, err = parsedKey.Serialize()
			if err != nil {
				c.panic(err, func() {})

				break
			}

			privateKeyExportArmored, err = parsedKey.Armor()
			if err != nil {
				c.panic(err, func() {})

				break
			}

			break
		}
	}

	publicKey := GPGKey{}
	publicKeyExport := []byte{}
	publicKeyExportArmored := ""
	for _, candidate := range c.keys {
		if candidate.ID == c.publicKeyID {
			publicKey = candidate

			parsedKey, err := crypto.NewKeyFromArmored(publicKey.Content)
			if err != nil {
				c.panic(err, func() {})

				break
			}

			publicKeyExport, err = parsedKey.GetPublicKey()
			if err != nil {
				c.panic(err, func() {})

				break
			}

			publicKeyExportArmored, err = parsedKey.GetArmoredPublicKey()
			if err != nil {
				c.panic(err, func() {})

				break
			}

			break
		}
	}

	selectedKey := GPGKey{}
	for _, candidate := range c.keys {
		if candidate.ID == c.selectedKeyID {
			selectedKey = candidate

			break
		}
	}

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
			app.If(
				c.keySuccessfullyGeneratedModalOpen,
				&SingleActionModal{
					ID:          "key-successfully-generated-modal",
					Icon:        "fas fa-check",
					Title:       "Key Successfully Generated!",
					Class:       "pf-m-success",
					Body:        "It has been added to the key list.",
					ActionLabel: "Continue to key list",

					OnClose: func() {
						c.keySuccessfullyGeneratedModalOpen = false

						c.Update()
					},
					OnAction: func() {
						c.keySuccessfullyGeneratedModalOpen = false
					},
				},
			),
			app.If(
				c.keyImportPasswordModalOpen,
				&PasswordModal{
					Title:         "Enter Key Password",
					WrongPassword: c.wrongPassword,
					ClearWrongPassword: func() {
						c.wrongPassword = false
					},
					OnSubmit: func(password string) {
						c.keyPasswordChan <- password
					},
					OnCancel: func() {
						c.confirmModalClose = func() {
							c.keyPasswordChan <- ""
						}
						c.confirmCloseModalOpen = true

						c.Update()
					},
				},
			),
			app.If(
				c.keySuccessfullyImportedModalOpen,
				&SingleActionModal{
					ID:          "key-successfully-imported-modal",
					Icon:        "fas fa-check",
					Title:       "Key Successfully Imported!",
					Class:       "pf-m-success",
					Body:        "It has been added to the key list.",
					ActionLabel: "Continue to key list",

					OnClose: func() {
						c.keySuccessfullyImportedModalOpen = false

						c.Update()
					},
					OnAction: func() {
						c.keySuccessfullyImportedModalOpen = false
					},
				},
			),
			app.If(
				c.keyDuplicateModalOpen,
				&SingleActionModal{
					ID:          "key-successfully-imported-modal",
					Icon:        "fas fa-info-circle",
					Title:       "Key Already Exists",
					Class:       "pf-m-info",
					Body:        "This key is already in the key list.",
					ActionLabel: "Continue to key list",

					OnClose: func() {
						c.keyDuplicateModalOpen = false

						c.Update()
					},
					OnAction: func() {
						c.keyDuplicateModalOpen = false
					},
				},
			),
			app.If(
				c.encryptAndSignPasswordModalOpen,
				&PasswordModal{
					Title: `Enter password for key "` + c.privateKeyID + `"`,
					OnSubmit: func(password string) {
						c.encryptAndSignPasswordModalOpen = false
						c.encryptAndSignDownloadModalOpen = true
					},
					OnCancel: func() {
						c.confirmModalClose = func() {
							c.encryptAndSignPasswordModalOpen = false
						}
						c.confirmCloseModalOpen = true

						c.Update()
					},
				},
			),
			app.If(
				c.decryptAndVerifyPasswordModalOpen,
				&PasswordModal{
					Title: `Enter password for key "` + c.privateKeyID + `"`,
					OnSubmit: func(password string) {
						c.decryptAndVerifyPasswordModalOpen = false
						c.decryptAndVerifyDownloadModalOpen = true
					},
					OnCancel: func() {
						c.confirmModalClose = func() {
							c.decryptAndVerifyPasswordModalOpen = false
						}
						c.confirmCloseModalOpen = true

						c.Update()
					},
				},
			),
			app.If(
				c.encryptAndSignDownloadModalOpen,
				&DownloadOrViewModal{
					SubjectA:     c.privateKeyID != "",
					SubjectANoun: "signature",
					SubjectAVerb: "signed",

					SubjectB:     c.publicKeyID != "",
					SubjectBNoun: "cypher",
					SubjectBVerb: "encrypted",

					OnClose: func(used bool) {
						if used {
							c.encryptAndSignDownloadModalOpen = false

							c.Update()

							return
						}

						c.confirmModalClose = func() {
							c.encryptAndSignDownloadModalOpen = false
						}
						c.confirmCloseModalOpen = true

						c.Update()
					},
					OnDownload: func() {
						c.download([]byte("Hello, world"), "cypher.txt", "text/plain")

						if c.createDetachedSignature {
							c.download([]byte("asdf"), "signature.asc", "text/plain")
						}
					},
					OnView: func() {
						c.encryptAndSignDownloadModalOpen = false
						c.viewCypherAndSignatureModalOpen = true
					},
				},
			),
			app.If(
				c.decryptAndVerifyDownloadModalOpen,
				&DownloadOrViewModal{
					SubjectA:     c.privateKeyID != "",
					SubjectANoun: "file",
					SubjectAVerb: "decrypted",

					SubjectB:     c.publicKeyID != "",
					SubjectBNoun: "",
					SubjectBVerb: "verified",

					OnClose: func(used bool) {
						if used {
							c.decryptAndVerifyDownloadModalOpen = false

							c.Update()

							return
						}

						c.confirmModalClose = func() {
							c.decryptAndVerifyDownloadModalOpen = false
						}
						c.confirmCloseModalOpen = true

						c.Update()
					},
					OnDownload: func() {
						c.download([]byte("Hello, world"), "plaintext.txt", "text/plain")
					},
					OnView: func() {
						c.decryptAndVerifyDownloadModalOpen = false
						c.viewPlaintextModalOpen = true
					},
				},
			),
			app.Main().
				Class("pf-c-page__main").
				ID("gridge-main").
				TabIndex(-1).
				Body(
					app.If(
						len(c.keys) != 0,
						app.Section().
							Class("pf-c-page__main-section pf-m-light pf-m-no-padding pf-u-px-sm-on-xl").
							Body(
								&Toolbar{
									OnCreateKey: func() {
										c.createKeyModalOpen = !c.createKeyModalOpen
									},
									OnImportKey: func() {
										c.importKeyModal = !c.importKeyModal
									},

									OnEncryptAndSign: func() {
										c.encryptAndSignModalOpen = !c.encryptAndSignModalOpen
									},
									OnDecryptAndVerify: func() {
										c.decryptAndVerifyModalOpen = !c.decryptAndVerifyModalOpen
									},
								},
							),
					),
					app.Section().
						Class("pf-c-page__main-section pf-m-fill").
						Body(
							app.If(
								len(c.keys) == 0,
								&EmptyState{
									OnCreateKey: func() {
										c.createKeyModalOpen = !c.createKeyModalOpen
									},
									OnImportKey: func() {
										c.importKeyModal = !c.importKeyModal
									},
								},
							).Else(
								&KeyList{
									Keys: c.keys,

									OnExport: func(keyID string) {
										c.selectedKeyID = keyID
										c.exportKeyModalOpen = !c.exportKeyModalOpen

										for _, key := range c.keys {
											if key.ID == keyID {
												if key.Public {
													c.publicKeyID = key.ID
												} else {
													c.publicKeyID = ""
												}

												if key.Private {
													c.privateKeyID = key.ID
												} else {
													c.privateKeyID = ""
												}

												break
											}
										}
									},
									OnDelete: func(keyID string) {
										c.selectedKeyID = keyID
										c.deleteKeyConfirmModalOpen = !c.deleteKeyConfirmModalOpen
									},
								},
							),
						),
				),
			app.If(
				c.createKeyModalOpen,
				&CreateKeyModal{
					OnSubmit: func(fullName, email, password string) {
						key, err := helper.GenerateKey(fullName, email, []byte(password), "x25519", 0)
						if err != nil {
							c.createKeyModalOpen = false
							c.panic(err, func() {
								c.createKeyModalOpen = true
							})

							return
						}

						c.createKeyModalOpen = false
						c.keySuccessfullyGeneratedModalOpen = true

						fingerprints, err := helper.GetSHA256Fingerprints(key)
						if err != nil {
							c.createKeyModalOpen = false
							c.panic(err, func() {
								c.createKeyModalOpen = true
							})

							return
						}

						c.keys = append(c.keys, GPGKey{
							ID:       fingerprints[0],       // Since we don't generate subkeys, we'll only have one fingerprint
							Label:    fingerprints[0][0:10], // We can safely assume that the fingerprint is at least 10 chars long
							FullName: fullName,
							Email:    email,
							Private:  true,
							Public:   true,
							Content:  key,
						})
					},
					OnCancel: func(dirty bool, clear chan struct{}) {
						c.handleCancel(dirty, clear, func() {
							c.createKeyModalOpen = false
						})
					},
				},
			),
			app.If(
				c.importKeyModal,
				&ImportKeyModal{
					OnSubmit: func(key string) {
						c.importKeyModal = false

						parsedKey, err := crypto.NewKeyFromArmored(key)
						if err != nil {
							c.panic(err, func() {})

							return
						}

						go func() {
							// We might have to unlock a private key first
							if parsedKey.IsPrivate() {
								locked, err := parsedKey.IsLocked()
								if err != nil {
									c.panic(err, func() {})

									return
								}

								if locked {
									c.keyImportPasswordModalOpen = true

									c.Update()

									for {
										password := <-c.keyPasswordChan

										// Stop import if no password has been entered
										if password == "" {
											c.keyImportPasswordModalOpen = false

											return
										}

										newParsedKey, err := parsedKey.Unlock([]byte(password))
										if err != nil {
											c.handleWrongPassword(err)

											continue
										}

										parsedKey = newParsedKey
										c.keyImportPasswordModalOpen = false
										c.clearWrongPassword()

										c.Update()

										break
									}
								}
							}

							parsedPublicKey, err := parsedKey.GetArmoredPublicKey()
							if err != nil {
								c.panic(err, func() {})

								return
							}

							fingerprints, err := helper.GetSHA256Fingerprints(parsedPublicKey)
							if err != nil {
								c.createKeyModalOpen = false
								c.panic(err, func() {
									c.createKeyModalOpen = true
								})

								return
							}

							var id *openpgp.Identity
							for _, candidate := range parsedKey.GetEntity().Identities {
								id = candidate

								break
							}

							if id == nil {
								c.panic(errors.New("no identity found in key"), func() {})

								return
							}

							newKeys := []GPGKey{}
							for _, candidate := range c.keys {
								if candidate.ID == fingerprints[0] {
									// Replace the key if the existing key is a public key and the imported key is a private key
									if !candidate.Private && parsedKey.IsPrivate() {
										continue
									} else {
										// Don't add the duplicate key
										c.keyDuplicateModalOpen = true

										c.Update()

										return
									}
								}

								newKeys = append(newKeys, candidate)
							}

							newKeys = append(newKeys, GPGKey{
								ID:       fingerprints[0],       // Since we don't generate subkeys, we'll only have one fingerprint
								Label:    fingerprints[0][0:10], // We can safely assume that the fingerprint is at least 10 chars long
								FullName: id.Name,
								Email:    id.UserId.Email,
								Private:  parsedKey.IsPrivate(),
								Public:   true,
								Content:  key,
							})
							c.keys = newKeys

							c.keySuccessfullyImportedModalOpen = true

							c.Update()
						}()
					},
					OnCancel: func(dirty bool, clear chan struct{}) {
						c.handleCancel(dirty, clear, func() {
							c.importKeyModal = false
						})
					},
				},
			),
			app.If(
				c.encryptAndSignModalOpen,
				&EncryptAndSignModal{
					Keys: c.keys,

					OnSubmit: func(file []byte, publicKeyID, privateKeyID string, createDetachedSignature, enableArmor bool) {
						c.publicKeyID = publicKeyID
						c.privateKeyID = privateKeyID
						c.createDetachedSignature = createDetachedSignature

						c.encryptAndSignModalOpen = false
						c.encryptAndSignPasswordModalOpen = true

						// TODO: Only use if c.publicKeyID != ""
						publicKey, err := c.getPublicKeyByID(c.publicKeyID)
						if err != nil {
							c.panic(err, func() {})

							return
						}

						armoredCyphertext, err := helper.EncryptBinaryMessageArmored(publicKey, file)
						if err != nil {
							c.panic(err, func() {})

							return
						}

						if enableArmor {
							log.Println(armoredCyphertext)

							return
						}

						rawCyphertext, err := armor.Unarmor(armoredCyphertext)
						if err != nil {
							c.panic(err, func() {})

							return
						}

						log.Println(rawCyphertext)

						// TODO: Add signing based on private key; requires unlocking with modal
					},
					OnCancel: func(dirty bool, clear chan struct{}) {
						c.handleCancel(dirty, clear, func() {
							c.encryptAndSignModalOpen = false
						})
					},
				},
			),
			app.If(
				c.decryptAndVerifyModalOpen,
				&DecryptAndVerifyModal{
					Keys: c.keys,

					OnSubmit: func(file []byte, publicKeyID, privateKeyID, detachedSignature string) {
						c.publicKeyID = publicKeyID
						c.privateKeyID = privateKeyID

						c.decryptAndVerifyModalOpen = false
						c.decryptAndVerifyPasswordModalOpen = true
					},
					OnCancel: func(dirty bool, clear chan struct{}) {
						c.handleCancel(dirty, clear, func() {
							c.decryptAndVerifyModalOpen = false
						})
					},
				},
			),
			app.If(
				c.viewCypherAndSignatureModalOpen,
				func() app.UI {
					tabs := []TextOutputModalTab{
						{
							Language: "text/plain",
							Title:    "cypher.txt",
							Body:     "Hello, world",
						},
					}
					title := "View Cypher"

					if c.createDetachedSignature {
						tabs = append(
							tabs,
							TextOutputModalTab{
								Language: "text/plain",
								Title:    "signature.asc",
								Body:     "uas-02rioj23jd",
							},
						)
						title += " and Signature"
					}

					return &TextOutputModal{
						Title: title,
						Tabs:  tabs,
						OnClose: func() {
							c.viewCypherAndSignatureModalOpen = false
							c.encryptAndSignDownloadModalOpen = true

							c.Update()
						},
					}
				}(),
			),
			app.If(
				c.viewPlaintextModalOpen,
				func() app.UI {
					return &TextOutputModal{
						Title: "View File",
						Tabs: []TextOutputModalTab{
							{
								Language: "text/plain",
								Title:    "plaintext.txt",
								Body:     "Hello, world",
							},
						},
						OnClose: func() {
							c.viewPlaintextModalOpen = false
							c.decryptAndVerifyDownloadModalOpen = true

							c.Update()
						},
					}
				}(),
			),
			app.If(
				c.exportKeyModalOpen,
				&ExportKeyModal{
					PublicKey: c.publicKeyID != "",
					OnDownloadPublicKey: func(armor bool) {
						if armor {
							c.download([]byte(publicKeyExportArmored), publicKey.Label+".pub", "text/plain")
						} else {
							c.download(publicKeyExport, publicKey.Label+".pub", "application/octet-stream")
						}
					},
					OnViewPublicKey: func() {
						c.exportKeyModalOpen = false
						c.viewPrivateKey = false
						c.viewKeyModalOpen = true
					},

					PrivateKey: c.privateKeyID != "",
					OnDownloadPrivateKey: func(armor bool) {
						if armor {
							c.download([]byte(privateKeyExportArmored), privateKey.Label, "text/plain")
						} else {
							c.download(privateKeyExport, privateKey.Label, "application/octet-stream")
						}
					},
					OnViewPrivateKey: func() {
						c.exportKeyModalOpen = false
						c.viewPrivateKey = true
						c.viewKeyModalOpen = true
					},

					OnOK: func() {
						c.exportKeyModalOpen = false

						c.Update()
					},
				},
			),
			app.If(
				c.viewKeyModalOpen,
				func() app.UI {
					tabs := []TextOutputModalTab{
						{
							Language: "text/plain",
							Title:    publicKey.Label + ".pub",
							Body:     publicKeyExportArmored,
						},
					}
					title := `View Public Key "` + publicKey.Label + `"`

					if c.viewPrivateKey {
						tabs = []TextOutputModalTab{
							{
								Language: "text/plain",
								Title:    privateKey.Label,
								Body:     privateKeyExportArmored,
							},
						}
						title = `View Private Key "` + privateKey.Label + `"`
					}

					return &TextOutputModal{
						Title: title,
						Tabs:  tabs,
						OnClose: func() {
							c.viewKeyModalOpen = false
							c.exportKeyModalOpen = true

							c.Update()
						},
					}
				}(),
			),
			app.If(
				c.confirmCloseModalOpen,
				&ConfirmationModal{
					ID:    "confirmation-modal",
					Icon:  "fas fa-exclamation-triangle",
					Title: "Are you sure?",
					Class: "pf-m-danger",
					Body:  "Unsaved changes might be lost.",

					ActionLabel: "Yes, delete unsaved changes",
					ActionClass: "pf-m-danger",

					CancelLabel: "Cancel",

					OnClose: func() {
						c.confirmCloseModalOpen = false

						c.Update()
					},
					OnAction: func() {
						c.confirmCloseModalOpen = false

						c.confirmModalClose()

						c.Update()
					},
				},
			),
			app.If(
				c.deleteKeyConfirmModalOpen,
				&ConfirmationModal{
					ID:    "delete-key-confirmation-modal",
					Icon:  "fas fa-exclamation-triangle",
					Title: "Are you sure?",
					Class: "pf-m-danger",
					Body:  "After deletion, you will not be able to restore the key.",

					ActionLabel: `Yes, delete key "` + selectedKey.Label + `"`,
					ActionClass: "pf-m-danger",

					CancelLabel: "Cancel",

					OnClose: func() {
						c.deleteKeyConfirmModalOpen = false

						c.Update()
					},
					OnAction: func() {
						c.deleteKeyConfirmModalOpen = false

						c.Update()
					},
				},
			),
			app.If(
				c.err != nil,
				&ErrorModal{
					ID:          "error-modal",
					Icon:        "fas fa-times",
					Title:       "An Error Occurred",
					Class:       "pf-m-danger",
					Body:        "The following details may be of help:",
					Error:       c.err,
					ActionLabel: "Close",

					OnClose: func() {
						c.recover()
					},
					OnAction: func() {
						c.recover()
					},
				},
			),
		)
}

func (c *Home) download(content []byte, name string, mimetype string) {
	buf := app.Window().JSValue().Get("Uint8Array").New(len(content))
	app.CopyBytesToJS(buf, content)

	blob := app.Window().JSValue().Get("Blob").New(app.Window().JSValue().Get("Array").New(buf), map[string]interface{}{
		"type": mimetype,
	})

	link := app.Window().Get("document").Call("createElement", "a")
	link.Set("href", app.Window().Get("URL").Call("createObjectURL", blob))
	link.Set("download", name)
	link.Call("click")
}

func (c *Home) handleCancel(dirty bool, clear chan struct{}, confirm func()) {
	if !dirty {
		confirm()

		clear <- struct{}{}

		c.Update()

		return
	}

	c.confirmModalClose = func() {
		confirm()

		// Remove the native confirmation prompt
		app.Window().Set("onbeforeunload", app.Undefined())

		clear <- struct{}{}
	}

	// Add the native confirmation prompt
	app.Window().Set(
		"onbeforeunload",
		app.FuncOf(func(this app.Value, args []app.Value) interface{} {
			args[0].Set("returnValue", "")

			return nil
		}),
	)
	c.confirmCloseModalOpen = true

	c.Update()
}

func (c *Home) panic(err error, onRecover func()) {
	log.Println(err)

	c.onRecover = onRecover
	c.err = err

	c.Update()
}

func (c *Home) recover() {
	c.err = nil
	c.onRecover()
}

func (c *Home) handleWrongPassword(err error) {
	log.Println(err)

	c.wrongPassword = true

	c.Update()
}

func (c *Home) clearWrongPassword() {
	c.wrongPassword = false
}

func (c *Home) OnAppUpdate(ctx app.Context) {
	if ctx.AppUpdateAvailable() {
		ctx.Reload()
	}
}

func (c *Home) getPublicKeyByID(publicKeyID string) (string, error) {
	publicKeyExportArmored := ""
	for _, candidate := range c.keys {
		if candidate.ID == c.publicKeyID {
			publicKey := candidate

			parsedKey, err := crypto.NewKeyFromArmored(publicKey.Content)
			if err != nil {
				return "", err
			}

			publicKeyExportArmored, err = parsedKey.GetArmoredPublicKey()
			if err != nil {
				return "", err
			}

			break
		}
	}

	return publicKeyExportArmored, nil
}
