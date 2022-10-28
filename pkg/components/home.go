package components

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"unicode/utf8"

	"github.com/ProtonMail/go-crypto/openpgp"
	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/keygaen/pkg/crypt"
)

const (
	keyringStorageKey = "keygaenKeys"
	auditStorageKey   = "keygaenAudit"
)

// Home is the home page
type Home struct {
	app.Compo

	createKeyModalOpen                bool
	importKeyModalOpen                bool
	encryptAndSignModalOpen           bool
	decryptAndVerifyModalOpen         bool
	keySuccessfullyGeneratedModalOpen bool

	keyPasswordModalOpen             bool
	keyPasswordModalKeyID            string
	keySuccessfullyImportedModalOpen bool

	publicKeyID             string
	privateKeyID            string
	createDetachedSignature bool

	encryptAndSignDownloadModalOpen bool

	confirmCloseModalOpen bool
	confirmModalClose     func()

	viewCypherAndSignatureModalOpen bool

	decryptAndVerifyDownloadModalOpen bool

	viewPlaintextModalOpen bool

	selectedKeyID             string
	deleteKeyConfirmModalOpen bool
	confirmDeleteKey          func()

	exportKeyModalOpen bool
	viewKeyModalOpen   bool

	viewPrivateKey bool

	viewArmor  bool
	viewBase64 bool

	err       error
	onRecover func()

	keys []PGPKey

	keyPasswordChan chan string

	wrongPassword bool

	keyDuplicateModalOpen bool

	outputCyphertext []byte
	outputSignature  []byte

	outputPlaintext []byte
	outputVerified  bool

	removeEventListeners []func()

	showAuditModal bool
}

