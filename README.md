go-tag-test
===========

Testing repository for versioned `go get` imports.

Patch go
--------

Apply the included patch on a checkout of the go source-code, build and run!

In particular, this allows `@<branch>` or `@<tag>` to Github checkouts, ex:

    go get github.com/msiebuhr/go-tag-test@v0.0.0
	go get github.com/msiebuhr/go-tag-test@v0.0.1

Will give you two checkouts of this repository at the commits corresponding to
the given tags.

If the tag is placed somewhere else, `go get -u` will fix it (i.e. have
`stable` or `builds-ok`-tags from CI systems).

Test program
------------

A quick test-program is included, which imports this library (without any tags)
and prints whatever `TagTest.Tag()` returns.

    go run ./TestProgram/main.go

Try changing the import-statements in the top  of `main.go` and verify the
output changes.

Corner-cases
------------

Tonnes!

 - Source from any other place than Github.
 - You can include different versions of the program without error, which will
   lead to all sorts of corner-cases.
 - Simplistic (I tried to make the patch as minimal as possible, so a lot of
   other stuff is probably broken rather badly.)
 - ...

