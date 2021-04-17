# bin-vendor

A no-frills way to vendor tools your projects use.
Stop fighting conflicting tool versions today.

* [Status](#status)
* [Getting Started](#getting-started)
  * [Installing](#installing)
  * [Using the GitHub Action](#using-the-github-action)
  * [Using in a Makefile](#using-in-a-makefile)
* [bin.yaml](#binyaml)
  * [Variables](#variables)
  * [Replacements](#replacements)

## Status

[![Release](https://img.shields.io/github/release/mjpitz/bin-vendor.svg?style=for-the-badge)](https://github.com/mjpitz/bin-vendor/releases/latest)
[![MIT License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=for-the-badge)](/LICENSE)

## Getting started

Using a `bin.yaml` file, projects can communicate their required tooling and associated versions.
`bin-vendor` then downloads the assets and unpacks them into a local `bin` directory.

### Installing

The easiest way to get started is with curl.
Alternatively, just pop over to the releases section.

```bash
curl -sSL https://raw.githubusercontent.com/mjpitz/bin-vendor/main/install.sh | VERSION=v0.0.5 bash -s -
```

### Using the GitHub Action

The GitHub action makes working with bin-vendor a completely transparent process.
Once setup, it vendors your project's tooling and prepends the system path with the projects bin dir.
The block of yaml below shows how you can easily add bin-vendor to your build process and configure it.

```yaml
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v2

      - name: Vendor tooling
        uses: mjpitz/bin-vendor@main
        # with:
        #   version: v0.0.5
        #   config: alternate-bin.yaml

      - name: Build
        run: |
          # invoke your binaries here ...
```

### Using in a Makefile

Updating a Makefile to support binaries downloaded using bin-vendor is easy.
Simply add a line like the one below to the top of your Makefile.

```makefile
export PATH := $(shell pwd)/bin:$(PATH)

target1:
    # invoke your binaries here ...
```

## bin.yaml

`bin.yaml` is how you define the tools that a project needs.
The block below shows a full block of configuration.
Not all elements are needed.
See the `bin.yaml` in bin-vendors repo for an example.

```yaml
# -- Your projects name.
name: bin-vendor

tools:
  - # -- The name of the tool.
    name: goreleaser
    # -- The version of the tool you want to download.
    version: v0.155.1
    # -- We currently only support GitHub release assets.
    github_release:
      # -- The repo where the tool can be downloaded from.
      repo: goreleaser/goreleaser
      # -- Format of the asset name you want to download. Used when the platform
      #    specific key is not set. See section below on variables and replacements.
      format: goreleaser_{{ .OS }}_{{ .Arch }}.tar.gz
      # -- A Linux specific format.
      linux: goreleaser_Linux_{{ .Arch }}.tar.gz
      # -- An OSX specific format.
      osx: goreleaser_Darwin_{{ .Arch }}.tar.gz
      # -- A Windows specific format.
      windows: goreleaser_Windows_{{ .Arch }}.zip
      # -- Replacements for adjusting the {{ .OS }} and {{ .Arch }} variables.
      #    See section below.
      replacements:
        386: i386
        amd64: x86_64
```

### Variables

Format strings support a small set of variable substitutions.
This is intended to help with matching archives for platforms.
The table below shows the various variables and some examples of each.

| Variable         | Description          | Examples                           |
|:-----------------|:---------------------|:-----------------------------------|
| `{{ .Name }}`    | Name of the tool.    | `goreleaser`                       |
| `{{ .Version }}` | Version of the tool. | `v0.155.1`                         |
| `{{ .OS }}`      | Operating System     | `windows`, `linux`, `darwin` (OSX) |
| `{{ .Arch }}`    | Architecture         | `amd64`, `arm64`, `386`            |

### Replacements

Replacements allow you to substitute one value for another.
It's an easy way to handle case sensitivity or architecture renames.
Replacements only support `OS` and `Arch`.
