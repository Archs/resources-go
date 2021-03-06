resources-go
============

An assets-loading package for Go.

Applications can use this package to load assets from zip-files (incuding a zip file bundled in the executable),
the filesystem, or other sources through a single interface. Also allows for the building of a search path to access
files sequentially through a set of application defined locations.

This is the development version of the API. It may change at any time, if you want stability use one of the versions 
listed below. They are also stored in branches of the git repository.

[![Build Status](https://travis-ci.org/cookieo9/resources-go.svg)](https://travis-ci.org/cookieo9/resources-go)
[![GoDoc](https://godoc.org/github.com/cookieo9/resources-go?status.png)](https://godoc.org/github.com/cookieo9/resources-go)
[![Coverage](http://gocover.io/_badge/github.com/cookieo9/resources-go)](http://gocover.io/github.com/cookieo9/resources-go)

Package Path
------------

This package is go-gettable with the path:

	github.com/cookieo9/resources-go

Previous Versions
-----------------

 - [v1](https://github.com/cookieo9/resources-go/tree/v1):
       `gopkg.in/cookieo9/resources-go.v1`
       [![Build Status](https://travis-ci.org/cookieo9/resources-go.svg?branch=v1)](https://travis-ci.org/cookieo9/resources-go)
       [![GoDoc](https://godoc.org/gopkg.in/cookieo9/v1/resources-go?status.png)](https://godoc.org/gopkg.in/cookieo9/resources-go.v1)
       [![Coverage](http://gocover.io/_badge/gopkg.in/cookieo9/resources-go.v1)](http://gocover.io/gopkg.in/cookieo9/resources-go.v1)

 - [v2](https://github.com/cookieo9/resources-go/tree/v2):
       `gopkg.in/cookieo9/resources-go.v2`
       [![Build Status](https://travis-ci.org/cookieo9/resources-go.svg?branch=v2)](https://travis-ci.org/cookieo9/resources-go)
       [![GoDoc](https://godoc.org/gopkg.in/cookieo9/v2/resources-go?status.png)](https://godoc.org/gopkg.in/cookieo9/resources-go.v2)
       [![Coverage](http://gocover.io/_badge/gopkg.in/cookieo9/resources-go.v2)](http://gocover.io/gopkg.in/cookieo9/resources-go.v2)

To migrate code that used v1 (or v2) of the API, change the import path: `github.com/cookieo9/resources-go/v1/resources` to 
`gopkg.in/cookieo9/resources-go.v1`. All the code will still be under the same package name, so your code should work
without any further modification.
    
Documentation
-------------

The autogenerated online documentation can be found at: <https://godoc.org/github.com/cookieo9/resources-go>

Embedding Zip-Files
-------------------

To embed a zip file into your executable do the following:
 - Create executable (eg: go build -o myApp)
 - Create zip file  (eg: zip -r assets.zip assets)
 - Append zip file to executable (eg: cat assets.zip >> myApp)
 - Adjust the offsets in the zip header (optional) (zip -A myApp)

License
-------
http://cookieo9.mit-license.org/2012
