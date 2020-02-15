# dotenv

[![Build Status][build-img]][build-url]
[![GoDoc][doc-img]][doc-url]
[![Go Report Card][reportcard-img]][reportcard-url]
[![Coverage][coverage-img]][coverage-url]

dotenv is a module to load environment variables from a `.env` files.

## Features

* Simplest API.
* Dependency-free.

## Install

Go version 1.13

```
go get github.com/cristalhq/dotenv
```

## Example

```go
errLoad := dotenv.Load("my_dev.env")
if errLoad != nil {
    os.Exit(1)
}
```

## Documentation

See [these docs](https://godoc.org/github.com/cristalhq/dotenv).

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/dotenv/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/dotenv/actions
[doc-img]: https://godoc.org/github.com/cristalhq/dotenv?status.svg
[doc-url]: https://godoc.org/github.com/cristalhq/dotenv
[reportcard-img]: https://goreportcard.com/badge/cristalhq/dotenv
[reportcard-url]: https://goreportcard.com/report/cristalhq/dotenv
[coverage-img]: https://codecov.io/gh/cristalhq/dotenv/branch/master/graph/badge.svg
[coverage-url]: https://codecov.io/gh/cristalhq/dotenv
