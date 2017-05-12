---
License: MIT
LicenseFile: LICENSE
LicenseColor: yellow
---
# {{.Name}}

{{template "badge/travis" .}} {{template "badge/appveyor" .}} {{template "badge/goreport" .}} {{template "badge/godoc" .}} {{template "license/shields" .}}

{{pkgdoc}}

s/Choose your gun!/[Aux armes!](https://www.youtube.com/watch?v=hD-wD_AMRYc&t=7)/

# {{toc 5}}

# Install
{{template "go/install" .}}

## Usage

#### $ {{exec "lister" "-help" | color "sh"}}

## Cli examples

```sh
# Create a typed slice version of Tomate to Tomates
lister Tomate:gen/Tomates
```
# API example

Following example demonstates a program using it to generate a slice of `Tomate`

#### > {{cat "demo/lib.go" | color "go"}}

Following code is the generated code for a slice of `Tomate`.

#### > {{cat "demo/gen/tomates.go" | color "go"}}

# Recipes

#### Release the project

```sh
gump patch -d # check
gump patch # bump
```

# History

[CHANGELOG](CHANGELOG.md)