func (c *Home) Render() app.UI {
	if c.keys == nil {
		c.readFromLocalStorage()
	}

	if c.keyPasswordChan == nil {
		c.keyPasswordChan = make(chan string)
	}

	privateKey := PGPKey{}
	privateKeyExport := []byte{}
	privateKeyExportBase64 := ""
	privateKeyExportArmored := ""
	privateKeyExportArmoredBase64 := ""
	for _, candidate := range c.keys {
		if candidate.ID == c.privateKeyID {
			privateKey = candidate

			rawKey, err := crypt.Unarmor(privateKey.Content)
			if err != nil {
				c.panic(err, func() {})

				break
			}

			privateKeyExportBase64 = base64.StdEncoding.EncodeToString(rawKey)

			parsedKey, err := crypto.NewKey(rawKey)
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

			privateKeyExportArmoredBase64 = base64.StdEncoding.EncodeToString(privateKeyExport)

			break
		}
	}

	publicKey := PGPKey{}
	publicKeyExport := []byte{}
	publicKeyExportBase64 := ""
	publicKeyExportArmored := ""
	publicKeyExportArmoredBase64 := ""
	for _, candidate := range c.keys {
		if candidate.ID == c.publicKeyID {
			publicKey = candidate

			rawKey, err := crypt.Unarmor(publicKey.Content)
			if err != nil {
				c.panic(err, func() {})

				break
			}

			publicKeyExportBase64 = base64.StdEncoding.EncodeToString(rawKey)

			parsedKey, err := crypto.NewKey(rawKey)
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

			publicKeyExportArmoredBase64 = base64.StdEncoding.EncodeToString(publicKeyExport)

			break
		}
	}

	selectedKey := PGPKey{}
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
				Href("#keygaen-main").
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
				c.keyPasswordModalOpen,
				&PasswordModal{
					Title:         `Enter Password for Key "` + c.keyPasswordModalKeyID + `"`,
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
				c.encryptAndSignDownloadModalOpen,
				&DownloadOrViewModal{
					SubjectA:          c.privateKeyID != "",
					SubjectANoun:      "signature",
					SubjectAAdjective: "signed",

					SubjectB:          c.publicKeyID != "",
					SubjectBNoun:      "cypher",
					SubjectBAdjective: "encrypted",

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
						cext, cmime := getEncryptedExtensionAndMIME(c.outputCyphertext, "cyphertext")
						c.download(c.outputCyphertext, cext, cmime)

						if c.createDetachedSignature {
							sext, smime := getEncryptedExtensionAndMIME(c.outputSignature, "signature")
							c.download(c.outputSignature, sext, smime)
						}
					},
					OnView: func() {
						c.encryptAndSignDownloadModalOpen = false
						c.viewCypherAndSignatureModalOpen = true
					},

					ShowView: utf8.Valid(c.outputCyphertext) && utf8.Valid(c.outputSignature),
				},
			),
			app.If(
				c.decryptAndVerifyDownloadModalOpen,
				&DownloadOrViewModal{
					SubjectA:          c.privateKeyID != "",
					SubjectANoun:      "file",
					SubjectAAdjective: "decrypted",

					SubjectB:          c.publicKeyID != "",
					SubjectBNoun:      "",
					SubjectBAdjective: "verified",

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
						pext, pmime := getDecryptedExtensionAndMIME(c.outputPlaintext, "plaintext")
						c.download(c.outputPlaintext, pext, pmime)
					},
					OnView: func() {
						c.decryptAndVerifyDownloadModalOpen = false
						c.viewPlaintextModalOpen = true
					},

					ShowView: utf8.Valid(c.outputPlaintext),
				},
			),
			app.Main().
				Class("pf-c-page__main").
				ID("keygaen-main").
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
										c.importKeyModalOpen = !c.importKeyModalOpen
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
										c.importKeyModalOpen = !c.importKeyModalOpen
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

										c.confirmDeleteKey = func() {
											newKeys := []PGPKey{}
											for _, candidate := range c.keys {
												if candidate.ID == c.selectedKeyID {
													continue
												}

												newKeys = append(newKeys, candidate)
											}

											c.keys = newKeys
											c.writeToLocalStorage()

											c.Update()
										}
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
						key, err := crypt.GenerateKey(fullName, email, password)
						if err != nil {
							c.createKeyModalOpen = false
							c.panic(err, func() {
								c.createKeyModalOpen = true
							})

							return
						}

						c.createKeyModalOpen = false
						c.keySuccessfullyGeneratedModalOpen = true

						parsedKey, _, err := crypt.ReadKey(key, password)
						if err != nil {
							c.panic(err, func() {})

							return
						}

						id := parsedKey.PrimaryIdentity()
						if id == nil {
							c.panic(errors.New("no identity found in key"), func() {})

							return
						}

						c.keys = append(c.keys, PGPKey{
							ID:       parsedKey.PrimaryKey.KeyIdString(),
							Label:    parsedKey.PrimaryKey.KeyIdShortString(),
							FullName: id.Name,
							Email:    id.UserId.Email,
							Private:  parsedKey.PrivateKey != nil,
							Public:   true,
							Content:  key,
						})
						c.writeToLocalStorage()
					},
					OnCancel: func(dirty bool, clear chan struct{}) {
						c.handleCancel(dirty, clear, func() {
							c.createKeyModalOpen = false
						})
					},
				},
			),
			app.If(
				c.importKeyModalOpen,
				&ImportKeyModal{
					OnSubmit: func(key []byte) {
						c.importKeyModalOpen = false

						go func() {
							var parsedKey *openpgp.Entity

							// We might have to unlock a private key first
							locked, fingerprint, err := crypt.IsKeyLocked(key)
							if err != nil {
								c.panic(err, func() {})

								return
							}

							if locked {
								password := c.getPasswordForKey(
									fingerprint,
									func(password string) error {
										pk, fp, err := crypt.ReadKey(key, password)
										if err == nil {
											parsedKey = pk
											fingerprint = fp
										}

										return err
									},
								)

								// Stop import if no password has been entered
								if password == "" {
									c.keyPasswordModalOpen = false

									return
								}
							}

							if parsedKey == nil {
								parsedKey, _, err = crypt.ReadKey(key, "")
								if err != nil {
									c.panic(err, func() {})

									return
								}
							}

							id := parsedKey.PrimaryIdentity()
							if id == nil {
								c.panic(errors.New("no identity found in key"), func() {})

								return
							}

							newKeys := []PGPKey{}
							for _, candidate := range c.keys {
								if candidate.ID == parsedKey.PrimaryKey.KeyIdString() {
									// Replace the key if the existing key is a public key and the imported key is a private key
									if !candidate.Private && parsedKey.PrivateKey != nil {
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

							newKeys = append(newKeys, PGPKey{
								ID:       parsedKey.PrimaryKey.KeyIdString(),
								Label:    parsedKey.PrimaryKey.KeyIdShortString(),
								FullName: id.Name,
								Email:    id.UserId.Email,
								Private:  parsedKey.PrivateKey != nil,
								Public:   true,
								Content:  key,
							})
							c.keys = newKeys
							c.writeToLocalStorage()

							c.keySuccessfullyImportedModalOpen = true

							c.Update()
						}()
					},
					OnCancel: func(dirty bool, clear chan struct{}) {
						c.handleCancel(dirty, clear, func() {
							c.importKeyModalOpen = false
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

						go func() {
							// Encrypt
							var encryptConfig *crypt.EncryptConfig
							if c.publicKeyID != "" {
								rawPublicKey, err := c.getPublicKeyByID(c.publicKeyID)
								if err != nil {
									c.panic(err, func() {})

									return
								}

								publicKey, _, err := crypt.ReadKey(rawPublicKey, "")
								if err != nil {
									c.panic(err, func() {})

									return
								}

								encryptConfig = &crypt.EncryptConfig{
									PublicKey:       publicKey,
									ArmorCyphertext: enableArmor,
								}
							}

							// Sign
							var signatureConfig *crypt.SignatureConfig
							if c.privateKeyID != "" {
								rawPrivateKey, err := c.getPrivateKeyByID(c.privateKeyID)
								if err != nil {
									c.panic(err, func() {})

									return
								}

								var privateKey *openpgp.Entity

								// We might have to unlock a private key first
								locked, fingerprint, err := crypt.IsKeyLocked(rawPrivateKey)
								if err != nil {
									c.panic(err, func() {})

									return
								}

								if locked {
									password := c.getPasswordForKey(
										fingerprint,
										func(password string) error {
											pk, fp, err := crypt.ReadKey(rawPrivateKey, password)
											if err == nil {
												privateKey = pk
												fingerprint = fp
											}

											return err
										},
									)

									// Stop import if no password has been entered
									if password == "" {
										c.keyPasswordModalOpen = false

										return
									}
								}

								if privateKey == nil {
									privateKey, fingerprint, err = crypt.ReadKey(rawPrivateKey, "")
									if err != nil {
										c.panic(err, func() {})

										return
									}
								}

								signatureConfig = &crypt.SignatureConfig{
									PrivateKey:      privateKey,
									ArmorSignature:  enableArmor,
									DetachSignature: createDetachedSignature,
								}
							}

							cyphertext, signature, err := crypt.EncryptSign(
								encryptConfig,
								signatureConfig,
								file,
							)
							if err != nil {
								c.panic(err, func() {})

								return
							}

							c.outputCyphertext = cyphertext
							c.outputSignature = signature

							c.encryptAndSignDownloadModalOpen = true

							c.Update()
						}()
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

					OnSubmit: func(file []byte, publicKeyID, privateKeyID string, detachedSignature []byte) {
						c.publicKeyID = publicKeyID
						c.privateKeyID = privateKeyID

						c.decryptAndVerifyModalOpen = false

						go func() {
							// Decrypt
							var decryptConfig *crypt.DecryptConfig
							if c.privateKeyID != "" {
								rawPrivateKey, err := c.getPrivateKeyByID(c.privateKeyID)
								if err != nil {
									c.panic(err, func() {})

									return
								}

								var privateKey *openpgp.Entity

								// We might have to unlock a private key first
								locked, fingerprint, err := crypt.IsKeyLocked(rawPrivateKey)
								if err != nil {
									c.panic(err, func() {})

									return
								}

								if locked {
									password := c.getPasswordForKey(
										fingerprint,
										func(password string) error {
											pk, fp, err := crypt.ReadKey(rawPrivateKey, password)
											if err == nil {
												privateKey = pk
												fingerprint = fp
											}

											return err
										},
									)

									// Stop import if no password has been entered
									if password == "" {
										c.keyPasswordModalOpen = false

										return
									}
								}

								if privateKey == nil {
									privateKey, fingerprint, err = crypt.ReadKey(rawPrivateKey, "")
									if err != nil {
										c.panic(err, func() {})

										return
									}
								}

								decryptConfig = &crypt.DecryptConfig{
									PrivateKey: privateKey,
								}
							}

							// Verify
							var verifyConfig *crypt.VerifyConfig
							if c.publicKeyID != "" {
								rawPublicKey, err := c.getPublicKeyByID(c.publicKeyID)
								if err != nil {
									c.panic(err, func() {})

									return
								}

								publicKey, _, err := crypt.ReadKey(rawPublicKey, "")
								if err != nil {
									c.panic(err, func() {})

									return
								}

								verifyConfig = &crypt.VerifyConfig{
									PublicKey:         publicKey,
									DetachedSignature: detachedSignature,
								}
							}

							plaintext, verified, err := crypt.DecryptVerify(
								decryptConfig,
								verifyConfig,
								file,
							)
							if err != nil {
								c.panic(err, func() {})

								return
							}

							c.outputPlaintext = plaintext
							c.outputVerified = verified

							c.decryptAndVerifyDownloadModalOpen = true

							c.Update()
						}()
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
							Title:    "cyphertext.asc",
							Body:     string(c.outputCyphertext),
						},
					}
					title := "View Cypher"

					if c.createDetachedSignature {
						tabs = append(
							tabs,
							TextOutputModalTab{
								Language: "text/plain",
								Title:    "signature.asc",
								Body:     string(c.outputSignature),
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
								Title:    "plaintext",
								Body:     string(c.outputPlaintext),
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
					OnDownloadPublicKey: func(armor, base64encode bool) {
						if armor {
							if base64encode {
								c.download([]byte(publicKeyExportArmoredBase64), publicKey.Label+".asc.txt", "text/plain")
							} else {
								c.download([]byte(publicKeyExportArmored), publicKey.Label+".asc", "text/plain")
							}
						} else {
							if base64encode {
								c.download([]byte(publicKeyExportBase64), publicKey.Label+".pgp.txt", "application/octet-stream")
							} else {
								c.download(publicKeyExport, publicKey.Label+".pgp", "application/octet-stream")
							}
						}
					},
					OnViewPublicKey: func(armor, base64encode bool) {
						c.exportKeyModalOpen = false
						c.viewPrivateKey = false
						c.viewKeyModalOpen = true
						c.viewArmor = armor
						c.viewBase64 = base64encode
					},

					PrivateKey: c.privateKeyID != "",
					OnDownloadPrivateKey: func(armor, base64encode bool) {
						if armor {
							if base64encode {
								c.download([]byte(privateKeyExportArmoredBase64), privateKey.Label+".asc.txt", "text/plain")
							} else {
								c.download([]byte(privateKeyExportArmored), privateKey.Label+".asc", "text/plain")
							}
						} else {
							if base64encode {
								c.download([]byte(privateKeyExportBase64), privateKey.Label+".pgp.txt", "application/octet-stream")
							} else {
								c.download(privateKeyExport, privateKey.Label+".pgp", "application/octet-stream")
							}
						}
					},
					OnViewPrivateKey: func(armor, base64encode bool) {
						c.exportKeyModalOpen = false
						c.viewPrivateKey = true
						c.viewKeyModalOpen = true
						c.viewArmor = armor
						c.viewBase64 = base64encode
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
							Title: publicKey.Label + func() string {
								if c.viewArmor {
									if c.viewBase64 {
										return ".asc.txt"
									}

									return ".asc"
								}

								return ".txt"
							}(),
							Body: func() string {
								if c.viewBase64 {
									return publicKeyExportArmoredBase64
								}

								return publicKeyExportArmored
							}(),
						},
					}
					title := `View Public Key "` + publicKey.Label + `"`

					if c.viewPrivateKey {
						tabs = []TextOutputModalTab{
							{
								Language: "text/plain",
								Title: privateKey.Label + ".asc" + func() string {
									if c.viewArmor {
										if c.viewBase64 {
											return ".asc.txt"
										}

										return ".asc"
									}

									return ".txt"
								}(),
								Body: func() string {
									if c.viewBase64 {
										return privateKeyExportArmoredBase64
									}

									return privateKeyExportArmored
								}(),
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
							c.viewBase64 = false

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

						c.confirmDeleteKey()

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
			app.If(
				c.showAuditModal,
				&ConfirmationModal{
					ID:    "audit-modal",
					Icon:  "fas fa-exclamation-triangle",
					Title: "keygaen has not yet been audited!",
					Class: "pf-m-warning",
					Body:  "While we try to make keygaen as secure as possible, it has not yet undergone a formal security audit by a third party. Please keep this in mind if you use it for security-critical applications.",

					ActionLabel: "Yes, I understand",
					ActionClass: "pf-m-warning",

					CancelLink:  "https://en.wikipedia.org/wiki/Information_security_audit",
					CancelLabel: "What is an audit?",

					OnClose: func() {
						c.showAuditModal = false
						c.writeToLocalStorage()

						c.Update()
					},
					OnAction: func() {
						c.showAuditModal = false
						c.writeToLocalStorage()

						c.Update()
					},
				},
			),
		)
}

func (c *Home) OnMount(ctx app.Context) {
	c.removeEventListeners = []func(){
		app.Window().AddEventListener("storage", func(ctx app.Context, e app.Event) { // This event only fires in other tabs; it does not lead to local race conditions with c.writeKeysToLocalStorage
			c.readFromLocalStorage()

			c.Update()
		}),
	}
}

func (c *Home) OnDismount() {
	if c.removeEventListeners != nil {
		for _, clearListener := range c.removeEventListeners {
			clearListener()
		}
	}
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

		clear <- struct{}{}
	}
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

func (c *Home) getPublicKeyByID(ID string) ([]byte, error) {
	for _, publicKey := range c.keys {
		if publicKey.ID == c.publicKeyID {
			rawKey, err := crypt.Unarmor(publicKey.Content)
			if err != nil {
				return []byte{}, err
			}

			parsedKey, err := crypto.NewKey(rawKey)
			if err != nil {
				return []byte{}, err
			}

			publicKeyExport, err := parsedKey.GetArmoredPublicKey()
			if err != nil {
				return []byte{}, err
			}

			return []byte(publicKeyExport), nil
		}
	}

	return []byte{}, errors.New("could not find public key")
}

func (c *Home) getPrivateKeyByID(ID string) ([]byte, error) {
	for _, privateKey := range c.keys {
		if privateKey.ID == c.privateKeyID {
			rawKey, err := crypt.Unarmor(privateKey.Content)
			if err != nil {
				return []byte{}, err
			}

			parsedKey, err := crypto.NewKey(rawKey)
			if err != nil {
				return []byte{}, err
			}

			privateKeyExport, err := parsedKey.Armor()
			if err != nil {
				return []byte{}, err
			}

			return []byte(privateKeyExport), nil
		}
	}

	return []byte{}, errors.New("could not find private key")
}

func (c *Home) getPasswordForKey(ID string, checkPassword func(password string) error) string {
	c.keyPasswordModalOpen = true
	c.keyPasswordModalKeyID = ID

	c.Update()

	for {
		password := <-c.keyPasswordChan

		// Stop if no password has been entered
		if password == "" {
			c.keyPasswordModalOpen = false

			return ""
		}

		if err := checkPassword(password); err != nil {
			c.handleWrongPassword(err)

			continue
		}

		c.keyPasswordModalOpen = false
		c.clearWrongPassword()

		c.Update()

		return password
	}
}

func (c *Home) readFromLocalStorage() {
	if showAuditModal := app.Window().Get("localStorage").Call("getItem", auditStorageKey); showAuditModal.IsNull() || showAuditModal.IsUndefined() || showAuditModal.String() == "true" {
		c.showAuditModal = true
	}

	marshalledKeys := app.Window().Get("localStorage").Call("getItem", keyringStorageKey).String()

	// Ignore errors in JSON parsing
	newKeys := []PGPKey{}
	_ = json.Unmarshal([]byte(marshalledKeys), &newKeys)

	if newKeys == nil {
		c.keys = []PGPKey{}

		return
	}

	c.keys = newKeys
}

func (c *Home) writeToLocalStorage() {
	app.Window().Get("localStorage").Call("setItem", auditStorageKey, c.showAuditModal)

	marshalledKeys, err := json.Marshal(c.keys)
	if err != nil {
		c.panic(err, func() {})

		return
	}

	app.Window().Get("localStorage").Call("setItem", keyringStorageKey, string(marshalledKeys))
}

func getEncryptedExtensionAndMIME(content []byte, filename string) (string, string) {
	if utf8.Valid(content) {
		return filename + ".asc", "text/plain"
	}

	return filename + ".pgp", "application/octet-stream"
}

func getDecryptedExtensionAndMIME(content []byte, filename string) (string, string) {
	if utf8.Valid(content) {
		return filename, "text/plain"
	}

	return filename, "application/octet-stream"
}
