# Fuzzy JSON highlighting ![build](https://github.com/michurin/jsonpainter/actions/workflows/ci.yaml/badge.svg) [![codecov](https://codecov.io/gh/michurin/jsonpainter/branch/master/graph/badge.svg?token=WHEKURGJZ6)](https://codecov.io/gh/michurin/jsonpainter) [![go reference](https://pkg.go.dev/badge/github.com/michurin/jsonpainter.svg)](https://pkg.go.dev/github.com/michurin/jsonpainter)

## Install

    go get github.com/michurin/jsonpainter

## Examples

    fmt.Println(paintjson.String(`{"x":12}`))

## Description

In fact, it doesn't perform full JSON parsing. It consider
spaces, quoted strings, brackets (including brackets balance),
colons (in context), commas... In addition,
it emphasizes quoted strings right before colons and mark them
as keys.

Thanks to this, it can treat semi-JSON strings like this:

    Body: {"ok": true}

## Todo

- Streaming: obtain `io.Reader`
- CLI tool
