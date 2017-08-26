#!/bin/sh

go get github.com/mattn/thank-you-stars/github-star
go list -f '{{join .Deps "\n"}}' |  xargs go list -f '{{if not .Standard}}{{.ImportPath}}{{end}}' | xargs -n 1 github-star
