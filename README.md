# asdf-bootstrap
A program that reads an .tool-versions and generates the `asdf` commands to sync up with it. The use case is as follows.
You're using asdf to manage your tools, you commit a `.tool-versions` file into a repo. Someone else, also using `asdf` 
clones the repo. They need to install all the required plugins and tool versions implied by the `.tool-versions`. That's
where `asdf-bootstrap` comes in. Run it, and save it's output as a script in the repo. 

# An Example

```bash
go build
./asdf-bootsstrap 
```

The output from that should look like what you find here in `./bin/asdf-bootstrap.sh`.

