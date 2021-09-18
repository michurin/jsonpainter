# Fuzzy JSON highlighting ![build](https://github.com/michurin/jsonpainter/actions/workflows/ci.yaml/badge.svg) [![codecov](https://codecov.io/gh/michurin/jsonpainter/branch/master/graph/badge.svg?token=WHEKURGJZ6)](https://codecov.io/gh/michurin/jsonpainter) [![go reference](https://pkg.go.dev/badge/github.com/michurin/jsonpainter.svg)](https://pkg.go.dev/github.com/michurin/jsonpainter)

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

## Examples

    fmt.Println(paintjson.String(`{"x":12}`))

## Install

    go get github.com/michurin/jsonpainter

## Todo

- Streaming: obtain `io.Reader`
- CLI tool
