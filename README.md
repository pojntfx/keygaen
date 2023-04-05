# keygaen

![Logo](./docs/logo-readme.png)

Sign, verify, encrypt and decrypt data with PGP in your browser.

⚠️ keygaen has not yet been audited! While we try to make keygaen as secure as possible, it has not yet undergone a formal security audit by a third party. Please keep this in mind if you use it for security-critical applications. ⚠️

[![hydrun CI](https://github.com/pojntfx/keygaen/actions/workflows/hydrun.yaml/badge.svg)](https://github.com/pojntfx/keygaen/actions/workflows/hydrun.yaml)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.18-61CFDD.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/pojntfx/keygaen.svg)](https://pkg.go.dev/github.com/pojntfx/keygaen)
[![Matrix](https://img.shields.io/matrix/keygaen:matrix.org)](https://matrix.to/#/#keygaen:matrix.org?via=matrix.org)
[![Binary Downloads](https://img.shields.io/github/downloads/pojntfx/keygaen/total?label=binary%20downloads)](https://github.com/pojntfx/keygaen/releases)

## Overview

keygaen is an app to work with PGP without having to install anything on your local system.

## Installation

The web app is available on [GitHub releases](https://github.com/pojntfx/keygaen/releases) in the form of a static `.tar.gz` archive; to deploy it, simply upload it to a CDN or copy it to a web server. For most users, this shouldn't be necessary though; thanks to [@maxence-charriere](https://github.com/maxence-charriere)'s [go-app package](https://go-app.dev/), keygaen is a progressive web app. By simply visiting the [public deployment](https://pojntfx.github.io/keygaen/) once, it will be available for offline use whenever you need it:

[<img src="https://github.com/alphahorizonio/webnetesctl/raw/main/img/launch.png" width="240">](https://pojntfx.github.io/keygaen/)

## Screenshots

Click on an image to see a larger version.

<a display="inline" href="./docs/empty.png?raw=true">
<img src="./docs/empty.png" width="45%" alt="Screenshot of the empty start screen" title="Screenshot of the empty start screen">
</a>

<a display="inline" href="./docs/key-create.png?raw=true">
<img src="./docs/key-create.png" width="45%" alt="Screenshot of the key creation modal" title="Screenshot of the key creation modal">
</a>

<a display="inline" href="./docs/key-import.png?raw=true">
<img src="./docs/key-import.png" width="45%" alt="Screenshot of the key import modal" title="Screenshot of the key import modal">
</a>

<a display="inline" href="./docs/key-list.png?raw=true">
<img src="./docs/key-list.png" width="45%" alt="Screenshot of the key list" title="Screenshot of the key list">
</a>

<a display="inline" href="./docs/encrypt-sign.png?raw=true">
<img src="./docs/encrypt-sign.png" width="45%" alt="Screenshot of the encrypt/sign modal" title="Screenshot of the encrypt/sign modal">
</a>

<a display="inline" href="./docs/view-cypher.png?raw=true">
<img src="./docs/view-cypher.png" width="45%" alt="Screenshot of the viewing the cypher" title="Screenshot of the viewing the cypher">
</a>

<a display="inline" href="./docs/download-cypher.png?raw=true">
<img src="./docs/download-cypher.png" width="45%" alt="Screenshot of the downloading the cypher" title="Screenshot of the downloading the cypher">
</a>

<a display="inline" href="./docs/decrypt-verify.png?raw=true">
<img src="./docs/decrypt-verify.png" width="45%" alt="Screenshot of the decrypt/verify modal" title="Screenshot of the decrypt/verify modal">
</a>

<a display="inline" href="./docs/view-plaintext.png?raw=true">
<img src="./docs/view-plaintext.png" width="45%" alt="Screenshot of the viewing the plaintext" title="Screenshot of the viewing the plaintext">
</a>

<a display="inline" href="./docs/download-plaintext.png?raw=true">
<img src="./docs/download-plaintext.png" width="45%" alt="Screenshot of the downloading the plaintext" title="Screenshot of the downloading the plaintext">
</a>

<a display="inline" href="./docs/export-key.png?raw=true">
<img src="./docs/export-key.png" width="45%" alt="Screenshot of the export key modal" title="Screenshot of the export key modal">
</a>

## Acknowledgements

- This project would not have been possible were it not for [@maxence-charriere](https://github.com/maxence-charriere)'s [go-app package](https://go-app.dev/); if you enjoy using keygaen, please donate to him!
- The open source [PatternFly design system](https://www.patternfly.org/v4/) provides the components for the project.
- [GopenPGP](https://gopenpgp.org/) is the PGP library in use.

## Contributing

To contribute, please use the [GitHub flow](https://guides.github.com/introduction/flow/) and follow our [Code of Conduct](./CODE_OF_CONDUCT.md).

To build and start a development version of keygaen locally, run the following:

```shell
$ git clone https://github.com/pojntfx/keygaen.git
$ cd keygaen
$ make depend
$ make run-pwa/keygaen-pwa
```

Have any questions or need help? Chat with us [on Matrix](https://matrix.to/#/#keygaen:matrix.org?via=matrix.org)!

## License

keygaen (c) 2023 Felicitas Pojtinger and contributors

SPDX-License-Identifier: AGPL-3.0
