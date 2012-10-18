status
======

This package gathers information about the status of Go packages.

The following information is currently made available for packages:

   * CanGet:
       - Specifies if it's possible to "go get" the package.
   * CanBuild:
       - Specifies if it's possible to "go build" the package.

ideas
-----

The following ideas may be implemented in the future.

   * Documentation maturity:
       - How many exported identifiers have comments?
       - What is the quality of these comments, in regards to
         [godoc conventions][].
   * Test coverage:
       - How many functions and methods have test cases?
   * Benchmark coverage:
       - How many functions and methods have benchmarking tests?

[godoc conventions]: http://golang.org/doc/articles/godoc_documenting_go_code.html

documentation
-------------

Documentation provided by GoPkgDoc:

   - [status][]

[status]: http://go.pkgdoc.org/github.com/mewmew/status

installation
------------

    $ go get github.com/mewmew/status/cmd/gos

public domain
-------------

I hereby release this code into the *[public domain][]*.

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
