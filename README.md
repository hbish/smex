# smex

A blazing fast CLI application that processes sitemaps in Go Lang. 

---

[![GoDoc](https://godoc.org/github.com/hbish.smex?status.svg)](https://godoc.org/github.com/hbish/smex)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fhbish%2Fsmex.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fhbish%2Fsmex?ref=badge_shield)
![](https://img.shields.io/badge/license-MIT-blue.svg)
[![smex](https://circleci.com/gh/hbish/smex.svg?style=shield)](https://circleci.com/gh/hbish/smex)
[![Go Report Card](https://goreportcard.com/badge/github.com/hbish/smex)](https://goreportcard.com/report/github.com/hbish/smex)
[![Coverage Status](https://coveralls.io/repos/github/hbish/smex/badge.svg?branch=develop)](https://coveralls.io/github/hbish/smex?branch=develop)

Smex is short for **S**ite**M**ap **EX**trator. It can handle various operations related to sitemaps with more being 
build over time. Smex has grown from the frustration of migrating multiple websites in the past and handling changes in 
URL semantics and massive amount of assets like images.  

Smex supports: 
- [x] extraction of urls
- [x] process local/remote sitemaps
- [x] output to csv/json
- [x] pattern matching on urls
- [x] extraction of images
- [x] extraction of video
- [x] extraction of news
- [ ] basic auth to retrieve sitemap
- [ ] url status checking
- [ ] support sitemap indexes
- [ ] support for sitemaps with multiple languages

Note: smex is not a sitemap validator and would not check the validity of sitemaps against the xsd. It will try to parse
the sitemaps on best effort.

[![asciicast](https://asciinema.org/a/327587.svg)](https://asciinema.org/a/327587)

## Usages

### Extract

To extract information from sitemap

The following command extracts only the urls from yoast's post sitemaps and prints to stdout

`smex extract https://yoast.com/post-sitemap.xml --remote --loc`

You can also perform the extraction at a sitemap locally

`smex extract ~/Download/sitemap.xml`

By default, smex write to stdout. Due to the limited space only important fields are displayed. If you want a full extract 
of the sitemap, then add `--format` or `-f` flag to write to a file, currently smex only support csv & json.

`smex extract https://yoast.com/post-sitemap.xml --remote --format csv`

To filter the URLs you can supply a valid regex pattern using `--pattern` or `-p` flag

`smex extract https://yoast.com/post-sitemap.xml --remote  --pattern ".*seo.*" --format csv`

The `--output` or `-o` flag will change the filename (no extension required), this is defaulted to `smex-output.(csv|json)`

### Check

To check the status of the pages

__TODO: this feature has not yet being implemented__

### Help

To get help simply run `smex` without any commands and flags.

## Installation

### Using Go

`go get -u github.com/hbish/smex`

### Using Homebrew

`brew install hbish/taps/smex`

### Binary

Currently cross-compiled for:

- Mac (64 bit)
- Linux (32/64 bit)
- Windows (32/64 bit)

Latest versions can be downloaded via [Releases](https://github.com/hbish/smex/releases).

## Documentation

For package documentation please check on [pkg.go.dev](https://pkg.go.dev/github.com/hbish/smex).

## Contribute

### Getting the code

Clone the repo

`git clone git@github.com:hbish/smex.git`

Initialise local environment and install commit hook

`make init`

Running the source

`go run github.com/hbish/smex [command] [flags...]`

Running tests

`make test`

Make the changes, the linter is set up to run when you commit your code, if it passes feel free to submit a PR!

## Last Thing

Smex is my first stab at building with Go. If you have any feedback, comments or notice any bugs, I'd be more than happy 
to consider them through the github issue tracker or better yet send me a pull request! 



## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fhbish%2Fsmex.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fhbish%2Fsmex?ref=badge_large)
