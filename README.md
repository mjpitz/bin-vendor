# bin-vendor

A very simple way to vendor binary tools your projects use.
Regardless of what language you work with, you often need specific versions of developer tools.
`bin-vendor` makes it easy to manage tool versions.

## Status

![experimental](https://img.shields.io/badge/-experimental-yellow?style=for-the-badge)
[![Release](https://img.shields.io/github/release/mjpitz/bin-vendor.svg?style=for-the-badge)](https://github.com/mjpitz/bin-vendor/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=for-the-badge)](/LICENSE)

## Getting started

1. Download and install the `bin-vendor` tool.

1. Create a `bin.yaml` file in your projects directory.
   See this projects `bin.yaml` file for an example.

1. Run `bin-vendor`.

That's it!
All the tools you requested will be stored under the `bin/` directory of the project.
You'll probably want to add your `bin/` directory to your `.gitignore`.

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
        #   version: v0.0.4
        #   config: alternate-bin.yaml

      - name: Build
        run: |
          # invoke your binaries here ....
```
