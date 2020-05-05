# smex

A blazing fast CLI application that processes sitemaps in Go Lang. 

Smex is short for **S**ite**M**ap **EX**trator. It can handle various operations related to sitemaps with more being 
build over time. Smex has grown from the frustration of migrating multiple websites in the past and handling changes in 
URL semantics and massive amount of assets like images.  

Smex supports: 
- [x] extraction of urls
- [x] process local/remote sitemaps
- [x] output to csv/json
- [x] pattern matching on urls
- [ ] url status checking
- [ ] extraction of images
- [ ] extraction of video
- [ ] extraction of news
- [ ] support for sitemaps with multiple languages

Note: smex is not a sitemap validator and would not check the validity of sitemaps.

## Usages

### Extract

To extract information from sitemap

The following command extracts yoast's post sitemaps and prints to stdout
`smex extract https://yoast.com/post-sitemap.xml --remote --loc`

You can also point the extract at a sitemap locally
`smex extract ~/Download/sitemap.xml`

### Check

To check the status of the pages

TODO

### Help

To get help simply run `smex` without any commands and flags.

## Installation

### Using Go

`go get -u github.com/hbish/smex`

### Binary

TODO

## Documentation

TODO - once published
For package documentation please check on [pkg.go.dev](https://pkg.go.dev/github.com/hbish/smex).

## Contribute

### Getting the code

Clone the repo

`git clone git@github.com:hbish/smex.git`

Initialise local environment

`make init`

Running the source

`go run github.com/hbish/smex [command] [flags...]`

Running tests

`make test`

Make the changes, the linter is set up to run when you commit your code, if it passes feel free to submit a PR!

## Last Thing

Smex is my first stab at building with Go. If you have any feedback, comments or notice any bugs, I'd be more than happy 
to consider them through the github issue tracker or better yet send me patches! 

