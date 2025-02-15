[CommonMark]:https://commonmark.org/
[Fleet]: https://jetbrains.com/fleet
[HTTP]: https://developer.mozilla.org/en-US/docs/Web/HTTP
[JSON]: https://www.json.org/json-en.html
[MDN]: https://developer.mozilla.org/en-US/docs/Web/API
[Rust Language]: https://rust-lang.org
[RustRover]: https://jetbrains.com/rust
[TOML]: https://toml.io/en/
[VSCode]: https://code.visualstudio.com/docs
[YAML]: https://yaml.org/
[WAI-ARIA]: https://www.w3.org/WAI/ARIA/apg/patterns/

<a href="https://github.com/HyaenaTechnologies/data-interchange-web">
  <h1>
    <picture>
      <img src="https://github.com/HyaenaTechnologies/data-interchange-web/blob/main/assets/di_markdown.png" alt="">
    </picture>
  </h1>
</a>

[![Rust Workflow](https://github.com/HyaenaTechnologies/data-interchange-web/actions/workflows/rust.yml/badge.svg)](https://github.com/HyaenaTechnologies/data-interchange-web/actions/workflows/rust.yml)

# Data Interchange Web

Data Interchange is a Network Application Programming Interface Development Platform

## Features

- Comma-Separated Values
- DNS Protocol
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

- [Rust][Rust Language]
- [JetBrains Fleet][Fleet]
- [JetBrains RustRover][RustRover]
- [Mozilla Developer Network Web Documentation][MDN]
- [Visual Studio Code][VSCode]
- [Web Accessibility Initiative][WAI-ARIA]

## Build Web Server

```shell
git clone

cargo build --release \ 
mv ./target/release/data-interchange-server ./binary \ 
./binary/data-interchange-server serve
```

OR

```shell
git clone

make build-release \ 
mv ./target/release/data-interchange-server ./binary \ 
./binary/data-interchange-server serve
```

OR

```shell
git clone

./shell/build.sh \ 
mv ./target/release/data-interchange-server ./binary \ 
./binary/data-interchange-server serve
```

## Install Web Server

```shell
echo 'export PATH="$PATH:/usr/local/bin/htdinet"' >> ~/.bashrc && sudo echo 'export PATH="$PATH:/usr/local/bin/htdinet"' >> /etc/skel/.bashrc
```