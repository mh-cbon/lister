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
{{template "glide/install" .}}

## Usage

#### $ {{exec "lister" "-help" | color "sh"}}

## Cli examples

```sh
# Create a slice of Tomate to Tomates to tomates.go
lister Tomate:Tomates
# Create a slice of strings to stdout
lister -p main - string:StringSlice
```
# API example

Following example demonstates a program using it to generate a slice of `Tomate`

#### > {{cat "demo/main.go" | color "go"}}

Following is the generated code for a slice of `Tomate`.

#### > {{cat "demo/tomates.go" | color "go"}}

# Recipes

#### Release the project

```sh
gump patch -d # check
gump patch # bump
```

# History

[CHANGELOG](CHANGELOG.md)
