[![license](https://img.shields.io/github/license/nwillc/asdf-bootstrap.svg)](https://tldrlegal.com/license/-isc-license)
[![CI](https://github.com/nwillc/asdf-bootstrap/workflows/CI/badge.svg)](https://github.com/nwillc/asdf-bootstrap/actions?query=workflow%3CI)
[![Go Report Card](https://goreportcard.com/badge/github.com/nwillc/asdf-bootstrap)](https://goreportcard.com/report/github.com/nwillc/asdf-bootstrap)
-----
# asdf-bootstrap
An [asdf](https://github.com/asdf-vm/asdf) plugin that reads an `.tool-versions` and generates the `asdf` commands to sync up with it. The use case is as follows.
You and your teammates are using asdf to manage tools, you commit a `.tool-versions` update into a repo. Someone else 
clones the repo. They need to install all the required plugins and tool versions implied by the `.tool-versions`. That's
where `asdf-bootstrap` comes in. Run it, and save it's output as a script in the repo. 

# Usage

```bash
asdf plugin-add asdf-bootstrap https://github.com/nwillc/asdf-bootstrap.git
asdf list all asdf-bootstrap
asdf install asdf-bootstrap 0.0.13
asdf local asdf-bootstrap 0.0.13
asdf asdf-bootstrap
```

The output from that should look like what you find here in `./bin/asdf-bootstrap.sh`.

