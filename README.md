WIP
---

This project is a *work in progress*. The implementation is *incomplete* and
subject to change. The documentation can be inaccurate.

status
======

This package gathers information about the status of Go packages.

The following information is currently made available for packages:

- CanGet:
	- Specifies if it's possible to "go get" the package.
- CanBuild:
	- Specifies if it's possible to "go build" the package.

Ideas
-----

The following ideas may be implemented in the future.

- Use the [golint][] tool locate potential issues with established conventions.
- Example coverage:
	- How many functions and methods have examples?
- Documentation maturity:
	- How many exported identifiers have comments?
	- What is the quality of these comments, in regards to [godoc conventions][].
- Test coverage:
	- How many functions and methods have test cases?
- Benchmark coverage:
	- How many functions and methods have benchmarking tests?

[godoc conventions]: http://golang.org/doc/articles/godoc_documenting_go_code.html
[golint]: https://github.com/golang/lint/golint

Documentation
-------------

Documentation provided by GoDoc.

- [status][]

[status]: http://godoc.org/github.com/mewmew/status

Installation
------------

	$ go get github.com/mewmew/status/cmd/gos

public domain
-------------

This code is hereby released into the *[public domain][]*.

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
