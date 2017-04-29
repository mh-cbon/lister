---
License: MIT
LicenseFile: LICENSE
LicenseColor: yellow
---
# {{.Name}}

{{template "badge/travis" .}} {{template "badge/appveyor" .}} {{template "badge/goreport" .}} {{template "badge/godoc" .}} {{template "license/shields" .}}

{{pkgdoc}}

# {{toc 5}}

# Install
{{template "go/install" .}}

## Usage

#### $ {{exec "lister" "-help" | color "sh"}}

## Cli examples

```sh
# Create a typed slice version of Tomate to Tomates
lister tomates_gen.go Tomate:Tomates
```
# API example

Following example demonstates a program using it to generate a lister version of a type.

#### > {{cat "demo/lib.go" | color "go"}}

Following code is the generated implementation of `Tomates` type.

#### > {{cat "demo/vegetables_gen.go" | color "go"}}

# Recipes

#### Release the project

```sh
gump patch -d # check
gump patch # bump
```

# History

[CHANGELOG](CHANGELOG.md)
