[CommonMark]:https://commonmark.org/
[Dart]: https://dart.dev/
[HTTP]: https://developer.mozilla.org/en-US/docs/Web/HTTP
[IDEA]: https://jetbrains.com/idea/
[JSON]: https://www.json.org/json-en.html
[TOML]: https://toml.io/en/
[VSCode]: https://code.visualstudio.com/docs
[YAML]: https://yaml.org/

<a href="https://github.com/HyaenaTechnologies/data-interchange-web">
  <h1>
    <picture>
      <img src="https://github.com/HyaenaTechnologies/data-interchange-web/blob/main/assets/di_markdown.png" alt="">
    </picture>
  </h1>
</a>

[![Dart Workflow](https://github.com/HyaenaTechnologies/data-interchange-web/actions/workflows/android.yml/badge.svg)](https://github.com/HyaenaTechnologies/data-interchange-web/actions/workflows/android.yml)

# Data Interchange Web

Data Interchange is a Network Application Programming Interface Platform

Data Interchange Web is the Web implementation of Data Interchange

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

- [Dart Language][Dart]
- [IntelliJ IDEA][IDEA]
- [Visual Studio Code][VSCode]

```shell
git clone

dart pub upgrade

dart run build_runner build

dart run build_runner serve

dart run build_runner test

dart run build_runner watch
```

OR

```shell
git clone

dart pub upgrade

dart pub global activate webdev

webdev build

webdev serve
```
