# keygaen

![Logo](./assets/logo-readme.png)

Sign, verify, encrypt and decrypt data with GPG in your browser.

⚠️ keygaen has not yet been audited! While we try to make keygaen as secure as possible, it has not yet undergone a formal security audit by a third party. Please keep this in mind if you use it for security-critical applications. ⚠️

[![hydrun CI](https://github.com/pojntfx/keygaen/actions/workflows/hydrun.yaml/badge.svg)](https://github.com/pojntfx/keygaen/actions/workflows/hydrun.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/pojntfx/keygaen.svg)](https://pkg.go.dev/github.com/pojntfx/keygaen)
[![Matrix](https://img.shields.io/matrix/keygaen:matrix.org)](https://matrix.to/#/#keygaen:matrix.org?via=matrix.org)

## Installation

The web app is available on [GitHub releases](https://github.com/pojntfx/keygaen/releases) in the form of a static `.tar.gz` archive; to deploy it, simply upload it to a CDN or copy it to a web server. For most users, this shouldn't be necessary though; thanks to [@maxence-charriere](https://github.com/maxence-charriere)'s [go-app package](https://go-app.dev/), keygaen is a progressive web app. By simply visiting the [public deployment](https://pojntfx.github.io/keygaen/) once, it will be available for offline use whenever you need it:

[<img src="https://github.com/alphahorizonio/webnetesctl/raw/main/img/launch.png" width="240">](https://pojntfx.github.io/keygaen/)

## Screenshots

Click on an image to see a larger version.

<a display="inline" href="./assets/empty.png?raw=true">
<img src="./assets/empty.png" width="45%" alt="Screenshot of the empty start screen" title="Screenshot of the empty start screen">
</a>

<a display="inline" href="./assets/key-create.png?raw=true">
<img src="./assets/key-create.png" width="45%" alt="Screenshot of the key creation modal" title="Screenshot of the key creation modal">
</a>

<a display="inline" href="./assets/key-import.png?raw=true">
<img src="./assets/key-import.png" width="45%" alt="Screenshot of the key import modal" title="Screenshot of the key import modal">
</a>

<a display="inline" href="./assets/key-list.png?raw=true">
<img src="./assets/key-list.png" width="45%" alt="Screenshot of the key list" title="Screenshot of the key list">
</a>

<a display="inline" href="./assets/encrypt-sign.png?raw=true">
<img src="./assets/encrypt-sign.png" width="45%" alt="Screenshot of the encrypt/sign modal" title="Screenshot of the encrypt/sign modal">
</a>

<a display="inline" href="./assets/view-cypher.png?raw=true">
<img src="./assets/view-cypher.png" width="45%" alt="Screenshot of the viewing the cypher" title="Screenshot of the viewing the cypher">
</a>

<a display="inline" href="./assets/download-cypher.png?raw=true">
<img src="./assets/download-cypher.png" width="45%" alt="Screenshot of the downloading the cypher" title="Screenshot of the downloading the cypher">
</a>

<a display="inline" href="./assets/decrypt-verify.png?raw=true">
<img src="./assets/decrypt-verify.png" width="45%" alt="Screenshot of the decrypt/verify modal" title="Screenshot of the decrypt/verify modal">
</a>

<a display="inline" href="./assets/view-plaintext.png?raw=true">
<img src="./assets/view-plaintext.png" width="45%" alt="Screenshot of the viewing the plaintext" title="Screenshot of the viewing the plaintext">
</a>

<a display="inline" href="./assets/download-plaintext.png?raw=true">
<img src="./assets/download-plaintext.png" width="45%" alt="Screenshot of the downloading the plaintext" title="Screenshot of the downloading the plaintext">
</a>

<a display="inline" href="./assets/export-key.png?raw=true">
<img src="./assets/export-key.png" width="45%" alt="Screenshot of the export key modal" title="Screenshot of the export key modal">
</a>

## Acknowledgements

- This project would not have been possible were it not for [@maxence-charriere](https://github.com/maxence-charriere)'s [go-app package](https://go-app.dev/); if you enjoy using keygaen, please donate to him!
- The open source [PatternFly design system](https://www.patternfly.org/v4/) provides the components for the project.
- [GopenPGP](https://gopenpgp.org/) is the GPG library in use.
- All the rest of the authors who worked on the dependencies used! Thanks a lot!

## Contributing

To contribute, please use the [GitHub flow](https://guides.github.com/introduction/flow/) and follow our [Code of Conduct](./CODE_OF_CONDUCT.md).

To build and start a development version of keygaen locally, run the following:

```shell
$ git clone https://github.com/pojntfx/keygaen.git
$ cd keygaen
$ make run
```

Have any questions or need help? Chat with us [on Matrix](https://matrix.to/#/#keygaen:matrix.org?via=matrix.org)!

## License

keygaen (c) 2021 Felix Pojtinger and contributors

SPDX-License-Identifier: AGPL-3.0
