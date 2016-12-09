
# go-trends
Command-Line app for querying trending GitHub repositories.

## Install
`go get github.com/pmbenjamin/go-trends/cmd/gittr`

## Usage
```
$ gittr
Usage of gittr:
  -l string
    	Search by Language

```

```
$ gittr -l go
Name: docker
Desc: Docker - the open-source application container engine
URL : https://api.github.com/repos/docker/docker

Name: go
Desc: The Go programming language
URL : https://api.github.com/repos/golang/go

Name: lantern
Desc: :izakaya_lantern: Open Internet for everyone. Lantern is a free application that delivers fast, reliable and secure access to the open Internet for users in censored regions. It uses a variety of techniques to stay unblocked, including domain fronting, p2p, and pluggable transports.
URL : https://api.github.com/repos/getlantern/lantern

...
```