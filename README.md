[![CircleCI](https://circleci.com/gh/stoewer/go-week.svg?style=svg)](https://circleci.com/gh/stoewer/go-week)
[![codecov](https://codecov.io/gh/stoewer/go-week/branch/master/graph/badge.svg)](https://codecov.io/gh/stoewer/go-week)
[![GoDoc](https://godoc.org/github.com/stoewer/go-week?status.svg)](https://godoc.org/github.com/stoewer/go-week)
---

go-week
=======

The package `go-week` provides a simple data type representing a week date as defined by [ISO 8601](https://en.wikipedia.org/wiki/ISO_week_date).

Versions and stability
----------------------

This package can be considered stable and ready to use. All releases follow the rules of 
[semantic versioning](http://semver.org).

Although the master branch is supposed to remain stable, there is not guarantee that braking changes will not
be merged into master when major versions are released. Therefore the repository contains version tags in 
order to support vendoring tools. The tag names follow common conventions and have the following format `v1.0.0`. 
This package supports Go modules introduced with version 1.11.

Dependencies
------------

### Build dependencies

* `github.com/pkg/errors`

### Test dependencies

* `github.com/DATA-DOG/go-sqlmock`
* `github.com/lib/pq` (integration tests only)
* `github.com/stretchr/testify`

Run unit and integration tests
------------------------------ 

Since some of the linters ran by gometalinter don't support go modules yet, all dependencies have to be
loaded to the vendor directory first and gometalinter itself must run disabled module support:

```
go mod vendor
GO111MODULE=off gometalinter --config=.gometalinter.json --deadline=10m .
```

To run the test use the following commands:

```
# without integration tests
go test .

# with integration tests (requires test db)
go test -tags=integration .
```

License
-------

This project is open source an published under the [MIT license](LICENSE).
