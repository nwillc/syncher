[![license](https://img.shields.io/github/license/nwillc/syncher.svg)](https://tldrlegal.com/license/-isc-license)
[![CI](https://github.com/nwillc/syncher/workflows/CI/badge.svg)](https://github.com/nwillc/syncher/actions?query=workflow%3CI)
[![Go Report Card](https://goreportcard.com/badge/github.com/nwillc/syncher)](https://goreportcard.com/report/github.com/nwillc/syncher)
-----
# asdf syncher
An [asdf](https://github.com/asdf-vm/asdf) plugin that reads an `.tool-versions` and generates the `asdf` commands to sync up with it. The use case is as follows.
You and your teammates are using asdf to manage tools, you commit a `.tool-versions` update into a repo. Someone else 
clones the repo. They need to install all the required plugins and tool versions implied by the `.tool-versions`. That's
where `syncher` comes in. Run it, and save it's output as a script in the repo. 

# Usage

```bash
asdf plugin-add syncher https://github.com/nwillc/syncher.git
asdf list all syncher
asdf install syncher 0.0.13
asdf local syncher 0.0.13
asdf syncher
```

The output from that will be a shell script you can share.
