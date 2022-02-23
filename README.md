<p align="center">
  <h1 align="center">Go Hateoas - Package to generate hateoas level to structs in go</h1>
  <p align="center">
    <a href="https://pkg.go.dev/github.com/booscaaa/go-hateoas"><img alt="Reference" src="https://img.shields.io/badge/go-reference-purple?style=for-the-badge"></a>
    <a href="https://github.com/booscaaa/go-hateoas/releases/latest"><img alt="Release" src="https://img.shields.io/github/v/release/booscaaa/go-hateoas.svg?style=for-the-badge"></a>
    <a href="/LICENSE"><img alt="Software License" src="https://img.shields.io/badge/license-MIT-red.svg?style=for-the-badge"></a>
    <a href="https://github.com/booscaaa/go-hateoas/actions/workflows/test.yaml"><img alt="Test status" src="https://img.shields.io/github/workflow/status/booscaaa/go-hateoas/Test?label=TESTS&style=for-the-badge"></a>
    <a href="https://codecov.io/gh/booscaaa/go-hateoas"><img alt="Coverage" src="https://img.shields.io/codecov/c/github/booscaaa/go-hateoas/master.svg?style=for-the-badge"></a>
  </p>
</p>

<br>

## Why?

This project is part of my personal portfolio, so, I'll be happy if you could provide me any feedback about the project, code, structure or anything that you can report that could make me a better developer!

Email-me: boscardinvinicius@gmail.com

Connect with me at [LinkedIn](https://www.linkedin.com/in/booscaaa/).

<br>

## Functionalities

- Generate hateoas struct for restfull level

<br>

## Starting

### Installation

```sh
go get github.com/booscaaa/go-hateoas
```

<br>

### Usage

```go
import "github.com/booscaaa/go-hateoas/hateoas"

mockStruct := struct {
    ID      int
    Make    string
    Model   string
    Mileage int
    hateoas.Hateoas
}{
    ID:      10,
    Make:    "Ford",
    Model:   "Taurus",
    Mileage: 200000,
}

hateoas.Generate(&mockStruct, "vehicle", "http://localhost:3000")
```

The output:

```js
{
    "ID": 10,
    "Make": "Ford",
    "Model": "Taurus",
    "Mileage": 200000,
    "_links": [
        {
            "href": "http://localhost:3000/vehicle",
            "method": "POST",
            "doc": "http://localhost:3000/doc/index.html#/vehicle/post_vehicle",
            "key": "new"
        },
        {
            "href": "http://localhost:3000/vehicle/11",
            "method": "GET",
            "doc": "http://localhost:3000/doc/index.html#/vehicle/get_vehicle",
            "key": "next"
        },
        {
            "href": "http://localhost:3000/vehicle/9",
            "method": "GET",
            "doc": "http://localhost:3000/doc/index.html#/vehicle/get_vehicle",
            "key": "prev"
        }
    ],
    "_embeded": [
        {
            "href": "http://localhost:3000/vehicle/10",
            "method": "GET",
            "doc": "http://localhost:3000/doc/index.html#/vehicle/get_vehicle"
        },
        {
            "href": "http://localhost:3000/vehicle/10",
            "method": "PUT",
            "doc": "http://localhost:3000/doc/index.html#/vehicle/put_vehicle__id_"
        },
        {
            "href": "http://localhost:3000/vehicle/10",
            "method": "DELETE",
            "doc": "http://localhost:3000/doc/index.html#/vehicle/delete_vehicle__id_"
        }
    ]
}
```

## Contributing

You can send how many PR's do you want, I'll be glad to analyze and accept them! And if you have any question about the project...

Email-me: boscardinvinicius@gmail.com

Connect with me at [LinkedIn](https://www.linkedin.com/in/booscaaa/)

Thank you!

## License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/booscaaa/go-hateoas/blob/master/LICENSE) file for details