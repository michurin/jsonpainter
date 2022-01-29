# Fuzzy JSON highlighting

[![build](https://github.com/michurin/jsonpainter/actions/workflows/ci.yaml/badge.svg)](https://github.com/michurin/jsonpainter/actions/workflows/ci.yaml)
[![codecov](https://codecov.io/gh/michurin/jsonpainter/branch/master/graph/badge.svg?token=WHEKURGJZ6)](https://codecov.io/gh/michurin/jsonpainter)
[![go reference](https://pkg.go.dev/badge/github.com/michurin/jsonpainter.svg)](https://pkg.go.dev/github.com/michurin/jsonpainter)
[![play.golang.org](https://shields.io/badge/go.dev-play-089?logo=go&logoColor=white&style=flat)](https://go.dev/play/p/kfpFC9Pm7r6)

## Description

In fact, it doesn't perform full JSON parsing. It consider
spaces, quoted strings, brackets (including brackets balance),
colons (in context), commas... In addition,
it emphasizes quoted strings right before colons and mark them
as keys.

Thanks to this, it can treat semi-JSON strings like this:

    12:00:00 INFO Request: {"msg": "hellow"}, Response: {"ok": true}

## Demo

![JSON painter demo](https://raw.githubusercontent.com/michurin/jsonpainter/demo/demo/demo.png)

You can find [code of this demo](https://raw.githubusercontent.com/michurin/jsonpainter/demo/demo/main.go) in `demo` brunch.

You can also [play with library online](https://go.dev/play/p/kfpFC9Pm7r6).
However, this online sandbox doesn't allow to display colored text.
Instead of painted text it you will show you control sequences ([ANSI escape codes](https://en.wikipedia.org/wiki/ANSI_escape_code)) as is.

## Examples

    fmt.Println(paintjson.String(`{"x":12}`))

## Install

    go get github.com/michurin/jsonpainter

## Todo

- Streaming: obtain `io.Reader`
- CLI tool
