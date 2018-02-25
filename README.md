[![Build Status](https://travis-ci.org/stoewer/go-week.svg?branch=master)](https://travis-ci.org/stoewer/go-week)
[![Coverage Status](https://coveralls.io/repos/github/stoewer/go-week/badge.svg?branch=master)](https://coveralls.io/github/stoewer/go-week?branch=master)
[![GoDoc](https://godoc.org/github.com/stoewer/go-week?status.svg)](https://godoc.org/github.com/stoewer/go-week)
---

go-week
=======

The package `go-week` provides a simple data type representing a week date as defined by [ISO 8601](https://en.wikipedia.org/wiki/ISO_week_date).

Versions and stability
----------------------

This package can be considered mostly stable but *not yet* ready to use. All releases follow the rules of 
[semantic versioning](http://semver.org).

Although the master branch is supposed to remain stable, there is not guarantee that braking changes will not
be merged into master when major versions are released. Therefore the repository contains version tags in 
order to support vendoring tools such as [glide](https://glide.sh). The tag names follow common conventions 
and have the following format `v1.0.0`.

Dependencies
------------

### Build dependencies

* `github.com/pkg/errors`

### Test dependencies

* `github.com/DATA-DOG/go-sqlmock`
* `github.com/lib/pq` (integration tests only)
* `github.com/stretchr/testify/...`

Run unit and integration tests
------------------------------ 

Commands used for linting and testing:

```
gometalinter --config=.gometalinter.json .
go test .

# with integration tests (requires test db)
go test -tags=integration .
```

License
-------

This project is open source an published under the [MIT license](LICENSE).
