
## Setup go path in VS code
settings.json file
```
{
  "window.zoomLevel": 1,
  "editor.renderWhitespace": "all",
  "files.trimTrailingWhitespace": true,
  "editor.renderControlCharacters": false,
  "breadcrumbs.enabled": true,
  "editor.minimap.enabled": true,
  "editor.insertSpaces": false,
  "editor.detectIndentation": false,
  "diffEditor.ignoreTrimWhitespace": true,
  "go.formatTool": "goimports",
  "go.useLanguageServer": true,
  "go.gopath": "/Users/junaid/Documents/gitRepos/go-workspace",
}
```

## Setup go path in .bash_profile
```
# GOLANG SETTINGS
export GOPATH=$HOME/Documents/gitRepos/go-workspace # don't forget to change your path correctly!
export PATH=$PATH:$GOPATH
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
export GOBIN=$GOPATH/bin

```
## Under workspace->src directory
```
go install apiProject
```
## Run from src folder
```
./apiProject
```

## Dependency installed
```
  go get -u github.com/gorilla/mux
  go get -u github.com/stamblerre/gocode
  go get -u github.com/ramya-rao-a/go-outline
  go get -u github.com/sqs/goreturns
  go get -u github.com/spf13/viper
```