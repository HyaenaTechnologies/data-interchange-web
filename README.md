[CommonMark]:https://commonmark.org/
[Dart]: https://dart.dev/
[Go Language]: https://go.dev
[HTTP]: https://developer.mozilla.org/en-US/docs/Web/HTTP
[JavaScript Language]: https://developer.mozilla.org/en-US/docs/Web/JavaScript
[JSON]: https://www.json.org/json-en.html
[MDN]: https://developer.mozilla.org/en-US/docs/Web/API
[TOML]: https://toml.io/en/
[VSCode]: https://code.visualstudio.com/docs
[YAML]: https://yaml.org/
[WAI-ARIA]: https://www.w3.org/WAI/ARIA/apg/patterns/
[Webstorm]: https://jetbrains.com/webstorm

<a href="https://github.com/HyaenaTechnologies/data-interchange-web">
  <h1>
    <picture>
      <img src="https://github.com/HyaenaTechnologies/data-interchange-web/blob/main/assets/di_markdown.png" alt="">
    </picture>
  </h1>
</a>

[![Go Workflow](https://github.com/HyaenaTechnologies/data-interchange-web/actions/workflows/go.yml/badge.svg)](https://github.com/HyaenaTechnologies/data-interchange-web/actions/workflows/go.yml)

# Data Interchange Web

Data Interchange is a Network Application Programming Interface Development Platform

## Features

- Comma-Separated Values
- Extensible Markup Language
- [Hypertext Transfer Protocol][HTTP]
- Internet Protocol
- [JavaScript Object Notation][JSON]
- [Markdown][CommonMark]
- [Tom's Obvious Minimal Language][TOML]
- Transmission Control Protocol
- Transport Layer Security
- User Datagram Protocol
- Web Socket Protocol
- [YAML Ain't Markup Language][YAML]

## Build

- [Go][Go Language]
- [JetBrains Webstorm][Webstorm]
- [JavaScript][JavaScript Language]
- [Mozilla Developer Network Web Documentation][MDN]
- [Visual Studio Code][VSCode]
- [Web Accessibility Initiative][WAI-ARIA]

## Build Web Server

```shell
git clone

GOOS=linux GOARCH=amd64 go build -o ./binary/htdinet ./server/source/main.go
```

## Install Web Server

```shell
echo 'export PATH="$PATH:/usr/local/bin/htdinet"' >> ~/.bashrc && sudo echo 'export PATH="$PATH:/usr/local/bin/htdinet"' >> /etc/skel/.bashrc
```