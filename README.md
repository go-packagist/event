# event

![Go](https://badgen.net/badge/Go/%3E=1.16/orange)
[![Go Version](https://badgen.net/github/release/go-packagist/event/stable)](https://github.com/go-packagist/event/releases)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-packagist/event)](https://pkg.go.dev/github.com/go-packagist/event)
[![codecov](https://codecov.io/gh/go-packagist/event/branch/master/graph/badge.svg?token=5TWGQ9DIRU)](https://codecov.io/gh/go-packagist/event)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-packagist/event)](https://goreportcard.com/report/github.com/go-packagist/event)
[![tests](https://github.com/go-packagist/event/actions/workflows/go.yml/badge.svg)](https://github.com/go-packagist/event/actions/workflows/go.yml)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

## Installation

```bash
go get github.com/go-packagist/event
```

## Usage

```go
package main

import "github.com/go-packagist/event"

type Event struct {
	Stop bool
}

func (e *Event) IsStop() bool {
	return e.Stop
}

func (e *Event) Val() string {
	return "event"
}

type Listener1 struct {
}

func (l *Listener1) Handle(event event.Event) {
	println("listener1:" + event.(*Event).Val())

	event.(*Event).Stop = true
}

type Listener2 struct{}

func (l *Listener2) Handle(event event.Event) {
	println("listener2:" + event.(*Event).Val())
}

var _ event.Event = (*Event)(nil)
var _ event.Listener = (*Listener1)(nil)
var _ event.Listener = (*Listener2)(nil)

func main() {
	d := event.NewDispatcher()

	e := &Event{
		Stop: false,
	}

	d.Listen(e, &Listener1{})
	d.Listen(e, &Listener2{})

	d.Dispatch(e) // echo: listener1:event (because listener1 set Stop to true)
}
```

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.