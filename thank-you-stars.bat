@echo off

setlocal enabledelayedexpansion

go get github.com/mattn/thank-you-stars/github-star
set DEPS=go list -f "{{join .Deps \"\n\"}}"
for /F %%i in ('%DEPS%') do (
  set PKGLIST=go list -f "{{if not .Standard}}{{.ImportPath}}{{end}}" %%i ^| findstr "^github\.com/"
  for /F %%j in ('!PKGLIST!') do (
    github-star %%j
    echo Stared %%j
  ) 
)
