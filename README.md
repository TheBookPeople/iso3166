iso3166
=======

A simple [Go](https://golang.org) [ISO-3166-2](https://en.wikipedia.org/wiki/ISO_3166-2) library.

Usage
-----

Import with:


    import "github.com/TheBookPeople/iso3166"


Only one function exists so far:

    text, err = iso3166.Decode("GB", "") // will return "United Kingdom"
    text, err = iso3166.Decode("GB", "CHS") // will return "Cheshire"
    text, err = iso3166.Decode("", "CHS") // will return an error
    text, err = iso3166.Decode("", "foobar") // will return an error


Contributing
------------

This data has not all been tested. Tests for the regions you care about are welcome.

Copyright
---------
Copyright The Book People Ltd. 2016

License
-------
This library is distributed under the GPLv3 license found in the LICENSE file.
