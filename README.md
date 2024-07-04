# My Kùzu Bindings for Go

A learning exercise in cgo, using [`c-for-go`][c-for-go] to wrap the [C API][c-api] for the embeddable graph database [Kùzu][kuzu] in a Go library

[c-api]: https://kuzudb.com/api-docs/c/kuzu_8h.html
[kuzu]: https://kuzudb.com
[c-for-go]: https://c.for-go.com

## Caveats

This is all *very* specific to my system, and just barely reached proof-of-concept status. Don't use it at work. Heck, I wouldn't even recommend using it for play yet.

## Knowingly broken

Ignoring Arrow functionality until I get better at cgo and `c-for-go`
