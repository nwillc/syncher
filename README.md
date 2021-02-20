[![license](https://img.shields.io/github/license/nwillc/syncher.svg)](https://tldrlegal.com/license/-isc-license)
[![CI](https://github.com/nwillc/syncher/workflows/CI/badge.svg?branch=master)](https://github.com/nwillc/syncher/actions?query=workflow%3CI)
[![Go Report Card](https://goreportcard.com/badge/github.com/nwillc/syncher)](https://goreportcard.com/report/github.com/nwillc/syncher)
-----
# asdf syncher
An [asdf](https://github.com/asdf-vm/asdf) plugin that reads a `.tool-versions` and generates the `asdf` commands to sync 
up with it. 

The use case is, you and your teammates are using asdf to manage tools, you commit a `.tool-versions` into a repo. 
Someone else clones the repo. They need to install all the required plugins and tool versions implied by the 
`.tool-versions`. That's where `syncher` comes in. Run it, and save it's output as a script in the repo. Others then can
clone the repo, run the `syncher` script, and they will add all the plugins needed as well as install the required tool
versions.

# Usage

```bash
# Initial setup
asdf plugin-add syncher https://github.com/nwillc/syncher.git
asdf install syncher $(asdf list all syncher | tail -1)
asdf global syncher $(asdf list all syncher | tail -1)
asdf reshim syncher

# Use going forward
asdf syncher
```

The output from `asdf syncher` will be a shell script that syncs up the plugins and installations.

# Architectures Released
The releases are available for:

 - darwin (OSX): amd64, arm64
 - linux: amd64, 386
