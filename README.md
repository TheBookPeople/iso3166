iso3166
=======

A simple [Go](https://golang.org) [ISO-3166-2](https://en.wikipedia.org/wiki/ISO_3166-2) library for dealing with country/regions/county codes.

Usage
-----

Import with:

```go
    import "github.com/TheBookPeople/iso3166"
```

Only one function exists so far:

```go
    text, err = iso3166.Decode("GB", "", true) // will return "United Kingdom"
    text, err = iso3166.Decode("GB", "CHS", true) // will return "Cheshire"
    text, err = iso3166.Decode("", "CHS", true) // will return an error
    text, err = iso3166.Decode("", "foobar", true) // will return an error
```

The final bool argument specifies if additional values are to be used. E.g. GB-DVV is not a valid ISO code,
but is still used in the wild. If you are translating real-world addresses, best set this to true.


Contributing
------------

This data has not all been tested. Tests for the regions you care about are welcome.

Copyright
---------
Copyright The Book People Ltd. 2016

License
-------
This library is distributed under the GPLv3 license found in the LICENSE file.
