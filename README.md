### The beginnings of a simple rest client for Atlassian Jira

Build by documenting their API with a Swagger spec, and then generating with go-swagger a Go client.

### Development

I have a GOPATH_source_me.sh at the top folder for this work.

GOPATH_source_me.sh

```bash
echo "you should source me, not execute me! >>  source <me>"
export GOPATH=/Users/jaw/me/Development/go-opensource-work
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOPATH/bin
echo $GOPATH
```

That folder contains:

```
>ls -a /Users/jaw/me/Development/go-opensource-work
.                       bin
..                      go-jira-restclient.iml
.idea                   pkg
GOPATH_source_me.sh     src
```

When I want to start working on this, I source that `. GOPATH_source_me.sh`.

I use IntelliJ with the golang plugin.

I ran `go get -u github.com/go-swagger/go-swagger/cmd/swagger` and in that created directory I then ran `go install ./...` to install the `swagger` binary in bin/ (and see my PATH edit above so it would include that bin/ folder).

See the `Makefile` for generating, building, and running.

When editing the `swagger.yml` it is helpful to test syntax here: http://editor.swagger.io/#/.
